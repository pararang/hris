package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/prrng/dealls/domain/entity"
	"github.com/prrng/dealls/domain/repository"
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

func (a *attendanceUseCase) CreateAttendancePeriod(ctx context.Context, period *entity.PayrollPeriod) (*entity.PayrollPeriod, error) {
	overlappedID, err := a.attendanceRepo.FindOverlappingPayrollPeriod(ctx, period.StartDate)
	if err != nil {
		return nil, fmt.Errorf("error on check overllaped periode: %w", err)
	}

	if overlappedID != uuid.Nil {
		return nil, errors.New("period start should be greater than existing periods")
	}

	createdPeriod, err := a.attendanceRepo.CreatePayrollPeriod(ctx, period)
	if err != nil {
		return nil, fmt.Errorf("error on create payroll period: %w", err)
	}

	return createdPeriod, nil
}
