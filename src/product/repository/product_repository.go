package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
	webProduct "github.com/mhmdiamd/go-restapi-future-store/model/web/product"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, product webProduct.CreateProductRequest) (domain.Product, error)
	FindById(ctx context.Context, tx *sql.Tx, productId uuid.UUID) (domain.Product, error)
	Update(ctx context.Context, tx *sql.Tx, product domain.Product) (domain.Product, error)
	Delete(ctx context.Context, tx *sql.Tx, productId uuid.UUID)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Product
}
