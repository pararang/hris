package repository

import (
	"context"
	"database/sql"

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
