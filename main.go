package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/mhmdiamd/go-restapi-future-store/app"
	"github.com/mhmdiamd/go-restapi-future-store/helpers"
	"github.com/mhmdiamd/go-restapi-future-store/src/category/controller"
	"github.com/mhmdiamd/go-restapi-future-store/src/category/repository"
	"github.com/mhmdiamd/go-restapi-future-store/src/category/service"
	pController "github.com/mhmdiamd/go-restapi-future-store/src/product/controller"
	pRepository "github.com/mhmdiamd/go-restapi-future-store/src/product/repository"
	pService "github.com/mhmdiamd/go-restapi-future-store/src/product/service"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	var router = httprouter.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, *validate)
	categoryController := controller.NewCategoryController(categoryService)

	productRepository := pRepository.NewProductRepository()
	productService := pService.NewProductService(productRepository, db, *validate)
	productController := pController.NewProductController(productService)

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.GET("/api/products", productController.FindAll)
	router.GET("/api/products/:productId", productController.FindById)
	router.POST("/api/products", productController.Create)
	router.PUT("/api/products/:productId", productController.Update)
	router.DELETE("/api/products/:productId", productController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()

	helpers.PanicIfError(err)

}
