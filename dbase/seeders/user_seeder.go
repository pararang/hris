package seeders

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/prrng/dealls/domain/entity"
	"golang.org/x/crypto/bcrypt"
)

func SeedUser(db *sql.DB) error {
	ctx := context.Background()

	// Create a simple repository implementation for seeding
	repo := &userSeeder{db: db}

	baseSalary := func() float64 {
		return 10_000_000.0 + float64(rand.Intn(10)*500_000)
	}

	// Seed admin
	adminPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	baseModel := entity.BaseModel{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: "system-seeder",
	}

	admin := &entity.User{
		Name:       "Sadmin Mariyanto",
		Email:      "admin@comp.id",
		Password:   string(adminPassword),
		IsAdmin:    true,
		BaseSalary: baseSalary(),
		BaseModel:  baseModel,
	}

	adminID, err := repo.CreateEmployee(ctx, admin)
	if err != nil {
		return err
	}

	if err := repo.SetAsAdmin(ctx, adminID); err != nil {
		return err
	}

	for i := 1; i <= 100; i++ {
		password, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("password%d", i)), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		employee := &entity.User{
			Name:       fmt.Sprintf("Employee %d", i),
			Email:      fmt.Sprintf("emp%d@comp.id", i),
			Password:   string(password),
			BaseSalary: baseSalary(),
			BaseModel:  baseModel,
		}

		_, err = repo.CreateEmployee(ctx, employee)
		if err != nil {
			return err
		}
	}

	log.Println("Seed data completed successfully")
	return nil
}

// userSeeder is a simple implementation for seeding
type userSeeder struct {
	db *sql.DB
}

func (s *userSeeder) CreateEmployee(ctx context.Context, user *entity.User) (uint, error) {
	var id uint
	query := `INSERT INTO users (name, email, password, base_salary, is_admin, created_at, updated_at, created_by)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	err := s.db.QueryRowContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.Password,
		user.BaseSalary,
		user.IsAdmin,
		user.CreatedAt,
		user.UpdatedAt,
		user.CreatedBy,
	).Scan(&id)

	return id, err
}

func (s *userSeeder) SetAsAdmin(ctx context.Context, id uint) error {
	query := `UPDATE users SET is_admin = true WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, id)
	return err
}
