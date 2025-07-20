package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/prrng/dealls/domain/entity"
	"github.com/prrng/dealls/domain/repository"
)

type overtimeRepository struct {
	*BaseRepository
}

func NewOvertimeRepository(db *sql.DB) repository.OvertimeRepository {
	return &overtimeRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

func (r *overtimeRepository) CreateOvertime(ctx context.Context, overtime *entity.Overtime) (*entity.Overtime, error) {
	sqlStat := `INSERT INTO overtimes (user_id, date, hours_taken, payroll_period_id, status, reason, created_at, created_by)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	err := r.executor(ctx).QueryRowContext(ctx, sqlStat,
		overtime.UserID, overtime.Date, overtime.HoursTaken, overtime.PayrollPeriodID, overtime.Status, overtime.Reason, overtime.CreatedAt, overtime.CreatedBy).Scan(&overtime.ID)
	if err != nil {
		return nil, err
	}

	return overtime, nil
}

func (r *overtimeRepository) FindUserOvertimeByDate(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Overtime, error) {
	sqlStat := `SELECT id, user_id, date, hours_taken, payroll_period_id, status, reason, created_at, created_by
	FROM overtimes WHERE user_id = $1 AND date = $2`

	overtime := &entity.Overtime{}
	err := r.executor(ctx).QueryRowContext(ctx, sqlStat, userID, date).Scan(
		&overtime.ID, &overtime.UserID, &overtime.Date, &overtime.HoursTaken, &overtime.PayrollPeriodID, &overtime.Status, &overtime.Reason, &overtime.CreatedAt, &overtime.CreatedBy)
	if err != nil {
		return nil, err
	}

	return overtime, nil
}

func (r *overtimeRepository) CountUserOvertimeHoursInPeriod(ctx context.Context, userID, payrollPeriodID uuid.UUID) (int32, error) {
	// status ignored for demo purpose
	sqlStat := `SELECT SUM(hours_taken) FROM overtimes WHERE user_id = $1 AND payroll_period_id = $2`

	var hoursTaken sql.NullInt32
	err := r.executor(ctx).QueryRowContext(ctx, sqlStat, userID, payrollPeriodID).Scan(&hoursTaken)
	if err != nil {
		return 0, err
	}

	return hoursTaken.Int32, nil
}
