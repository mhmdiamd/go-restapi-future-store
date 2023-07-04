package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mhmdiamd/go-restapi-future-store/lib/controller/category"
	"github.com/mhmdiamd/go-restapi-future-store/lib/exceptions"
)

func CategoryRouter(categoryController category.CategoryController) *httprouter.Router {
	var router = httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exceptions.ErrorHandler

	return router
}
