package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/prrng/dealls/domain/entity"
	"github.com/prrng/dealls/domain/repository"
	"github.com/prrng/dealls/libs/auth"

	"github.com/prrng/dealls/libs"
)

type attendanceUseCase struct {
	attendanceRepo repository.AttendanceRepository
}

// NewAttendanceUseCase creates a new instance of AttendanceUseCase
func NewAttendanceUseCase(
	attendanceRepo repository.AttendanceRepository,
) *attendanceUseCase {
	return &attendanceUseCase{
		attendanceRepo: attendanceRepo,
	}
}

func (a *attendanceUseCase) hasOverlappingAttendance(ctx context.Context, startDate time.Time) (bool, error) {
	_, err := a.attendanceRepo.FindOverlappingPayrollPeriod(ctx, startDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (a *attendanceUseCase) CreateAttendancePeriod(ctx context.Context, period *entity.PayrollPeriod) (*entity.PayrollPeriod, error) {
	hasOverlap, err := a.hasOverlappingAttendance(ctx, period.StartDate)
	if err != nil {
		return nil, fmt.Errorf("error on check overllaped periode: %w", err)
	}

	if hasOverlap {
		return nil, errors.New("period start should be greater than existing periods")
	}

	createdPeriod, err := a.attendanceRepo.CreatePayrollPeriod(ctx, period)
	if err != nil {
		return nil, fmt.Errorf("error on create payroll period: %w", err)
	}

	return createdPeriod, nil
}

func (a *attendanceUseCase) ClockIn(ctx context.Context, userID uuid.UUID) (*entity.Attendance, error) {
	now := time.Now()

	// check if weekend
	if now.Weekday() == time.Saturday || now.Weekday() == time.Sunday {
		return nil, libs.ErrWeekendNotAllowed{}
	}

	// check if already checkin
	existingAttendance, err := a.attendanceRepo.FindUserAttendanceByDate(ctx, userID, now)
	if err == nil {
		return existingAttendance, nil
	}

	if err != sql.ErrNoRows {
		return nil, fmt.Errorf("error on check exisitng attendance: %w", err)
	}

	// err is sql.ErrNoRows, no attendance record, create new one
	payrollPeriod, err := a.attendanceRepo.FindLatestPayrollPeriod(ctx)
	if err != nil {
		return nil, fmt.Errorf("error on find payroll period: %w", err)
	}

	createdBy, ok := ctx.Value(auth.CtxKeyAuthUserEmail).(string)
	if !ok {
		return nil, fmt.Errorf("error on get updatedBy: %w", err)
	}

	createdAttendance, err := a.attendanceRepo.CreateAttendance(ctx, &entity.Attendance{
		UserID:          userID,
		Date:            now,
		ClockinAt:       now,
		PayrollPeriodID: payrollPeriod.ID,
		BaseModel: entity.BaseModel{
			CreatedBy: createdBy,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("error on create attendance: %w", err)
	}

	return createdAttendance, nil
}

func (a *attendanceUseCase) ClockOut(ctx context.Context, userID uuid.UUID) (*entity.Attendance, error) {
	now := time.Now()

	attendanceIn, err := a.attendanceRepo.FindUserAttendanceByDate(ctx, userID, now)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, libs.ErrShouldClockIn{}
		}

		return nil, fmt.Errorf("error on find existing attendance: %w", err)
	}

	updatedBy, ok := ctx.Value(auth.CtxKeyAuthUserEmail).(string)
	if !ok {
		return nil, fmt.Errorf("error on get updatedBy: %w", err)
	}

	attendanceIn.ClockoutAt = &now
	attendanceIn.UpdatedBy = updatedBy
	attendanceIn.UpdatedAt = now

	attendanceOut, err := a.attendanceRepo.UpdateAttendance(ctx, attendanceIn)
	if err != nil {
		return nil, fmt.Errorf("error on set clockout: %w", err)
	}

	return attendanceOut, nil
}
