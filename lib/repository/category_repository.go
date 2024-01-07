package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/google/uuid"
)

type CategoryRepository interface {
	Create(ctx context.Context, tx *sqlx.Tx, name string) (Category, error)
	Update(ctx context.Context, tx *sqlx.Tx, category Category) Category
	Delete(ctx context.Context, tx *sqlx.Tx, categoryId uuid.UUID) error
	FindById(ctx context.Context, tx *sqlx.Tx, categoryId uuid.UUID) (Category, error)
	FindAll(ctx context.Context, tx *sqlx.Tx) []Category
}
