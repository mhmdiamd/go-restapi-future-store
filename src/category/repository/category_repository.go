package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type CategoryRepository interface {
	Create(ctx context.Context, tx *sql.Tx, name string) (domain.Category, error)
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, categoryId uuid.UUID) error
	FindById(ctx context.Context, tx *sql.Tx, categoryId uuid.UUID) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
