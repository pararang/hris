package repository

import (
	"context"
	"database/sql"

	"github.com/pararang/hris/domain/repository"
	"github.com/pararang/hris/entity"
)

type userRepository struct {
	*BaseRepository
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

func (r *userRepository) GetEmployeeByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := r.executor(ctx).QueryRowContext(ctx, "SELECT id, name, email, password, is_admin FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.IsAdmin)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// ListEmployees returns a list of employees with their details.
// TODO: Implement pagination and filtering.
func (r *userRepository) ListEmployees(ctx context.Context) ([]*entity.User, error) {
	rows, err := r.executor(ctx).QueryContext(ctx, "SELECT id, name, email, base_salary FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.BaseSalary); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
