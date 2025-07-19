package entity

import (
	"time"
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
	EmployeeID uint      `json:"employee_id"`
	Date       time.Time `json:"date"`
	CheckIn    time.Time `json:"check_in"`
	CheckOut   time.Time `json:"check_out"`
	Status     string    `json:"status"` // e.g., "Present", "Absent", "Late", "Half-day"
	Notes      string    `json:"notes,omitempty"`
}
