package helpers

import (
	"encoding/json"
	"net/http"
)

func ReadRequestBody(req *http.Request, requestBody interface{}) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(requestBody)

	if err != nil {
		panic(err)
	}
}

func WriteToResponse(w http.ResponseWriter, contentType string, dataResponse interface{}) {
	w.Header().Add("Content-Type", contentType)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(dataResponse)
	PanicIfError(err)
}
