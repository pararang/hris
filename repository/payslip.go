package repository

import (
	"context"
	"database/sql"

	"github.com/pararang/hris/domain/repository"
	"github.com/pararang/hris/entity"
)

type payslipRepository struct {
	*BaseRepository
}

// NewPayslipRepository creates a new instance of PayslipRepository
func NewPayslipRepository(db *sql.DB) repository.PayslipRepository {
	return &payslipRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

func (r *payslipRepository) CreatePayslip(ctx context.Context, payslip *entity.Payslip) error {

	sqlInsert := `INSERT INTO payslips (
					user_id, payroll_period_id,
					base_salary, prorated_base_salary,
					overtime_pay, reimbursement_amount,
					take_home_pay, details_json, created_by)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	err := r.executor(ctx).QueryRowContext(ctx, sqlInsert,
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
