package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/mhmdiamd/go-restapi-future-store/exceptions"
)

func WriteToResponse(w http.ResponseWriter, contentType string, dataResponse interface{}) {
	w.Header().Add("Content-Type", contentType)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(dataResponse)
	exceptions.PanicIfError(err)
}
