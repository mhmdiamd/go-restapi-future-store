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

func ProductRouter(r *httprouter.Router, db *sqlx.DB, v *validator.Validate) {
	productRepository := repository.NewProductRepositoryImpl()
	productService := service.NewProductServiceImpl(productRepository, db, v)
  
  // Image 
  imageRepository := repository.NewImageRepositoryImpl()
  imageService := service.NewImageServiceImpl(imageRepository, db, v)

	productController := controller.NewProductController(productService, imageService)

  createController := middleware.IsAuthMiddleware(middleware.IsSellerMiddleware(productController.Create))

	r.GET("/api/v1/products", productController.FindAll)
	r.GET("/api/v1/products/:id", productController.FindById)
	r.POST("/api/v1/products",createController)
	r.PUT("/api/v1/products", productController.Update)
	r.DELETE("/api/v1/products", productController.Delete)
}
