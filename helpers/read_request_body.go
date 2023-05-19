package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/mhmdiamd/go-restapi-future-store/exceptions"
)

func ReadRequestBody(req *http.Request, requestBody interface{}) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(requestBody)
	exceptions.PanicIfError(err)
}
