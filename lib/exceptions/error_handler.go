package exceptions

import (
	"net/http"

	"github.com/mhmdiamd/go-restapi-future-store/lib/helpers"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web"
)

func ErrorHandler(w http.ResponseWriter, req *http.Request, err interface{}) {

	if notFoundError(w, req, err) {
		return
	}

	internalServerError(w, req, err)
}

func notFoundError(w http.ResponseWriter, req *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		helpers.WriteToResponse(w, "application/json", webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, req *http.Request, err interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
	}

	helpers.WriteToResponse(w, "application/json", webResponse)

}
