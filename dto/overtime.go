package dto

import (
	"time"

	"github.com/google/uuid"
)

// SubmitOvertimeRequest represents the request to submit an overtime record
type SubmitOvertimeRequest struct {
	Date       string `json:"date" binding:"required,datetime=2006-01-02"`
	HoursTaken uint8  `json:"hours_taken" binding:"required,number,min=1,max=3"`
	Reason     string `json:"reason" validate:"required"`
}

// OvertimeResponse represents the overtime data in responses
type OvertimeResponse struct {
	ID         uuid.UUID `json:"id"`
	Date       time.Time `json:"date"`
	HoursTaken uint8     `json:"hours_taken"`
	Reason     string    `json:"reason"`
	Status     string    `json:"status"`
}
