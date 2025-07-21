package repository

//go:generate mockery --name=UserRepository --output=. --outpkg=repository --filename=user_repository_mock.go

import (
	"context"

	"github.com/pararang/hris/entity"
)

// UserRepository defines the interface for user-related database operations
type UserRepository interface {
	GetEmployeeByEmail(ctx context.Context, email string) (*entity.User, error)
	ListEmployees(ctx context.Context) ([]*entity.User, error) //TODO: enable filter and pagination
}
