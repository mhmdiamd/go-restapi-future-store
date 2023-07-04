package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type CategoryRepository interface {
	Create(ctx context.Context, tx *sql.Tx, name string) (Category, error)
	Update(ctx context.Context, tx *sql.Tx, category Category) Category
	Delete(ctx context.Context, tx *sql.Tx, categoryId uuid.UUID) error
	FindById(ctx context.Context, tx *sql.Tx, categoryId uuid.UUID) (Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []Category
}
