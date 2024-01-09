package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/dto"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type ImageService interface {
  FindProductImageById(ctx context.Context, Id_product_image uuid.UUID) (domain.ProductImage) 
  UploadProductImage(ctx context.Context, body dto.CreateProductImageDto) (domain.ProductImage) 
  UpdateProductImage(ctx context.Context, body dto.UpdateProductImageDto) (domain.ProductImage) 
  DeleteProductImage(ctx context.Context, Id_product_image uuid.UUID) string

  GetAllByIdProduct(ctx context.Context, id_product uuid.UUID) []domain.ProductImage
}
