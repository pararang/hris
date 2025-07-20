package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/prrng/dealls/domain/entity"
)

type SubmitOvertimeParam struct {
	UserID     uuid.UUID
	HoursTaken uint8
	Date       time.Time
	Reason     string
}

type OvertimeUseCase interface {
	SubmitOvertime(ctx context.Context, param *SubmitOvertimeParam) (*entity.Overtime, error)
}
