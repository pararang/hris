package dbase

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/prrng/dealls/config"
)

// NewPostgresDB creates a new PostgreSQL database connection
func NewPostgresDB(dbConf config.DB) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConf.Host, dbConf.Port, dbConf.User, dbConf.Password, dbConf.Name, dbConf.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
