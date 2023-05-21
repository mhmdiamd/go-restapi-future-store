package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mhmdiamd/go-restapi-future-store/exceptions"
	"github.com/mhmdiamd/go-restapi-future-store/src/category/controller"
)

func CategoryRouter(categoryController controller.CategoryController) *httprouter.Router {
	var router = httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exceptions.ErrorHandler

	return router
}
