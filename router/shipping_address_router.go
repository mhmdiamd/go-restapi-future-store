
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

func ShippingAddressRouter(r *httprouter.Router, db *sqlx.DB, v *validator.Validate) {
	shippingAddressRepository := repository.NewShippingAddressRepositoryImpl()
	shippingAddressService := service.NewShippingAddressRepositoryImpl(shippingAddressRepository, db, v)

	shippingAddressController := controller.NewShippingAddressControllerImpl(shippingAddressService)

  findAll := middleware.IsAuthMiddleware(middleware.IsCustomerMiddleware(shippingAddressController.FindAll))
  findById := middleware.IsAuthMiddleware(middleware.IsCustomerMiddleware(shippingAddressController.FindById))
  findAllByIdUser := middleware.IsAuthMiddleware(middleware.IsCustomerMiddleware(shippingAddressController.FindAllByIdUser))

  create := middleware.IsAuthMiddleware(middleware.IsCustomerMiddleware(shippingAddressController.Create))
  update := middleware.IsAuthMiddleware(middleware.IsCustomerMiddleware(shippingAddressController.Update))
  delete := middleware.IsAuthMiddleware(middleware.IsCustomerMiddleware(shippingAddressController.Delete))

  activate := middleware.IsAuthMiddleware(middleware.IsCustomerMiddleware(shippingAddressController.Activate))

  // Shipping address users
  r.GET("/api/v1/shipping-addresses/users/get-all", findAllByIdUser)

  r.GET("/api/v1/shipping-address", findAll)
  r.GET("/api/v1/shipping-address/:id", findById)
  r.POST("/api/v1/shipping-address", create)
  r.DELETE("/api/v1/shipping-address/:id", delete)
  r.PUT("/api/v1/shipping-address/:id", update)

  r.POST("/api/v1/shipping-address/activate", activate)
}
