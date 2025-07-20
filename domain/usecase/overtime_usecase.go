package usecase

import (
	"context"

	"github.com/pararang/hris/dto"
	"github.com/pararang/hris/entity"
)

type OvertimeUseCase interface {
	SubmitOvertime(ctx context.Context, param *dto.SubmitOvertimeParam) (*entity.Overtime, error)
}
