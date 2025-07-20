package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/prrng/dealls/domain/entity"
	"github.com/prrng/dealls/domain/repository"
)

type reimbursementRepository struct {
	db *sql.DB
}

func NewReimbursementRepository(db *sql.DB) repository.ReimbursementRepository {
	return &reimbursementRepository{
		db: db,
	}
}

func (r *reimbursementRepository) CreateReimbursement(ctx context.Context, reimbursement *entity.Reimbursement) (*entity.Reimbursement, error) {
	sqlStatement := `INSERT INTO reimbursements (user_id, amount, description, transaction_date, payroll_period_id, status, created_by)
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id
	`
	err := r.db.QueryRowContext(ctx, sqlStatement,
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
	err := r.db.QueryRowContext(ctx, sqlStatement, userID, payrollPeriodID).Scan(&totalAmount)
	if err != nil {
		return 0, err
	}
	return totalAmount.Float64, nil
}
