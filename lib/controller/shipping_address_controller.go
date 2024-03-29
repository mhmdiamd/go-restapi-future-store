package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ShippingAddressController interface {
  Create(res http.ResponseWriter, req *http.Request, params httprouter.Params) 
  Update(res http.ResponseWriter, req *http.Request, params httprouter.Params) 
  Delete(res http.ResponseWriter, req *http.Request, params httprouter.Params) 
  FindAll(res http.ResponseWriter, req *http.Request, params httprouter.Params) 
  FindAllByIdUser(res http.ResponseWriter, req *http.Request, params httprouter.Params) 
  FindById(res http.ResponseWriter, req *http.Request, params httprouter.Params) 

  Activate(res http.ResponseWriter, req *http.Request, params httprouter.Params) 
}
