package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/mhmdiamd/go-restapi-future-store/helper"
	"github.com/mhmdiamd/go-restapi-future-store/lib/service"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/dto"
	"github.com/mhmdiamd/go-restapi-future-store/model/web"
)

type UserControllerImpl struct {
  service service.UserService
}

func NewUserControllerImpl(service service.UserService) UserController {
  return &UserControllerImpl{
    service: service,
  }
}

func (c *UserControllerImpl) Update(w http.ResponseWriter, req *http.Request, p httprouter.Params) {

  userId := p.ByName("id")
  newId := uuid.MustParse(userId)

  var requestBody dto.UserUpdateDTO
  err := json.NewDecoder(req.Body).Decode(&requestBody)

  requestBody.ID = newId

  helper.PanicIfError(err)

  res, err := c.service.Update(req.Context(), requestBody)

  if err != nil {
    helper.PanicIfError(err)
  }

  webResponse := web.WebResponse{
    Code: http.StatusOK,
    Status: "success",
    Data: res,
  }

  helper.WriteToResponseBody(w, webResponse)
}

func (c *UserControllerImpl) GetProfile(w http.ResponseWriter, req *http.Request, p httprouter.Params){
  
  userId := p.ByName("id")
  fmt.Println(userId)

  newId := uuid.MustParse(userId)

  res, err := c.service.GetProfile(req.Context(), newId)

  if err != nil {
    helper.PanicIfError(err)
  }

  webResponse := web.WebResponse{
    Code: http.StatusOK,
    Status: "success",
    Data: res,
  }

  helper.WriteToResponseBody(w, webResponse)
}



