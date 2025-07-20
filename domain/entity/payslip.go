package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type PayslipDetails struct {
	AttendanceDays    int16   `json:"attendance_days"`
	ProrateBaseSalary float64 `json:"prorate_base_salary"`
	OvertimePay       float64 `json:"overtime_pay"`
	OvertimeHours     int32   `json:"overtime_hours"`
	ReimbursementPay  float64 `json:"reimbursement_pay"`
}

func (pd PayslipDetails) Value() (driver.Value, error) {
	return json.Marshal(pd)
}

func (pd *PayslipDetails) Scan(value any) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &pd)
}

type Payslip struct {
	BaseModel
	UserID              uuid.UUID      `json:"employee_id"`
	PayrollPeriodID     uuid.UUID      `json:"payroll_period_id"`
	BaseSalary          float64        `json:"base_salary"`
	ProratedBaseSalary  float64        `json:"prorated_base_salary"`
	OvertimePay         float64        `json:"overtime_pay"`
	ReimbursementAmount float64        `json:"reimbursement_amount"`
	TakeHomePay         float64        `json:"take_home_pay"`
	DetailsJSON         PayslipDetails `json:"details"`
}
