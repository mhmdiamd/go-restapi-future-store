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
	productRepository := repository.NewProductRepositoryImpl()
	productService := service.NewProductServiceImpl(productRepository, db, v)

  // Image 
  imageRepository := repository.NewImageRepositoryImpl()
  imageService := service.NewImageServiceImpl(imageRepository, db, v)

  // Seller
  sellerRepository := repository.NewSellerRepositoryImpl()
  sellerService := service.NewSellerServiceImpl(sellerRepository, db, v)
	sellerController := controller.NewSellerControllerImpl(sellerService)

  sellerUpdate := middleware.IsAuthMiddleware(middleware.IsSellerMiddleware(sellerController.Update))

	productController := controller.NewProductController(productService, imageService)

  uploadPhoto := middleware.IsAuthMiddleware(middleware.IsSellerMiddleware(sellerController.UploadPhoto))

  r.PUT("/api/v1/sellers/:id", sellerUpdate)

  // Seller Product
  r.GET("/api/v1/sellers/users/products", middleware.IsAuthMiddleware(middleware.IsSellerMiddleware(productController.FindAllByIdSeller)))
  
  // Seller Upload
  r.POST("/api/v1/seller/users/upload",uploadPhoto)
}
