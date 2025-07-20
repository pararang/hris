package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pararang/hris/domain/entity"
)

type AttendanceRepository interface {
	CreatePayrollPeriod(ctx context.Context, period *entity.PayrollPeriod) (*entity.PayrollPeriod, error)
	FindOverlappingPayrollPeriod(ctx context.Context, startDate time.Time) (uuid.UUID, error)
	FindLatestPayrollPeriod(ctx context.Context) (*entity.PayrollPeriod, error)
	GetPayrollPeriodByID(ctx context.Context, id uuid.UUID) (*entity.PayrollPeriod, error)
	SetStatusPayrollPeriod(ctx context.Context, id uuid.UUID, status string) error

	FindUserAttendanceByDate(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Attendance, error)
	CreateAttendance(ctx context.Context, attendance *entity.Attendance) (*entity.Attendance, error)
	UpdateAttendance(ctx context.Context, attendance *entity.Attendance) (*entity.Attendance, error)
	CountAttendance(ctx context.Context, userID, payrollPeriodID uuid.UUID) (int, error)
}
