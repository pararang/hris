package mocks

import (
	"context"

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
