package router

import (
	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-restapi-future-store/lib/controller"
	"github.com/mhmdiamd/go-restapi-future-store/lib/repository"
	"github.com/mhmdiamd/go-restapi-future-store/lib/service"
	"github.com/mhmdiamd/go-restapi-future-store/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func ImageRouter(r *httprouter.Router, db *sqlx.DB, v *validator.Validate) {
  imageRepository := repository.NewImageRepositoryImpl()
  imageService := service.NewImageServiceImpl(imageRepository, db, v)
  imageController := controller.NewImageControllerImpl(imageService)

  uploadProductImage := middleware.IsAuthMiddleware(middleware.IsSellerMiddleware(imageController.UploadProductImage))
  updateProductImage := middleware.IsAuthMiddleware(middleware.IsSellerMiddleware(imageController.UpdateProductImage))

  r.POST("/api/v1/uploads/product-image", uploadProductImage)
  r.PUT("/api/v1/uploads/product-image/:id", updateProductImage)
}
