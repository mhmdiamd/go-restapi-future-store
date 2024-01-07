package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/dto"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/response"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type ProductService interface {
	Create(ctx context.Context, product dto.ProductCreateDTO) response.ProductCreateResponse
	Update(ctx context.Context, product response.UpdateProductRequest) domain.Product
  Delete(ctx context.Context, productId uuid.UUID)
	FindById(ctx context.Context, productId uuid.UUID) domain.Product
	FindAll(ctx context.Context) []response.ProductResponse

  // Seller Access
  FindAllByIdSeller(ctx context.Context, id_seller uuid.UUID) []response.ProductResponse
}
