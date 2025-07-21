package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pararang/hris/domain/repository"
	"github.com/pararang/hris/dto"
	"github.com/pararang/hris/entity"
	"github.com/pararang/hris/libs/auth"

	"github.com/pararang/hris/libs"
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

func (a *attendanceUseCase) CreateAttendancePeriod(ctx context.Context, param dto.CreateAttendancePeriodParam) (*entity.PayrollPeriod, error) {
	hasOverlap, err := a.hasOverlappingAttendance(ctx, param.StartDate)
	if err != nil {
		return nil, fmt.Errorf("error on check overllaped periode: %w", err)
	}

	if hasOverlap {
		return nil, errors.New("period start should be greater than existing periods")
	}

	createdBy, ok := ctx.Value(auth.CtxKeyAuthUserEmail).(string)
	if !ok {
		return nil, fmt.Errorf("error on get createdBy")
	}

	createdPeriod, err := a.attendanceRepo.CreatePayrollPeriod(ctx, &entity.PayrollPeriod{
		StartDate: param.StartDate,
		EndDate:   param.EndDate,
		Status:    entity.PayrollPeriodStatusPending,
		BaseModel: entity.BaseModel{
			CreatedBy: createdBy,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error on create payroll period: %w", err)
	}

	return createdPeriod, nil
}

func (a *attendanceUseCase) ClockIn(ctx context.Context, userID uuid.UUID) (*entity.Attendance, error) {
	now := time.Now()

	// check if weekend
	if libs.IsWeekend(now) {
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
		return nil, fmt.Errorf("error on get createdBy")
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
		return nil, fmt.Errorf("error on get updatedBy")
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
