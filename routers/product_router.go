package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mhmdiamd/go-restapi-future-store/lib/controller/product"
	"github.com/mhmdiamd/go-restapi-future-store/lib/exceptions"
)

func ProductRouter(productController product.ProductController) *httprouter.Router {
	var router = httprouter.New()

	router.GET("/api/products", productController.FindAll)
	router.GET("/api/products/:productId", productController.FindById)
	router.POST("/api/products", productController.Create)
	router.PUT("/api/products/:productId", productController.Update)
	router.DELETE("/api/products/:productId", productController.Delete)

	router.PanicHandler = exceptions.ErrorHandler

	return router
}
