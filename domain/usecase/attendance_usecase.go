package usecase

import (
	"context"

	"github.com/prrng/dealls/domain/entity"
)

// AttendanceUseCase defines the interface for attendance-related business logic
type AttendanceUseCase interface {
	CreateAttendancePeriod(ctx context.Context, period *entity.PayrollPeriod) (*entity.PayrollPeriod, error)
}
