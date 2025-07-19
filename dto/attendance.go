package dto

import (
	"time"

	"github.com/google/uuid"
)

// CreateAttendancePeriodRequest represents the request to create a new attendance period
type CreateAttendancePeriodRequest struct {
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
}

type AttendancePeriodResponse struct {
	ID        uuid.UUID `json:"id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Status    string    `json:"status"`
}

// SubmitAttendanceRequest represents the request to submit an attendance record
type SubmitAttendanceRequest struct {
	Date     time.Time `json:"date" validate:"required"`
	CheckIn  time.Time `json:"check_in" validate:"required"`
	CheckOut time.Time `json:"check_out" validate:"required,gtfield=CheckIn"`
	Notes    string    `json:"notes,omitempty"`
}

// AttendanceResponse represents the attendance data in responses
type AttendanceResponse struct {
	ID         uuid.UUID `json:"id"`
	EmployeeID uint      `json:"employee_id"`
	Date       time.Time `json:"date"`
	CheckIn    time.Time `json:"check_in"`
	CheckOut   time.Time `json:"check_out"`
	Status     string    `json:"status"`
	Notes      string    `json:"notes,omitempty"`
}
