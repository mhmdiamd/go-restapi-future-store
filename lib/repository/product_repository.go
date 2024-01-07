package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/dto"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sqlx.Tx, product dto.ProductCreateDTO) (domain.Product, error)
	FindById(ctx context.Context, tx *sqlx.Tx, productId uuid.UUID) (domain.Product, error)
	Update(ctx context.Context, tx *sqlx.Tx, product Product) domain.Product
	Delete(ctx context.Context, tx *sqlx.Tx, productId uuid.UUID)
	FindAll(ctx context.Context, tx *sqlx.Tx) []domain.Product

  // Seller service
  FindAllByIdSeller(ctx context.Context, tx *sqlx.Tx, id_seller uuid.UUID) []domain.Product
}
