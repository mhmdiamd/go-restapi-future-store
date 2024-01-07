package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/mhmdiamd/go-restapi-future-store/app"
	"github.com/mhmdiamd/go-restapi-future-store/exception"
	"github.com/mhmdiamd/go-restapi-future-store/helper"
	myRouter "github.com/mhmdiamd/go-restapi-future-store/router"
)


func main() {

  err := godotenv.Load(".env.local") 
  helper.PanicIfError(err)

  db := app.NewDB()
	validate := validator.New()
	router := httprouter.New()

	//Initialize Category router
	myRouter.AuthRouter(router, db, validate)
	myRouter.UserRouter(router, db, validate)
  myRouter.CategoryRouter(router, db, validate)
	myRouter.ProductRouter(router, db, validate)
	myRouter.SellerRouter(router, db, validate)
	myRouter.ImageRouter(router, db, validate)

	// authMiddleware := middleware.NewAuthMiddleware(router)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Handler: router,
		Addr:    "localhost:8000",
    
	}

	err = server.ListenAndServe()

	helper.PanicIfError(err)
}
