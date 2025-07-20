package repository

import (
	"context"
	"database/sql"

	"github.com/prrng/dealls/domain/entity"
	"github.com/prrng/dealls/domain/repository"
)

type payslipRepository struct {
	db *sql.DB
}

// NewPayslipRepository creates a new instance of PayslipRepository
func NewPayslipRepository(db *sql.DB) repository.PayslipRepository {
	return &payslipRepository{
		db: db,
	}
}

func (r *payslipRepository) CreatePayslip(ctx context.Context, payslip *entity.Payslip) error {

	sqlInsert := `INSERT INTO payslips (
					user_id, payroll_period_id,
					base_salary, prorated_base_salary,
					overtime_pay, reimbursement_amount,
					take_home_pay, details_json, created_by)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	err := r.db.QueryRowContext(ctx, sqlInsert,
		payslip.UserID, payslip.PayrollPeriodID,
		payslip.BaseSalary, payslip.ProratedBaseSalary,
		payslip.OvertimePay, payslip.ReimbursementAmount,
		payslip.TakeHomePay, payslip.DetailsJSON, payslip.CreatedBy).
		Scan(&payslip.ID)
	if err != nil {
		return err
	}

	return nil
}
