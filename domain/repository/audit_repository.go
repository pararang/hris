package repository

//go:generate mockery --name=AuditRepository --output=. --outpkg=repository --filename=audit_repository_mock.go

import (
	"context"

	"github.com/pararang/hris/entity"
)

// AuditRepository defines the interface for audit log-related database operations
type AuditRepository interface {
	CreateAuditLog(ctx context.Context, log *entity.AuditLog) error
}
