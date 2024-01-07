package middleware

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mhmdiamd/go-restapi-future-store/app"
	"github.com/mhmdiamd/go-restapi-future-store/exception"
	"github.com/mhmdiamd/go-restapi-future-store/helper"
	"github.com/mhmdiamd/go-restapi-future-store/model/web"
)

type AuthMiddleware struct {
	Handler httprouter.Handle
}

func NewAuthMiddleware(handler httprouter.Handle) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func  IsAuthMiddleware(next httprouter.Handle) httprouter.Handle {
  return func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
    if req.Header.Get("X-API-Key") != "" {
	    tokenWithBearer := req.Header.Get("X-API-Key")

      if tokenWithBearer == ""{
        panic(exception.NewErrorUnauthorized("Unauthorized"))
      }

      payload, err := app.ValidateToken(tokenWithBearer)

      if err != nil {
        helper.PanicIfError(err)
      }

      req.Header.Add("id_user", payload.ID.String())
      req.Header.Add("role", payload.Role)
		
      next(w, req, params)
	  } else {
		  w.Header().Set("Content-Type", "application/json")
		  w.WriteHeader(http.StatusUnauthorized)

		  webResponse := web.WebResponse{
			  Code:   http.StatusUnauthorized,
			  Status: "UNAUTHORIZE",
		  }

		  helper.WriteToResponseBody(w, webResponse)
	  }
  }
} 

func (m *AuthMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request, p httprouter.Params) {

	if req.Header.Get("X-API-Key") != "" {
	  tokenWithBearer := req.Header.Get("X-API-Key")

    if tokenWithBearer == ""{
      panic(exception.NewErrorUnauthorized("Unauthorized"))
    }

    payload, err := app.ValidateToken(tokenWithBearer)

    if err != nil {
      helper.PanicIfError(err)
    }

    fmt.Println(payload)

    req.Header.Add("id_user", payload.ID.String())

		m.Handler(w, req, p)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZE",
		}

		helper.WriteToResponseBody(w, webResponse)
	}
}
