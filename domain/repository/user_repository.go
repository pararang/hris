package repository

import (
	"context"

	"github.com/prrng/dealls/domain/entity"
)

// UserRepository defines the interface for user-related database operations
type UserRepository interface {
	GetEmployeeByEmail(ctx context.Context, email string) (*entity.User, error)
	ListEmployees(ctx context.Context) ([]*entity.User, error) //TODO: enable filter and pagination
}
