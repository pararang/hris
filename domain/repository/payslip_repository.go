package repository

import (
	"context"

	"github.com/pararang/hris/domain/entity"
)

// PayslipRepository defines the interface for payslip-related database operations
type PayslipRepository interface {
	CreatePayslip(ctx context.Context, payslip *entity.Payslip) error
}
