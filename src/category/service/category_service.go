package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/model/category/web"
)

type CategoryService interface {
	Create(ctx context.Context, req web.CategoryCreateRequest) (web.CategoryResponse, error)
	Update(ctx context.Context, req web.CategoryUpdateRequest) (web.CategoryResponse, error)
	Delete(ctx context.Context, categoryId uuid.UUID)
	FindById(ctx context.Context, categoryId uuid.UUID) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
