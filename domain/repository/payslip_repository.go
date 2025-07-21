package repository

//go:generate mockery --name=PayslipRepository --output=. --outpkg=repository --filename=payslip_repository_mock.go

import (
	"context"

	"github.com/google/uuid"
	"github.com/pararang/hris/entity"
)

// PayslipRepository defines the interface for payslip-related database operations
type PayslipRepository interface {
	CreatePayslip(ctx context.Context, payslip *entity.Payslip) error
	ListUserPayslips(ctx context.Context, userID uuid.UUID) ([]entity.Payslip, error) // TODO: add pagination
	GetPayslipByID(ctx context.Context, payslipID uuid.UUID) (*entity.Payslip, error)
}
