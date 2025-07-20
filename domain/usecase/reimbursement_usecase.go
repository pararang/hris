package usecase

import (
	"context"

	"github.com/pararang/hris/dto"
	"github.com/pararang/hris/entity"
)

type ReimbursementUseCase interface {
	SubmitReimbursement(ctx context.Context, param dto.SubmitReimbursementParam) (*entity.Reimbursement, error)
}
