package usecase

import (
	"context"
)

// AuditUseCase defines the interface for audit log-related business logic
type AuditUseCase interface {
	LogAction(ctx context.Context, userID uint, ipAddress, action, resource string, resourceID uint, oldValue, newValue string) error
	// GetAuditLogByID(ctx context.Context, id uint) (*entity.AuditLog, error)
	// ListAuditLogs(ctx context.Context, offset, limit int) ([]*entity.AuditLog, error)
	// ListAuditLogsByUser(ctx context.Context, userID uint, startTime, endTime time.Time) ([]*entity.AuditLog, error)
	// ListAuditLogsByResource(ctx context.Context, resource string, resourceID uint) ([]*entity.AuditLog, error)
}
