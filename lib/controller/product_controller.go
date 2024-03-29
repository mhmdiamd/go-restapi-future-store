package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductController interface {
	Create(w http.ResponseWriter, req *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, req *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, req *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, req *http.Request, params httprouter.Params)
	FindAll(w http.ResponseWriter, req *http.Request, params httprouter.Params)

  // Seller Access
  FindAllByIdSeller(w http.ResponseWriter, req *http.Request, params httprouter.Params)
}
