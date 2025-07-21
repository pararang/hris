package dto

import (
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
type ItemListPayslipResponse struct {
	ID            uuid.UUID `json:"id"`
	GeneratedDate string    `json:"generated_date"`
	BaseSalary    float64   `json:"base_salary"`
	Details       any       `json:"details"`
}

type PayslipBreakdownRembursement struct {
	TransactionDate string  `json:"transaction_date"`
	Description     string  `json:"description"`
	Amount          float64 `json:"amount"`
}

type PayslipBreakdownAttendance struct {
	Date     string `json:"date"`
	ClockIn  string `json:"clock_in"`
	ClockOut string `json:"clock_out"`
}

type PayslipBreakdownOvertime struct {
	Date  string `json:"date"`
	Hours int16  `json:"hours"`
}

type PayslipBreakdownResponse struct {
	ID            uuid.UUID                      `json:"id"`
	GeneratedDate string                         `json:"generated_date"`
	BaseSalary    float64                        `json:"base_salary"`
	TakeHomePay   float64                        `json:"take_home_pay"`
	Details       any                            `json:"details"`
	Attendances   []PayslipBreakdownAttendance   `json:"attendance"`
	Rembursements []PayslipBreakdownRembursement `json:"rembursements"`
	Overtimes     []PayslipBreakdownOvertime     `json:"overtime"`
}
