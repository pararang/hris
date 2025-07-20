package usecase

import (
	"context"

	"github.com/prrng/dealls/domain/entity"
	"github.com/prrng/dealls/dto"
)

type OvertimeUseCase interface {
	SubmitOvertime(ctx context.Context, param *dto.SubmitOvertimeParam) (*entity.Overtime, error)
}
