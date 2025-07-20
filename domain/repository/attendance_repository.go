package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/prrng/dealls/domain/entity"
)

type AttendanceRepository interface {
	CreatePayrollPeriod(ctx context.Context, period *entity.PayrollPeriod) (*entity.PayrollPeriod, error)
	FindOverlappingPayrollPeriod(ctx context.Context, startDate time.Time) (uuid.UUID, error)
	FindLatestPayrollPeriod(ctx context.Context) (*entity.PayrollPeriod, error)

	FindUserAttendanceByDate(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Attendance, error)
	CreateAttendance(ctx context.Context, attendance *entity.Attendance) (*entity.Attendance, error)
}
