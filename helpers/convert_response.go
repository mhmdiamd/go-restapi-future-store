package helpers

import (
	"github.com/mhmdiamd/go-restapi-future-store/model/category/web"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

func ConvertToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ConvertToSliceCategoryResponse(categories []domain.Category) []web.CategoryResponse {
	var dataCategories []web.CategoryResponse
	for _, category := range categories {
		dataCategories = append(dataCategories, ConvertToCategoryResponse(category))
	}

	return dataCategories
}
