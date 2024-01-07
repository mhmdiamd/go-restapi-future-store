package router

import (
	"github.com/jmoiron/sqlx"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/mhmdiamd/go-restapi-future-store/lib/controller"
	"github.com/mhmdiamd/go-restapi-future-store/lib/repository"
	"github.com/mhmdiamd/go-restapi-future-store/lib/service"
	"github.com/mhmdiamd/go-restapi-future-store/middleware"
)

func SellerRouter(r *httprouter.Router, db *sqlx.DB, v *validator.Validate) {
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, v)

  // Image 
  imageRepository := repository.NewImageRepositoryImpl()
  imageService := service.NewImageServiceImpl(imageRepository, db, v)

	productController := controller.NewProductController(productService, imageService)

  r.GET("/api/v1/users/seller/products", middleware.IsAuthMiddleware(middleware.IsSellerMiddleware(productController.FindAllByIdSeller)))
}
