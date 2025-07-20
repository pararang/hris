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
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *attendanceRepository) FindUserAttendanceByDate(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Attendance, error) {
	sqlStatement := `SELECT id, date, user_id, clockin_at, clockout_at, payroll_period_id, created_by FROM attendances WHERE user_id = $1 AND date = $2`

	var attendance entity.Attendance
	err := r.db.QueryRowContext(ctx, sqlStatement, userID, date.Format(time.DateOnly)).
		Scan(&attendance.ID, &attendance.Date, &attendance.UserID, &attendance.ClockinAt, &attendance.ClockoutAt, &attendance.PayrollPeriodID, &attendance.CreatedBy)
	if err != nil {
		return nil, err
	}

	return &attendance, nil
}

func (r *attendanceRepository) FindLatestPayrollPeriod(ctx context.Context) (*entity.PayrollPeriod, error) {
	sqlStatement := `SELECT id, start_date, end_date, status, created_by FROM payroll_periods ORDER BY end_date DESC LIMIT 1`

	var period entity.PayrollPeriod
	err := r.db.QueryRowContext(ctx, sqlStatement).Scan(&period.ID, &period.StartDate, &period.EndDate, &period.Status, &period.CreatedBy)
	if err != nil {
		return nil, err
	}

	return &period, nil
}

func (r *attendanceRepository) CreateAttendance(ctx context.Context, attendance *entity.Attendance) (*entity.Attendance, error) {
	sqlStatement := `INSERT INTO attendances (user_id, date, clockin_at, payroll_period_id, created_by)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`

	var id uuid.UUID
	err := r.db.QueryRowContext(ctx, sqlStatement,
		attendance.UserID, attendance.Date, attendance.ClockinAt, attendance.PayrollPeriodID, attendance.CreatedBy).
		Scan(&id)

	if err != nil {
		return nil, err
	}

	attendance.ID = id
	return attendance, nil
}
