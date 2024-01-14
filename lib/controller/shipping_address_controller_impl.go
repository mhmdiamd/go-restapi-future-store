package controller

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/mhmdiamd/go-restapi-future-store/helper"
	"github.com/mhmdiamd/go-restapi-future-store/lib/service"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/dto"
	"github.com/mhmdiamd/go-restapi-future-store/model/web"
)

type ShippingAddressControllerImpl struct {
  shippingAddressService service.ShippingAddressService
}

func NewShippingAddressControllerImpl(service service.ShippingAddressService) ShippingAddressController {
	return &ShippingAddressControllerImpl{
		shippingAddressService: service,
	}
}

func (c *ShippingAddressControllerImpl) Create(w http.ResponseWriter, req *http.Request, params httprouter.Params) {

  userId := req.Header.Get("id_user")
  parsedUuid := uuid.MustParse(userId)

  // Convert data stream to json
	requestBody := dto.CreateShippingAddressDTO{}
	helper.ReadRequestBody(req, &requestBody)

  requestBody.Id_user = parsedUuid

	// Call servicee
	response := c.shippingAddressService.Create(req.Context(), requestBody)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *ShippingAddressControllerImpl) Update(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// Get Id From Params
	shippindAddressId := params.ByName("id")
	// Convert to uuid from string
	parsedUuid := uuid.MustParse(shippindAddressId)

  // Convert data stream to json
	requestBody := dto.UpdateShippingAddressDTO{}
	helper.ReadRequestBody(req, &requestBody)

  requestBody.Id = parsedUuid

	// Call servicee
	response := c.shippingAddressService.Update(req.Context(), requestBody)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *ShippingAddressControllerImpl) Delete(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// Get Id From Params
	shippindAddressId := params.ByName("id")
	// Convert to uuid from string
	newId := uuid.MustParse(shippindAddressId)

	// Call service
  message := c.shippingAddressService.Delete(req.Context(), newId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
    Data: message,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *ShippingAddressControllerImpl) FindById(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// Get Id From Params
	shippindAddressId := params.ByName("id")

	// Convert to uuid from string
	newId, err := uuid.Parse(shippindAddressId)

	if err != nil {
		helper.PanicIfError(err)
	}

	// Call service
	response := c.shippingAddressService.FindById(req.Context(), newId)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (c *ShippingAddressControllerImpl) FindAll(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	responses := c.shippingAddressService.FindAll(req.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (c *ShippingAddressControllerImpl) FindAllByIdUser(w http.ResponseWriter, req *http.Request, params httprouter.Params){  

  id_seller := req.Header.Get("id_user")

  result := c.shippingAddressService.FindAllByIdUser(req.Context(), uuid.MustParse(id_seller))
 
  var webResponse web.WebResponse
  
  if len(result) < 1 {
    webResponse = web.WebResponse {
      Code : http.StatusOK,
      Status: "Success",
      Data: result,
    }
  }else {
    webResponse = web.WebResponse {
      Code : http.StatusOK,
      Status: "Success",
      Data: result,
    }
  }


  helper.WriteToResponseBody(w, webResponse)

}


func (c *ShippingAddressControllerImpl) Activate(w http.ResponseWriter, req *http.Request, params httprouter.Params) {

  // Convert data stream to json
	requestBody := dto.ActivateShippingAddressDTO{}
	helper.ReadRequestBody(req, &requestBody)

	// Call servicee
	response := c.shippingAddressService.Activate(req.Context(), requestBody.Id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(w, webResponse)
}
