package controller

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mhmdiamd/go-restapi-future-store/helper"
	"github.com/mhmdiamd/go-restapi-future-store/lib/service"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/response"
	"github.com/mhmdiamd/go-restapi-future-store/model/web"
)

type AuthControllerImpl struct {
  Service service.AuthService
}

func NewAuthControllerImpl(authService service.AuthService) AuthController {
  return &AuthControllerImpl{
    Service : authService,
  }
}

func(c *AuthControllerImpl) SignUp(res http.ResponseWriter, req *http.Request, params httprouter.Params){
  // create object request body
  requestBody := response.SignUpRequestDTO{
    Role: "customer",
  }

  // Create Decoder for decoded request body
  err := json.NewDecoder(req.Body).Decode(&requestBody)
  helper.PanicIfError(err)

  data := c.Service.SignUp(req.Context(), requestBody)

  webResponse := web.WebResponse{
    Status: "Success",
    Code: http.StatusOK,
    Data: data,
  }

  helper.WriteToResponseBody(res, webResponse)
}

func(s *AuthControllerImpl) SignIn(res http.ResponseWriter, req *http.Request, params httprouter.Params){
  body := response.SignInRequestDTO{}

  // Decoded data from request Body
  err := json.NewDecoder(req.Body).Decode(&body)
  helper.PanicIfError(err)

  result := s.Service.SignIn(req.Context(), body)

  webResponse := web.WebResponse{
    Status: "success",
    Code: http.StatusOK,
    Data: result,
  }

  helper.WriteToResponseBody(res, webResponse)
}
