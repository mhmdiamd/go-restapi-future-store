package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/response"
)

type CategoryService interface {
	Create(ctx context.Context, req response.CategoryCreateRequest) (response.CategoryResponse, error)
	Update(ctx context.Context, req response.CategoryUpdateRequest) (response.CategoryResponse, error)
	Delete(ctx context.Context, categoryId uuid.UUID)
	FindById(ctx context.Context, categoryId uuid.UUID) response.CategoryResponse
	FindAll(ctx context.Context) []response.CategoryResponse
}
