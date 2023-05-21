package service

import (
	"context"

	"github.com/google/uuid"
	categoryWeb "github.com/mhmdiamd/go-restapi-future-store/model/web/category"
)

type CategoryService interface {
	Create(ctx context.Context, req categoryWeb.CategoryCreateRequest) (categoryWeb.CategoryResponse, error)
	Update(ctx context.Context, req categoryWeb.CategoryUpdateRequest) (categoryWeb.CategoryResponse, error)
	Delete(ctx context.Context, categoryId uuid.UUID)
	FindById(ctx context.Context, categoryId uuid.UUID) categoryWeb.CategoryResponse
	FindAll(ctx context.Context) []categoryWeb.CategoryResponse
}
