package repository

import (
	"context"

	"github.com/prrng/dealls/domain/entity"
)

type ReimbursementRepository interface {
	CreateReimbursement(ctx context.Context, reimbursement *entity.Reimbursement) (*entity.Reimbursement, error)
}
