package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/prrng/dealls/domain/entity"
)

type OvertimeRepository interface {
	CreateOvertime(ctx context.Context, overtime *entity.Overtime) (*entity.Overtime, error)
	FindUserOvertimeByDate(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Overtime, error)
	CountUserOvertimeHoursInPeriod(ctx context.Context, userID, payrollPeriodID uuid.UUID) (int32, error)
}
