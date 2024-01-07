package router

import (
	"github.com/jmoiron/sqlx"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/mhmdiamd/go-restapi-future-store/lib/controller"
	"github.com/mhmdiamd/go-restapi-future-store/lib/repository"
	"github.com/mhmdiamd/go-restapi-future-store/lib/service"
)

func CategoryRouter(r *httprouter.Router, db *sqlx.DB, v *validator.Validate) {

  categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, v)
	categoryController := controller.NewCategoryController(categoryService)

	r.GET("/api/v1/categories", categoryController.FindAll)
	r.GET("/api/v1/categories/:id", categoryController.FindById)
	r.POST("/api/v1/categories", categoryController.Create) 
	r.PUT("/api/v1/categories/:id", categoryController.Update)
	r.DELETE("/api/v1/categories/:id", categoryController.Delete)
}
