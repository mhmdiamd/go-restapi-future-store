package product

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/mhmdiamd/go-restapi-future-store/lib/helpers"
	"github.com/mhmdiamd/go-restapi-future-store/lib/service"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/response"
)

type ProductControllerImpl struct {
	productService service.ProductService
}

func NewProductController(service service.ProductService) ProductController {
	return &ProductControllerImpl{
		productService: service,
	}
}

func (controller *ProductControllerImpl) Create(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// Convert data stream to json
	requestBody := response.CreateProductRequest{}
	helpers.ReadRequestBody(req, &requestBody)

	// Call servicee
	response, err := controller.productService.Create(req.Context(), requestBody)
	helpers.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helpers.WriteToResponse(w, "application/json", webResponse)
}

func (controller *ProductControllerImpl) FindById(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// Get Id From Params
	productId := params.ByName("productId")
	// Convert to uuid from string
	newId := uuid.MustParse(productId)

	// Call service
	response := controller.productService.FindById(req.Context(), newId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helpers.WriteToResponse(w, "application/json", webResponse)
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

	helpers.WriteToResponse(w, "application/json", webResponse)
}

func (controller *ProductControllerImpl) Update(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// Get Id From Params
	productId := params.ByName("productId")
	// Convert to uuid from string
	newId := uuid.MustParse(productId)

	requestBody := response.UpdateProductRequest{}
	helpers.ReadRequestBody(req, &requestBody)

	requestBody.Id = newId

	// Call service
	response, err := controller.productService.Update(req.Context(), requestBody)
	helpers.PanicIfError(err)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helpers.WriteToResponse(w, "application/json", webResponse)
}

func (controller *ProductControllerImpl) FindAll(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	responses := controller.productService.FindAll(req.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	}
	helpers.WriteToResponse(w, "application/json", webResponse)

}
