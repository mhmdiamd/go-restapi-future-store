package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CartController interface {
  Create(w http.ResponseWriter, req *http.Request, params httprouter.Params)
  Delete(w http.ResponseWriter, req *http.Request, params httprouter.Params)
  FindAll(w http.ResponseWriter, req *http.Request, params httprouter.Params)
  Increment(w http.ResponseWriter, req *http.Request, params httprouter.Params)
  Decrement(w http.ResponseWriter, req *http.Request, params httprouter.Params)
}
