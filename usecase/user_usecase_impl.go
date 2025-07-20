package usecase

import (
	"context"
	"fmt"

	"errors"

	"github.com/prrng/dealls/domain/entity"
	"github.com/prrng/dealls/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo  repository.UserRepository
	auditRepo repository.AuditRepository
}

// NewUserUseCase creates a new instance of UserUseCase
func NewUserUseCase(userRepo repository.UserRepository, auditRepo repository.AuditRepository) *userUseCase {
	return &userUseCase{
		userRepo:  userRepo,
		auditRepo: auditRepo,
	}
}

// AuthenticateEmployee authenticates an employee
func (u *userUseCase) Authenticate(ctx context.Context, email, password string) (*entity.User, error) {
	user, err := u.userRepo.GetEmployeeByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed get employee data: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

// GetEmployeeProfile gets an employee's profile
func (u *userUseCase) GetEmployeeProfile(ctx context.Context, id uint) (*entity.User, error) {
	// TODO: Implement get employee profile logic
	// 1. Get employee by ID
	return nil, nil
}

// UpdateEmployeeProfile updates an employee's profile
func (u *userUseCase) UpdateEmployeeProfile(ctx context.Context, employee *entity.User) error {
	// TODO: Implement update employee profile logic
	// 1. Get existing employee
	// 2. Update fields
	// 3. Save via repository
	return nil
}

// ListEmployees lists all employees
func (u *userUseCase) ListEmployees(ctx context.Context, offset, limit int) ([]*entity.User, error) {
	// TODO: Implement list employees logic
	// 1. Get employees from repository with pagination
	return nil, nil
}
