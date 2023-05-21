package service

import (
	"context"

	"github.com/google/uuid"
	webProduct "github.com/mhmdiamd/go-restapi-future-store/model/web/product"
)

type ProductService interface {
	Create(ctx context.Context, product webProduct.CreateProductRequest) (webProduct.ProductResponse, error)
	Update(ctx context.Context, product webProduct.UpdateProductRequest) (webProduct.ProductResponse, error)
	Delete(ctx context.Context, productId uuid.UUID)
	FindById(ctx context.Context, productId uuid.UUID) webProduct.ProductResponse
	FindAll(ctx context.Context) []webProduct.ProductResponse
}
