package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/pararang/hris/domain/entity"
	"github.com/pararang/hris/dto"
)

const (
	RegularWorkingHoursInDay int8 = 8 //9am-5pm
	OvertimeMultiplier       int8 = 2
)

// AttendanceUseCase defines the interface for attendance-related business logic
type AttendanceUseCase interface {
	CreateAttendancePeriod(ctx context.Context, period dto.CreateAttendancePeriodParam) (*entity.PayrollPeriod, error)

	ClockIn(ctx context.Context, userID uuid.UUID) (*entity.Attendance, error)
	ClockOut(ctx context.Context, userID uuid.UUID) (*entity.Attendance, error)
}
