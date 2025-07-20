package repository

import (
	"context"
	"database/sql"

	"github.com/prrng/dealls/dbase"
)

type QueryExecutor interface {
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

type BaseRepository struct {
	db *sql.DB
}

func NewBaseRepository(db *sql.DB) *BaseRepository {
	return &BaseRepository{db: db}
}

func (b *BaseRepository) executor(ctx context.Context) QueryExecutor {
	if tx := dbase.GetTxFromContext(ctx); tx != nil {
		return tx
	}
	return b.db
}
