package usecase

import (
	"context"

	"github.com/prrng/dealls/domain/entity"
	"github.com/prrng/dealls/dto"
)

type ReimbursementUseCase interface {
	SubmitReimbursement(ctx context.Context, param dto.SubmitReimbursementParam) (*entity.Reimbursement, error)
}
