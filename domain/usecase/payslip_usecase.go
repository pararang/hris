package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/pararang/hris/dto"
	"github.com/pararang/hris/entity"
)

type PayslipUseCase interface {
	GeneratePayslip(ctx context.Context, payrollPeriodID uuid.UUID) error
	GetListPayslip(ctx context.Context, userID uuid.UUID) ([]entity.Payslip, error)
	GetPayslipDetail(ctx context.Context, payslipID uuid.UUID) (dto.PayslipBreakdownResponse, error)
	GetPayrollPeriodSummary(ctx context.Context, payrollPeriodID uuid.UUID) (dto.PayrollSummaryResponse, error)
}
