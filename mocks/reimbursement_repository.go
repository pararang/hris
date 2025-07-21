package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/pararang/hris/entity"
	"github.com/stretchr/testify/mock"
)

type ReimbursementRepository struct {
	mock.Mock
}

func (m *ReimbursementRepository) CreateReimbursement(ctx context.Context, reimbursement *entity.Reimbursement) (*entity.Reimbursement, error) {
	args := m.Called(ctx, reimbursement)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Reimbursement), args.Error(1)
}

func (m *ReimbursementRepository) CountUserApprovedAmountReimbursementByPeriod(ctx context.Context, userID, payrollPeriodID uuid.UUID) (float64, error) {
	args := m.Called(ctx, userID, payrollPeriodID)
	return args.Get(0).(float64), args.Error(1)
}
