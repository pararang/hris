package mocks

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pararang/hris/entity"
	"github.com/stretchr/testify/mock"
)

type OvertimeRepository struct {
	mock.Mock
}

func (m *OvertimeRepository) CreateOvertime(ctx context.Context, overtime *entity.Overtime) (*entity.Overtime, error) {
	args := m.Called(ctx, overtime)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Overtime), args.Error(1)
}

func (m *OvertimeRepository) FindUserOvertimeByDate(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Overtime, error) {
	args := m.Called(ctx, userID, date)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Overtime), args.Error(1)
}

func (m *OvertimeRepository) CountUserOvertimeHoursInPeriod(ctx context.Context, userID, payrollPeriodID uuid.UUID) (int32, error) {
	args := m.Called(ctx, userID, payrollPeriodID)
	return args.Get(0).(int32), args.Error(1)
}

func (m *OvertimeRepository) GetUserOvertimeListByPeriod(ctx context.Context, userID uuid.UUID, payrollPeriodID uuid.UUID) ([]*entity.Overtime, error) {
	args := m.Called(ctx, userID, payrollPeriodID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.Overtime), args.Error(1)
}
