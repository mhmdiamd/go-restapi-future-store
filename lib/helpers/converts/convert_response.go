package converts

import (
	"github.com/mhmdiamd/go-restapi-future-store/lib/repository"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/response"
)

func ConvertToCategoryResponse(category repository.Category) response.CategoryResponse {
	return response.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ConvertToProductResponse(product repository.Product) response.ProductResponse {
	return response.ProductResponse{
		Id:           product.Id,
		Product_name: product.Product_name,
		Stock:        product.Stock,
		Price:        product.Price,
		Description:  product.Description,
		Photo:        product.Photo,
		Id_category:  product.Id_category,
	}
}

func ConvertToSliceProductResponse(products []repository.Product) []response.ProductResponse {
	var dataProducts []response.ProductResponse
	for _, Product := range products {
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
