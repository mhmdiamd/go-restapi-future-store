package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mhmdiamd/go-restapi-future-store/app"
	"github.com/mhmdiamd/go-restapi-future-store/exceptions"
	"github.com/mhmdiamd/go-restapi-future-store/src/category/controller"
	"github.com/mhmdiamd/go-restapi-future-store/src/category/repository"
	"github.com/mhmdiamd/go-restapi-future-store/src/category/service"
	"github.com/mhmdiamd/go-restapi-future-store/src/router"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, *validate)
	categoryController := controller.NewCategoryController(categoryService)

	categoryRouter := router.CategoryRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: categoryRouter,
	}

	err := server.ListenAndServe()

	exceptions.PanicIfError(err)

}
