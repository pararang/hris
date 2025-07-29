package repository

//go:generate mockery --name=ReimbursementRepository --output=. --outpkg=repository --filename=reimbursement_repository_mock.go

import (
	"context"

	"github.com/google/uuid"
	"github.com/pararang/hris/entity"
)

type ReimbursementRepository interface {
	CreateReimbursement(ctx context.Context, reimbursement *entity.Reimbursement) (*entity.Reimbursement, error)
	CountUserApprovedAmountReimbursementByPeriod(ctx context.Context, userID, payrollPeriodID uuid.UUID) (float64, error)
	GetUserReimbursementListByPeriod(ctx context.Context, userID uuid.UUID, payrollPeriodID uuid.UUID) ([]*entity.Reimbursement, error)
}
