package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mhmdiamd/go-restapi-future-store/exception"
)

func NewRouter(router *httprouter.Router, ) *httprouter.Router {
  router.PanicHandler = exception.ErrorHandler
  return router
}

