package router

import (
	"github.com/jmoiron/sqlx"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/mhmdiamd/go-restapi-future-store/lib/controller"
	"github.com/mhmdiamd/go-restapi-future-store/lib/repository"
	"github.com/mhmdiamd/go-restapi-future-store/lib/service"
)

func AuthRouter(r *httprouter.Router, db *sqlx.DB, v *validator.Validate) {
	authRepository := repository.NewAuthRepositoryImpl()
	authService := service.NewAuthServiceImpl(authRepository, db, v)
	authControlller := controller.NewAuthControllerImpl(authService)

	r.POST("/api/v1/signup", authControlller.SignUp)
	r.POST("/api/v1/signin", authControlller.SignIn)
}
