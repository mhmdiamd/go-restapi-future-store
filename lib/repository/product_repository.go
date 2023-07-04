package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/response"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, product response.ProductCreateRequest) (Product, error)
	FindById(ctx context.Context, tx *sql.Tx, productId uuid.UUID) (Product, error)
	Update(ctx context.Context, tx *sql.Tx, product Product) (Product, error)
	Delete(ctx context.Context, tx *sql.Tx, productId uuid.UUID)
	FindAll(ctx context.Context, tx *sql.Tx) []Product
}
