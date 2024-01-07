package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/dto"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type ImageService interface {
  UploadProductImage(ctx context.Context, body dto.CreateProductImageDto) (domain.ProductImage) 
  GetAllByIdProduct(ctx context.Context, id_product uuid.UUID) []domain.ProductImage
}
