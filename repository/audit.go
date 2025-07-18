package repository

import (
	"context"
	"database/sql"

	"github.com/prrng/dealls/domain/entity"
	"github.com/prrng/dealls/domain/repository"
)

type auditRepository struct {
	db *sql.DB
}

// NewAuditRepository creates a new instance of AuditRepository
func NewAuditRepository(db *sql.DB) repository.AuditRepository {
	return &auditRepository{
		db: db,
	}
}

// CreateAuditLog creates a new audit log
func (r *auditRepository) CreateAuditLog(ctx context.Context, log *entity.AuditLog) error {
	// TODO: Implement database query to create audit log
	return nil
}
