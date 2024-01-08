package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/dto"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type ImageRepository interface {
  UploadProductImage(ctx context.Context, tx *sqlx.Tx, data dto.CreateProductBody) (domain.ProductImage, error)
  UpdateProductImage(ctx context.Context, tx *sqlx.Tx, data dto.UpdateProductBody) (domain.ProductImage, error)
 
  GetAllByIdProduct(ctx context.Context, tx *sqlx.Tx, id_product uuid.UUID) []domain.ProductImage
}
