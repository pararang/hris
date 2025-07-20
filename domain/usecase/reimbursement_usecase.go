package usecase

import (
	"context"

	"github.com/pararang/hris/domain/entity"
	"github.com/pararang/hris/dto"
)

type ReimbursementUseCase interface {
	SubmitReimbursement(ctx context.Context, param dto.SubmitReimbursementParam) (*entity.Reimbursement, error)
}
