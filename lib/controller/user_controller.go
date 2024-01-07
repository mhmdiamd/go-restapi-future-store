package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
  Update(w http.ResponseWriter, req *http.Request, p httprouter.Params)
  GetProfile(w http.ResponseWriter, req *http.Request, p httprouter.Params)
}
