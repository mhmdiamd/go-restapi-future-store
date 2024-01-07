package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ImageController interface {
  UploadProductImage(w http.ResponseWriter, req *http.Request, params httprouter.Params)
}
