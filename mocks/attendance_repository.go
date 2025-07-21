package mocks

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pararang/hris/entity"
	"github.com/stretchr/testify/mock"
)

type AttendanceRepository struct {
	mock.Mock
}

func (m *AttendanceRepository) CreatePayrollPeriod(ctx context.Context, period *entity.PayrollPeriod) (*entity.PayrollPeriod, error) {
	args := m.Called(ctx, period)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.PayrollPeriod), args.Error(1)
}

func (m *AttendanceRepository) FindOverlappingPayrollPeriod(ctx context.Context, startDate time.Time) (uuid.UUID, error) {
	args := m.Called(ctx, startDate)
	return args.Get(0).(uuid.UUID), args.Error(1)
}

func (m *AttendanceRepository) FindLatestPayrollPeriod(ctx context.Context) (*entity.PayrollPeriod, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.PayrollPeriod), args.Error(1)
}

func (m *AttendanceRepository) GetPayrollPeriodByID(ctx context.Context, id uuid.UUID) (*entity.PayrollPeriod, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.PayrollPeriod), args.Error(1)
}

func (m *AttendanceRepository) SetStatusPayrollPeriod(ctx context.Context, id uuid.UUID, status, updatedBy string) error {
	args := m.Called(ctx, id, status, updatedBy)
	return args.Error(0)
}

func (m *AttendanceRepository) FindUserAttendanceByDate(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Attendance, error) {
	args := m.Called(ctx, userID, date)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Attendance), args.Error(1)
}

func (m *AttendanceRepository) CreateAttendance(ctx context.Context, attendance *entity.Attendance) (*entity.Attendance, error) {
	args := m.Called(ctx, attendance)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Attendance), args.Error(1)
}

func (m *AttendanceRepository) UpdateAttendance(ctx context.Context, attendance *entity.Attendance) (*entity.Attendance, error) {
	args := m.Called(ctx, attendance)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Attendance), args.Error(1)
}

func (m *AttendanceRepository) CountAttendance(ctx context.Context, userID, payrollPeriodID uuid.UUID) (int, error) {
	args := m.Called(ctx, userID, payrollPeriodID)
	return args.Get(0).(int), args.Error(1)
}

func (m *AttendanceRepository) GetUserAttendanceListByPeriod(ctx context.Context, userID, payrollPeriodID uuid.UUID) ([]*entity.Attendance, error) {
	args := m.Called(ctx, userID, payrollPeriodID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.Attendance), args.Error(1)
}
