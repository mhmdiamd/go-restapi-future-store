
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

func CartRouter(r *httprouter.Router, db *sqlx.DB, v *validator.Validate) {
	cartRepository := repository.NewCartRepositoryImpl()
	cartService := service.NewCartServiceImpl(cartRepository, db, v)

  // Product
	productRepository := repository.NewProductRepositoryImpl()
	productService := service.NewProductServiceImpl(productRepository, db, v)

	cartController := controller.NewCartControllerImpl(cartService, productService)

  findAll := middleware.IsAuthMiddleware(middleware.IsCustomerMiddleware(cartController.FindAll))
  create := middleware.IsAuthMiddleware(middleware.IsCustomerMiddleware(cartController.Create))
  delete := middleware.IsAuthMiddleware(middleware.IsCustomerMiddleware(cartController.Delete))
  increment := middleware.IsAuthMiddleware(middleware.IsCustomerMiddleware(cartController.Increment))
  decrement := middleware.IsAuthMiddleware(middleware.IsCustomerMiddleware(cartController.Decrement))

  r.GET("/api/v1/carts", findAll)
  r.POST("/api/v1/carts", create)
  r.DELETE("/api/v1/carts/:id", delete)
  r.POST("/api/v1/carts/increment", increment)
  r.POST("/api/v1/carts/decrement", decrement)

}
