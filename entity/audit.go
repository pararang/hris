package entity

import (
	"time"
)

// AuditLog represents a record of user actions in the system
type AuditLog struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	IPAddress string    `json:"ip_address"`
	Action    string    `json:"action"`
	Resource  string    `json:"resource"`
	ResourceID uint      `json:"resource_id,omitempty"`
	OldValue  string    `json:"old_value,omitempty"`
	NewValue  string    `json:"new_value,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}
