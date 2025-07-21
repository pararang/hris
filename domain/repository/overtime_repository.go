package repository

//go:generate mockery --name=OvertimeRepository --output=. --outpkg=repository --filename=overtime_repository_mock.go

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pararang/hris/entity"
)

type OvertimeRepository interface {
	CreateOvertime(ctx context.Context, overtime *entity.Overtime) (*entity.Overtime, error)
	FindUserOvertimeByDate(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Overtime, error)
	CountUserOvertimeHoursInPeriod(ctx context.Context, userID, payrollPeriodID uuid.UUID) (int32, error)
}
