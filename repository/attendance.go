package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/pararang/hris/domain/repository"
	"github.com/pararang/hris/entity"
)

type attendanceRepository struct {
	*BaseRepository
}

func NewAttendanceRepository(db *sql.DB) repository.AttendanceRepository {
	return &attendanceRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

func (r *attendanceRepository) CreatePayrollPeriod(ctx context.Context, period *entity.PayrollPeriod) (*entity.PayrollPeriod, error) {
	sqlStatement := `INSERT INTO payroll_periods (start_date, end_date, status, created_by)
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	var id uuid.UUID
	err := r.executor(ctx).QueryRowContext(ctx, sqlStatement, period.StartDate, period.EndDate, period.Status, period.CreatedBy).Scan(&id)
	if err != nil {
		return nil, err
	}

	period.ID = id
	return period, nil
}

func (r *attendanceRepository) FindOverlappingPayrollPeriod(ctx context.Context, startDate time.Time) (uuid.UUID, error) {
	sqlStatement := `SELECT id FROM payroll_periods WHERE end_date >= $1 LIMIT 1`

	var id uuid.UUID
	err := r.executor(ctx).QueryRowContext(ctx, sqlStatement, startDate).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *attendanceRepository) FindUserAttendanceByDate(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Attendance, error) {
	sqlStatement := `SELECT id, date, user_id, clockin_at, clockout_at, payroll_period_id, created_by FROM attendances WHERE user_id = $1 AND date = $2 LIMIT 1`

	var attendance entity.Attendance
	err := r.executor(ctx).QueryRowContext(ctx, sqlStatement, userID, date.Format(time.DateOnly)).
		Scan(&attendance.ID, &attendance.Date, &attendance.UserID, &attendance.ClockinAt, &attendance.ClockoutAt, &attendance.PayrollPeriodID, &attendance.CreatedBy)
	if err != nil {
		return nil, err
	}

	return &attendance, nil
}

func (r *attendanceRepository) FindLatestPayrollPeriod(ctx context.Context) (*entity.PayrollPeriod, error) {
	sqlStatement := `SELECT id, start_date, end_date, status, created_by FROM payroll_periods ORDER BY end_date DESC LIMIT 1`

	var period entity.PayrollPeriod
	err := r.executor(ctx).QueryRowContext(ctx, sqlStatement).Scan(&period.ID, &period.StartDate, &period.EndDate, &period.Status, &period.CreatedBy)
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
	err := r.executor(ctx).QueryRowContext(ctx, sqlStatement,
		attendance.UserID, attendance.Date, attendance.ClockinAt, attendance.PayrollPeriodID, attendance.CreatedBy).
		Scan(&id)

	if err != nil {
		return nil, err
	}

	attendance.ID = id
	return attendance, nil
}

func (r *attendanceRepository) UpdateAttendance(ctx context.Context, attendance *entity.Attendance) (*entity.Attendance, error) {
	// currently used only for cloclout
	sqlStatement := `UPDATE attendances
	SET clockout_at = $2, updated_by = $3, updated_at = $4
	WHERE id = $1 RETURNING clockout_at`

	err := r.executor(ctx).QueryRowContext(ctx, sqlStatement, attendance.ID, attendance.ClockoutAt, attendance.UpdatedBy, attendance.UpdatedAt).
		Scan(&attendance.ClockoutAt)

	if err != nil {
		return nil, err
	}

	return attendance, nil
}

func (r *attendanceRepository) GetPayrollPeriodByID(ctx context.Context, id uuid.UUID) (*entity.PayrollPeriod, error) {
	sqlStatement := `SELECT id, start_date, end_date, status FROM payroll_periods WHERE id = $1`

	var period entity.PayrollPeriod
	err := r.executor(ctx).QueryRowContext(ctx, sqlStatement, id).Scan(&period.ID, &period.StartDate, &period.EndDate, &period.Status)
	if err != nil {
		return nil, err
	}

	return &period, nil
}

func (r *attendanceRepository) CountAttendance(ctx context.Context, userID, payrollPeriodID uuid.UUID) (int, error) {
	sqlStatement := `SELECT COUNT(1) FROM attendances WHERE user_id = $1 AND payroll_period_id = $2`

	var count int
	err := r.executor(ctx).QueryRowContext(ctx, sqlStatement, userID, payrollPeriodID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *attendanceRepository) SetStatusPayrollPeriod(ctx context.Context, id uuid.UUID, status, updatedBy string) error {
	sqlStatement := `UPDATE payroll_periods SET status = $2, updated_by = $3, updated_at = $4 WHERE id = $1`

	_, err := r.executor(ctx).ExecContext(ctx, sqlStatement, id, status, updatedBy, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (r *attendanceRepository) GetUserAttendanceListByPeriod(ctx context.Context, userID uuid.UUID, payrollPeriodID uuid.UUID) ([]*entity.Attendance, error) {
	sqlStatement := `SELECT id, date, user_id, payroll_period_id, clockin_at, clockout_at, created_by, created_at, updated_by, updated_at FROM attendances WHERE user_id = $1 AND payroll_period_id = $2`

	rows, err := r.executor(ctx).QueryContext(ctx, sqlStatement, userID, payrollPeriodID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attendances []*entity.Attendance
	for rows.Next() {
		var attendance entity.Attendance
		err := rows.Scan(&attendance.ID, &attendance.Date, &attendance.UserID, &attendance.PayrollPeriodID, &attendance.ClockinAt, &attendance.ClockoutAt, &attendance.CreatedBy, &attendance.CreatedAt, &attendance.UpdatedBy, &attendance.UpdatedAt)
		if err != nil {
			return nil, err
		}
		attendances = append(attendances, &attendance)
	}

	return attendances, nil
}
