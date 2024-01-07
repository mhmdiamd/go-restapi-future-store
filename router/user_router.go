package router

import (
	"github.com/jmoiron/sqlx"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/mhmdiamd/go-restapi-future-store/lib/controller"
	"github.com/mhmdiamd/go-restapi-future-store/lib/repository"
	"github.com/mhmdiamd/go-restapi-future-store/lib/service"
)

func UserRouter(r *httprouter.Router, db *sqlx.DB, v *validator.Validate) {
	userRepository := repository.NewUserRepositoryImpl()
	userService := service.NewUserServiceImpl(userRepository, db, v)
	userController := controller.NewUserControllerImpl(userService)

  r.PUT("/api/v1/user/:id", userController.Update)
  r.GET("/api/v1/user/:id", userController.GetProfile)
}
