package exception

import (
	"fmt"
	"net/http"

	"github.com/mhmdiamd/go-restapi-future-store/helper"
	"github.com/mhmdiamd/go-restapi-future-store/model/web"
)

func ErrorHandler(w http.ResponseWriter, req *http.Request, err interface{}) {

	if unauthorizedError(w, req, err) {
		return
	}

	if notFoundError(w, req, err) {
		return
	}

  if forbiddenError(w, req, err){
    return 
  }

	internalServerError(w, req, err)
}

func notFoundError(w http.ResponseWriter, req *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception,
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
	}
}

func unauthorizedError(w http.ResponseWriter, req *http.Request, err interface{}) bool {
	_, ok := err.(ErrorUnauthorized)

	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Status: "Unauthorized",
			Code:   http.StatusUnauthorized,
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	} else {
		return false
  }
}


func forbiddenError(w http.ResponseWriter, r *http.Request, err interface{}) bool {

  exception, ok := err.(ErrorForbidden)

  if ok  {

  w.Header().Set("Content-Type", "application/json")

  webResponse := web.WebResponse {
    Code: http.StatusForbidden,
    Status: "Forbidden",
    Data : exception,
  }

  helper.WriteToResponseBody(w, webResponse)
return true  
}else {
    return false
  }
}

func internalServerError(w http.ResponseWriter, req *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

  fmt.Println(err)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
  }

	helper.WriteToResponseBody(w, webResponse)

}
