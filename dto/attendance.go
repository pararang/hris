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
	StartDate string    `json:"start_date"`
	EndDate   string    `json:"end_date"`
	Status    string    `json:"status"`
}

// SubmitAttendanceRequest represents the request to submit an attendance record
type SubmitAttendanceRequest struct {
	Date     time.Time `json:"date" validate:"required"`
	CheckIn  time.Time `json:"check_in" validate:"required"`
	CheckOut time.Time `json:"check_out" validate:"required,gtfield=CheckIn"`
	Notes    string    `json:"notes,omitempty"`
}

type ClockinResponse struct {
	ID         uuid.UUID  `json:"id"`
	Date       string     `json:"date"`
	ClockinAt  time.Time  `json:"clockin_at"`
	ClockoutAt *time.Time `json:"clockout_at"`
}
