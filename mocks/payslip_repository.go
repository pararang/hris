package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/pararang/hris/entity"
	"github.com/stretchr/testify/mock"
)

type PayslipRepository struct {
	mock.Mock
}

func (m *PayslipRepository) CreatePayslip(ctx context.Context, payslip *entity.Payslip) error {
	args := m.Called(ctx, payslip)
	return args.Error(0)
}

func (m *PayslipRepository) GetPayslipsInPeriod(ctx context.Context, periodID uuid.UUID) ([]*entity.Payslip, error) {
	args := m.Called(ctx, periodID)
	return args.Get(0).([]*entity.Payslip), args.Error(1)
}
