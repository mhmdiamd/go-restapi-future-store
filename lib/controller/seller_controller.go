package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type SellerController interface {
  Update(w http.ResponseWriter, req *http.Request, params httprouter.Params)

  UploadPhoto(w http.ResponseWriter, req *http.Request, params httprouter.Params)
} 
