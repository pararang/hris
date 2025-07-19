package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/prrng/dealls/domain/entity"
	"github.com/prrng/dealls/domain/repository"
)

type attendanceRepository struct {
	db *sql.DB
}

func NewAttendanceRepository(db *sql.DB) repository.AttendanceRepository {
	return &attendanceRepository{
		db: db,
	}
}

func (r *attendanceRepository) CreatePayrollPeriod(ctx context.Context, period *entity.PayrollPeriod) (*entity.PayrollPeriod, error) {
	sqlStatement := `INSERT INTO payroll_periods (start_date, end_date, status, created_by)
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	var id uuid.UUID
	err := r.db.QueryRowContext(ctx, sqlStatement, period.StartDate, period.EndDate, period.Status, period.CreatedBy).Scan(&id)
	if err != nil {
		return nil, err
	}

	period.ID = id
	return period, nil
}

func (r *attendanceRepository) FindOverlappingPayrollPeriod(ctx context.Context, startDate time.Time) (uuid.UUID, error) {
	sqlStatement := `SELECT id FROM payroll_periods WHERE end_date >= $1 LIMIT 1`

	var id uuid.UUID
	err := r.db.QueryRowContext(ctx, sqlStatement, startDate).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		return uuid.Nil, err
	}

	return id, nil
}
