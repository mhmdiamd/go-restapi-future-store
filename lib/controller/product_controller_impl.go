package controller

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/mhmdiamd/go-restapi-future-store/helper"
	"github.com/mhmdiamd/go-restapi-future-store/lib/service"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/dto"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/response"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
	"github.com/mhmdiamd/go-restapi-future-store/model/web"
)

type ProductControllerImpl struct {
	productService service.ProductService
  imageService service.ImageService
}

func NewProductController(service service.ProductService, imageService service.ImageService) ProductController {
	return &ProductControllerImpl{
		productService: service,
    imageService: imageService,
	}
}

func (controller *ProductControllerImpl) Create(w http.ResponseWriter, req *http.Request, params httprouter.Params) {

  userId := req.Header.Get("id_user")
  parsedUuid := uuid.MustParse(userId)

  // Convert data stream to json
	requestBody := dto.ProductCreateDTO{}
	helper.ReadRequestBody(req, &requestBody)

  requestBody.Id_user = parsedUuid

	// Call servicee
	response := controller.productService.Create(req.Context(), requestBody)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *ProductControllerImpl) FindById(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// Get Id From Params
	productId := params.ByName("productId")

	// Convert to uuid from string
	newId, err := uuid.Parse(productId)

	if err != nil {
		helper.PanicIfError(err)
	}

	// Call service
	response := controller.productService.FindById(req.Context(), newId)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (controller *ProductControllerImpl) Delete(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// Get Id From Params
	productId := params.ByName("productId")
	// Convert to uuid from string
	newId := uuid.MustParse(productId)

	// Call service
	controller.productService.Delete(req.Context(), newId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *ProductControllerImpl) Update(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// Get Id From Params
	productId := params.ByName("productId")
	// Convert to uuid from string
	newId := uuid.MustParse(productId)

	requestBody := response.UpdateProductRequest{}
	helper.ReadRequestBody(req, &requestBody)

	requestBody.Id = newId

	// Call service
	response := controller.productService.Update(req.Context(), requestBody)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *ProductControllerImpl) FindAll(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	responses := controller.productService.FindAll(req.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (c *ProductControllerImpl) FindAllByIdSeller(w http.ResponseWriter, req *http.Request, params httprouter.Params){  

  id_seller := req.Header.Get("id_user")

  result := c.productService.FindAllByIdSeller(req.Context(), uuid.MustParse(id_seller))
 
  var webResponse web.WebResponse
  
  if len(result) < 1 {
    webResponse = web.WebResponse {
      Code : http.StatusOK,
      Status: "Success",
      Data: []domain.Product{},
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




