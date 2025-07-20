package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/prrng/dealls/domain/entity"
	"github.com/prrng/dealls/domain/repository"
	"github.com/prrng/dealls/domain/usecase"
	"github.com/prrng/dealls/libs"
	"github.com/prrng/dealls/libs/auth"
	"golang.org/x/sync/errgroup"
)

type payslipUseCase struct {
	db                *sql.DB
	payslipRepo       repository.PayslipRepository
	userRepo          repository.UserRepository
	attendanceRepo    repository.AttendanceRepository
	overtimeRepo      repository.OvertimeRepository
	reimbursementRepo repository.ReimbursementRepository
}

// NewPayslipUseCase creates a new instance of PayslipUseCase
func NewPayslipUseCase(
	db *sql.DB,
	payslipRepo repository.PayslipRepository,
	userRepo repository.UserRepository,
	attendanceRepo repository.AttendanceRepository,
	overtimeRepo repository.OvertimeRepository,
	reimbursementRepo repository.ReimbursementRepository,
) *payslipUseCase {
	return &payslipUseCase{
		db:                db,
		payslipRepo:       payslipRepo,
		userRepo:          userRepo,
		attendanceRepo:    attendanceRepo,
		overtimeRepo:      overtimeRepo,
		reimbursementRepo: reimbursementRepo,
	}
}

func (p *payslipUseCase) calculatePeriodeWorkingDay(startDate, endDate time.Time) int16 {
	var workingDays int16
	for date := startDate; date.Before(endDate); date = date.AddDate(0, 0, 1) {
		if !libs.IsWeekend(date) {
			workingDays++
		}
	}
	return workingDays
}

func (p *payslipUseCase) calculateOvertimePay(baseSalary float64, overtimeHours int32, workingDays int16) float64 {
	hourlyRate := baseSalary / float64(workingDays*int16(usecase.RegularWorkingHoursInDay))
	return float64(overtimeHours) * hourlyRate * float64(usecase.OvertimeMultiplier)
}

func (p *payslipUseCase) calculateProrateBaseSalary(baseSalary float64, attendDays, workingDays int16) float64 {
	return baseSalary * float64(attendDays/workingDays)
}

type calculateTHPParam struct {
	UserID          uuid.UUID
	PayrollPeriodID uuid.UUID
	BaseSalary      float64
	WorkingDays     int16
}

func (p *payslipUseCase) calculateTHPComponents(ctx context.Context, param calculateTHPParam) (entity.PayslipDetails, error) {
	g, childCtx := errgroup.WithContext(ctx)

	var attendanceDays int16
	var prorateBaseSalary float64
	var overtimePay float64
	var overtimeHours int32
	var reimbursementPay float64

	// Calculate Attendance and Prorated Salary
	g.Go(func() error {
		attendanceCount, err := p.attendanceRepo.CountAttendance(childCtx, param.UserID, param.PayrollPeriodID)
		if err != nil && err != sql.ErrNoRows {
			return fmt.Errorf("failed to count attendance: %w", err)
		}

		if attendanceCount == 0 {
			return nil
		}

		attendanceDays = int16(attendanceCount)
		prorateBaseSalary = p.calculateProrateBaseSalary(param.BaseSalary, attendanceDays, param.WorkingDays)
		return nil
	})

	// Calculate Overtime Hours and Pay
	g.Go(func() error {
		sumOvertimeHours, err := p.overtimeRepo.CountUserOvertimeHoursInPeriod(childCtx, param.UserID, param.PayrollPeriodID)
		if err != nil && err != sql.ErrNoRows {
			return fmt.Errorf("failed to count overtime hours: %w", err)
		}

		if sumOvertimeHours == 0 {
			return nil
		}

		overtimeHours = sumOvertimeHours
		overtimePay = p.calculateOvertimePay(param.BaseSalary, sumOvertimeHours, param.WorkingDays)
		return nil
	})

	// Calculate Reimbursement Pay
	g.Go(func() error {
		// Pass the childCtx to the repository call
		sumApprovedAmount, err := p.reimbursementRepo.CountUserApprovedAmountReimbursementByPeriod(childCtx, param.UserID, param.PayrollPeriodID)
		if err != nil && err != sql.ErrNoRows {
			return fmt.Errorf("failed to count reimbursement amount: %w", err)
		}

		if sumApprovedAmount == 0 {
			return nil
		}

		reimbursementPay = sumApprovedAmount
		return nil
	})

	if err := g.Wait(); err != nil {
		return entity.PayslipDetails{}, err
	}

	return entity.PayslipDetails{
		AttendanceDays:    attendanceDays,
		ProrateBaseSalary: prorateBaseSalary,
		OvertimePay:       overtimePay,
		OvertimeHours:     overtimeHours,
		ReimbursementPay:  reimbursementPay,
	}, nil
}

func (p *payslipUseCase) GeneratePayslip(ctx context.Context, payrollPeriodID uuid.UUID) error {
	tx, err := p.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer tx.Rollback()

	payrollPeriod, err := p.attendanceRepo.GetPayrollPeriodByID(ctx, payrollPeriodID)
	if err != nil {
		if err == sql.ErrNoRows {
			return libs.ErrPayrollPeriodNotFound{}
		}
		return fmt.Errorf("error on get payroll period: %w", err)
	}

	if payrollPeriod.Status != entity.PayrollPeriodStatusPending {
		return libs.ErrPayrollPeriodNotPending{}
	}

	employees, err := p.userRepo.ListEmployees(ctx)
	if err != nil {
		return fmt.Errorf("error on list employees: %w", err)
	}

	createdBy, ok := ctx.Value(auth.CtxKeyAuthUserEmail).(string)
	if !ok {
		return fmt.Errorf("error on get createdBy")
	}

	for _, employee := range employees {
		thpComponent, err := p.calculateTHPComponents(ctx, calculateTHPParam{
			UserID:          employee.ID,
			PayrollPeriodID: payrollPeriodID,
			BaseSalary:      employee.BaseSalary,
			WorkingDays:     p.calculatePeriodeWorkingDay(payrollPeriod.StartDate, payrollPeriod.EndDate),
		})

		if err != nil {
			return fmt.Errorf("error on calculate THP components: %w", err)
		}

		totalTHP := thpComponent.ProrateBaseSalary + thpComponent.OvertimePay + thpComponent.ReimbursementPay

		err = p.payslipRepo.CreatePayslip(ctx, &entity.Payslip{
			UserID:              employee.ID,
			PayrollPeriodID:     payrollPeriod.ID,
			BaseSalary:          employee.BaseSalary,
			ProratedBaseSalary:  thpComponent.ProrateBaseSalary,
			OvertimePay:         thpComponent.OvertimePay,
			ReimbursementAmount: thpComponent.ReimbursementPay,
			TakeHomePay:         totalTHP,
			DetailsJSON:         thpComponent,
			BaseModel: entity.BaseModel{
				CreatedBy: createdBy,
			},
		})

		if err != nil {
			return fmt.Errorf("error on create payslip: %w", err)
		}

		err = p.attendanceRepo.SetStatusPayrollPeriod(ctx, payrollPeriod.ID, entity.PayrollPeriodStatusCompleted)
		if err != nil {
			return fmt.Errorf("error on set status payroll period: %w", err)
		}

	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error on commit transaction: %w", err)
	}

	return nil
}
