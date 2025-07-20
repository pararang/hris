package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	StatusOvertimePending  = "pending"
	StatusOvertimeApproved = "approved"
	StatusOvertimeRejected = "rejected"
)

// Overtime represents an overtime record for an employee
type Overtime struct {
	BaseModel
	UserID          uuid.UUID `json:"user_id"`
	Date            time.Time `json:"date"`
	HoursTaken      uint8     `json:"hours_taken"`
	PayrollPeriodID uuid.UUID `json:"payroll_period_id"`
	Status          string    `json:"status"` // e.g., "pending", "approved", "rejected"
	Reason          string    `json:"reason"`
}
