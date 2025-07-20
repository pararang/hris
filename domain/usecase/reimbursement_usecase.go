package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/prrng/dealls/domain/entity"
)

type SubmitReimbursementParam struct {
	UserID          uuid.UUID
	Amount          float64
	Description     string
	TransactionDate time.Time
}

type ReimbursementUseCase interface {
	SubmitReimbursement(ctx context.Context, param SubmitReimbursementParam) (*entity.Reimbursement, error)
}
