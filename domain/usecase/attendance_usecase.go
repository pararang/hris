package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/prrng/dealls/domain/entity"
)

// AttendanceUseCase defines the interface for attendance-related business logic
type AttendanceUseCase interface {
	CreateAttendancePeriod(ctx context.Context, period *entity.PayrollPeriod) (*entity.PayrollPeriod, error)

	ClockIn(ctx context.Context, userID uuid.UUID, actorEmail string) (*entity.Attendance, error)
	ClockOut(ctx context.Context, userID uuid.UUID, actorEmail string) (*entity.Attendance, error)
}
