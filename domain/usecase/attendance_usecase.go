package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/prrng/dealls/domain/entity"
)

type CreateAttendancePeriodParam struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

// AttendanceUseCase defines the interface for attendance-related business logic
type AttendanceUseCase interface {
	CreateAttendancePeriod(ctx context.Context, period CreateAttendancePeriodParam) (*entity.PayrollPeriod, error)

	ClockIn(ctx context.Context, userID uuid.UUID) (*entity.Attendance, error)
	ClockOut(ctx context.Context, userID uuid.UUID) (*entity.Attendance, error)
}
