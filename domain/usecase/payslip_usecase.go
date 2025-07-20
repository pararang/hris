package usecase

import (
	"context"

	"github.com/google/uuid"
)

type PayslipUseCase interface {
	GeneratePayslip(ctx context.Context, payrollPeriodID uuid.UUID) error
}
