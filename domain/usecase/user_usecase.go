package usecase

import (
	"context"

	"github.com/pararang/hris/domain/entity"
)

// UserUseCase defines the interface for user-related business logic
type UserUseCase interface {
	Authenticate(ctx context.Context, username, password string) (*entity.User, error)
	// RegisterEmployee(ctx context.Context, employee *entity.User) error
	GetEmployeeProfile(ctx context.Context, id uint) (*entity.User, error)
	UpdateEmployeeProfile(ctx context.Context, employee *entity.User) error
	ListEmployees(ctx context.Context, offset, limit int) ([]*entity.User, error)
}
