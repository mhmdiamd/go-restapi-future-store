package helper

import (
	"encoding/json"
	"net/http"

	"github.com/mhmdiamd/go-restapi-future-store/model/web"
)

func ReadRequestBody(req *http.Request, requestBody interface{}) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(requestBody)

	if err != nil {
		panic(err)
	}
}

func WriteToResponseBody(w http.ResponseWriter, response web.WebResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)
  encoder := json.NewEncoder(w)
	err := encoder.Encode(response)

	PanicIfError(err)
}
