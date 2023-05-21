package helpers

import (
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
	web "github.com/mhmdiamd/go-restapi-future-store/model/web/category"
	webProduct "github.com/mhmdiamd/go-restapi-future-store/model/web/product"
)

func ConvertToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ConvertToProductResponse(product domain.Product) webProduct.ProductResponse {
	return webProduct.ProductResponse{
		Id:           product.Id,
		Product_name: product.Product_name,
		Stock:        product.Stock,
		Price:        product.Price,
		Description:  product.Description,
		Photo:        product.Photo,
		Id_category:  product.Id_category,
	}
}

func ConvertToSliceProductResponse(products []domain.Product) []webProduct.ProductResponse {
	var dataProducts []webProduct.ProductResponse
	for _, Product := range products {
		dataProducts = append(dataProducts, ConvertToProductResponse(Product))
	}

	return dataProducts
}

func ConvertToSliceCategoryResponse(categories []domain.Category) []web.CategoryResponse {
	var dataCategories []web.CategoryResponse
	for _, category := range categories {
		dataCategories = append(dataCategories, ConvertToCategoryResponse(category))
	}

	return dataCategories
}
