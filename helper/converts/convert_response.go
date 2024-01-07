package converts

import (
	"fmt"

	"github.com/mhmdiamd/go-restapi-future-store/lib/repository"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/response"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

func ConvertToCategoryResponse(category repository.Category) response.CategoryResponse {
	return response.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ConvertToProductResponse(product domain.Product) response.ProductResponse {
	return response.ProductResponse{
		ID:           product.ID,
		Name: product.Name,
		Stock:        product.Stock,
		Price:        product.Price,
		Description:  product.Description,
		Photo:        product.Photo,
		Id_category:  product.Id_category,
		Id_user:  product.Id_user,
	}
}

func ConvertToSliceProductResponse(products []domain.Product) []response.ProductResponse {
	var dataProducts []response.ProductResponse
	for _, Product := range products {
    fmt.Println(Product)
		dataProducts = append(dataProducts, ConvertToProductResponse(Product))
	}

	return dataProducts
}

func ConvertToSliceCategoryResponse(categories []repository.Category) []response.CategoryResponse {

	var dataCategories []response.CategoryResponse
	for _, category := range categories {
		dataCategories = append(dataCategories, ConvertToCategoryResponse(category))
	}

	return dataCategories
}
