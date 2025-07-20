package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	StatusReimbursementPending  = "pending"
	StatusReimbursementApproved = "approved"
	StatusReimbursementRejected = "rejected"
	StatusReimbursementPaid     = "paid"
)

// Reimbursement represents a reimbursement request from an employee
type Reimbursement struct {
	BaseModel
	UserID          uuid.UUID `json:"user_id"`
	Amount          float64   `json:"amount"`
	Description     string    `json:"description"`
	TransactionDate time.Time `json:"transaction_date"`
	PayrollPeriodID uuid.UUID `json:"payroll_period_id"`
	Status          string    `json:"status"` // e.g., "pending", "approved", "rejected", "paid"
}
