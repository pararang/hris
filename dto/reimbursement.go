package dto

import (
	"github.com/google/uuid"
)

// SubmitReimbursementRequest represents the request to submit a reimbursement record
type SubmitReimbursementRequest struct {
	Amount          int    `json:"amount" binding:"required,number,gt=0"`
	Description     string `json:"description" binding:"required"`
	TransactionDate string `json:"transaction_date" binding:"required,datetime=2006-01-02"`
}

// ReimbursementResponse represents the reimbursement data in responses
type ReimbursementResponse struct {
	ID              uuid.UUID `json:"id"`
	Amount          int       `json:"amount"`
	Description     string    `json:"description"`
	TransactionDate string    `json:"transaction_date"`
	Status          string    `json:"status"`
}
