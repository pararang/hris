package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/pararang/hris/domain/repository"
	"github.com/pararang/hris/entity"
)

type reimbursementRepository struct {
	*BaseRepository
}

func NewReimbursementRepository(db *sql.DB) repository.ReimbursementRepository {
	return &reimbursementRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

func (r *reimbursementRepository) CreateReimbursement(ctx context.Context, reimbursement *entity.Reimbursement) (*entity.Reimbursement, error) {
	sqlStatement := `INSERT INTO reimbursements (user_id, amount, description, transaction_date, payroll_period_id, status, created_by)
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id
	`
	err := r.executor(ctx).QueryRowContext(ctx, sqlStatement,
		reimbursement.UserID, reimbursement.Amount,
		reimbursement.Description, reimbursement.TransactionDate,
		reimbursement.PayrollPeriodID, reimbursement.Status, reimbursement.CreatedBy).
		Scan(&reimbursement.ID)
	if err != nil {
		return nil, err
	}
	return reimbursement, nil
}

func (r *reimbursementRepository) CountUserApprovedAmountReimbursementByPeriod(ctx context.Context, userID, payrollPeriodID uuid.UUID) (float64, error) {
	sqlStatement := `SELECT SUM(amount) FROM reimbursements WHERE user_id = $1 AND payroll_period_id = $2` // AND status = $3` ignored for demo purpose
	var totalAmount sql.NullFloat64
	err := r.executor(ctx).QueryRowContext(ctx, sqlStatement, userID, payrollPeriodID).Scan(&totalAmount)
	if err != nil {
		return 0, err
	}
	return totalAmount.Float64, nil
}

func (r *reimbursementRepository) GetUserReimbursementListByPeriod(ctx context.Context, userID uuid.UUID, payrollPeriodID uuid.UUID) ([]*entity.Reimbursement, error) {
	sqlStatement := `SELECT id, user_id, amount, description, transaction_date, payroll_period_id, status, created_by, created_at, updated_at FROM reimbursements WHERE user_id = $1 AND payroll_period_id = $2`
	rows, err := r.executor(ctx).QueryContext(ctx, sqlStatement, userID, payrollPeriodID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reimbursements []*entity.Reimbursement
	for rows.Next() {
		reimbursement := &entity.Reimbursement{}
		err := rows.Scan(&reimbursement.ID, &reimbursement.UserID, &reimbursement.Amount, &reimbursement.Description, &reimbursement.TransactionDate, &reimbursement.PayrollPeriodID, &reimbursement.Status, &reimbursement.CreatedBy, &reimbursement.CreatedAt, &reimbursement.UpdatedAt)
		if err != nil {
			return nil, err
		}
		reimbursements = append(reimbursements, reimbursement)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return reimbursements, nil
}
