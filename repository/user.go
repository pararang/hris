package repository

import (
	"context"
	"database/sql"

	"github.com/prrng/dealls/domain/entity"
	"github.com/prrng/dealls/domain/repository"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetEmployeeByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := r.db.QueryRowContext(ctx, "SELECT id, name, email, password, is_admin FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.IsAdmin)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateEmployee creates a new employee
func (r *userRepository) CreateEmployee(ctx context.Context, employee *entity.User) (id uint, err error) {
	// TODO: Implement database query to create employee
	return 0, nil
}

func (r *userRepository) SetAsAdmin(ctx context.Context, id uint) error {
	// TODO: Implement database query to create user
	return nil
}

// GetEmployeeByID gets an employee by ID
func (r *userRepository) GetEmployeeByID(ctx context.Context, id uint) (*entity.User, error) {
	// TODO: Implement database query to get employee by ID
	return nil, nil
}

// UpdateEmployee updates an employee
func (r *userRepository) UpdateEmployee(ctx context.Context, employee *entity.User) error {
	// TODO: Implement database query to update employee
	return nil
}

// DeleteEmployee deletes an employee
func (r *userRepository) DeleteEmployee(ctx context.Context, id uint) error {
	// TODO: Implement database query to delete employee
	return nil
}

// ListEmployees lists all employees
func (r *userRepository) ListEmployees(ctx context.Context, offset, limit int) ([]*entity.User, error) {
	// TODO: Implement database query to list employees with pagination
	return nil, nil
}
