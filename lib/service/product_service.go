package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/response"
)

type ProductService interface {
	Create(ctx context.Context, product response.CreateProductRequest) (response.ProductResponse, error)
	Update(ctx context.Context, product response.UpdateProductRequest) (response.ProductResponse, error)
	Delete(ctx context.Context, productId uuid.UUID)
	FindById(ctx context.Context, productId uuid.UUID) response.ProductResponse
	FindAll(ctx context.Context) []response.ProductResponse
}
