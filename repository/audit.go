package repository

import (
	"context"
	"database/sql"

	"github.com/pararang/hris/domain/entity"
	"github.com/pararang/hris/domain/repository"
)

type auditRepository struct {
	*BaseRepository
}

// NewAuditRepository creates a new instance of AuditRepository
func NewAuditRepository(db *sql.DB) repository.AuditRepository {
	return &auditRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// CreateAuditLog creates a new audit log
func (r *auditRepository) CreateAuditLog(ctx context.Context, log *entity.AuditLog) error {
	// TODO: Implement database query to create audit log
	return nil
}
