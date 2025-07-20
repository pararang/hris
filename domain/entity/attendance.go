package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	PayrollPeriodStatusPending    = "pending"
	PayrollPeriodStatusProcessing = "processing"
	PayrollPeriodStatusCompleted  = "completed"
)

// AttendancePeriod represents a period for which attendance is tracked and payroll is processed
type PayrollPeriod struct {
	BaseModel
	Status    string    `json:"status"` // e.g., "pending", "processing", "completed"
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

// Attendance represents an employee's attendance record for a specific date
type Attendance struct {
	BaseModel
	UserID          uuid.UUID  `json:"user_id"`
	Date            time.Time  `json:"date"`
	ClockinAt       time.Time  `json:"clockin_at"`
	ClockoutAt      *time.Time `json:"clockout_at"`
	PayrollPeriodID uuid.UUID  `json:"payroll_period_id"`
}
