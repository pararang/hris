package usecase

import (
	"context"

	"github.com/pararang/hris/domain/entity"
	"github.com/pararang/hris/dto"
)

type OvertimeUseCase interface {
	SubmitOvertime(ctx context.Context, param *dto.SubmitOvertimeParam) (*entity.Overtime, error)
}
