//go:build wireinject
// +build wireinject

package main

import (
  "github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"net/http"

  "github.com/mhmdiamd/go-restapi-future-store/middleware"
  "github.com/mhmdiamd/go-restapi-future-store/lib/repository"
  "github.com/mhmdiamd/go-restapi-future-store/lib/controller"
  "github.com/mhmdiamd/go-restapi-future-store/lib/service"
  "github.com/mhmdiamd/go-restapi-future-store/app"
)

var categorySet = wire.NewSet(
  repository.NewCategoryRepository,
  service.NewCategoryService,
  controller.NewCategoryController,
)

func InitializedServer() *http.Server{
  wire.Build(
    app.NewDB,
    validator.New,
    categorySet,
    app.NewRouter,
    wire.Bind(new(http.Handler), new(*httprouter.Router)),
    middleware.NewAuthMiddleware,
    NewServer,
  )

  return nil
} 
