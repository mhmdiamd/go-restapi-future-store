package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthController interface {
  SignUp(res http.ResponseWriter, req *http.Request, params httprouter.Params) 
  SignIn(res http.ResponseWriter, req *http.Request, params httprouter.Params) 
}
