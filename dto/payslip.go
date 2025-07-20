package dto

import (
	"time"

	"github.com/google/uuid"
)

type ProcessPayrollRequest struct {
	PayrollPeriodID uuid.UUID `json:"payroll_period_id" binding:"required"`
}

type PayslipComponentResponse struct {
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
}

// PayslipResponse represents the payslip data in responses
type PayslipResponse struct {
	ID                  uuid.UUID                  `json:"id"`
	EmployeeID          uint                       `json:"employee_id"`
	AttendancePeriodID  uint                       `json:"attendance_period_id"`
	GeneratedDate       time.Time                  `json:"generated_date"`
	BaseSalary          float64                    `json:"base_salary"`
	AttendanceEffect    float64                    `json:"attendance_effect"`
	OvertimeAmount      float64                    `json:"overtime_amount"`
	ReimbursementAmount float64                    `json:"reimbursement_amount"`
	Deductions          float64                    `json:"deductions"`
	TaxAmount           float64                    `json:"tax_amount"`
	TotalTakeHomePay    float64                    `json:"total_take_home_pay"`
	Components          []PayslipComponentResponse `json:"components"`
	Status              string                     `json:"status"`
}
