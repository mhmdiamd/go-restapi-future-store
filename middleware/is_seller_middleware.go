package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mhmdiamd/go-restapi-future-store/exception"
)

func IsSellerMiddleware(next httprouter.Handle) httprouter.Handle {
  return func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {

    role := req.Header.Get("role")

    if role != "seller" {
      panic(exception.NewForbiddenError("Roles are not permitted"))
    }

    next(w, req, params)
  }
}

