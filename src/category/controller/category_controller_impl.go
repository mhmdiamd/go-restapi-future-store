package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/mhmdiamd/go-restapi-future-store/helpers"
	httpResponse "github.com/mhmdiamd/go-restapi-future-store/model/web"
	categoryWeb "github.com/mhmdiamd/go-restapi-future-store/model/web/category"
	"github.com/mhmdiamd/go-restapi-future-store/src/category/service"

	_ "github.com/lib/pq"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(service service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: service,
	}
}

func (controller *CategoryControllerImpl) Create(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// Buat Decoder untuk melakukan decode request body dari data stream ke bentuk json
	decoder := json.NewDecoder(req.Body)
	categoryCreateRequest := categoryWeb.CategoryCreateRequest{}

	// Encode data json dan memasukanya ke dalam variable categoryCreateRequest
	err := decoder.Decode(&categoryCreateRequest)
	helpers.PanicIfError(err)

	// Kirim data ke service
	response, err := controller.CategoryService.Create(req.Context(), categoryCreateRequest)
	helpers.PanicIfError(err)

	webResponse := httpResponse.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	// Write to Response
	helpers.WriteToResponse(w, "application/json", webResponse)
}

func (controller *CategoryControllerImpl) Update(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// Get Id
	categoryId := params.ByName("categoryId")
	newId := uuid.MustParse(categoryId)

	// Decode request body to json
	categoryUpdateRequest := categoryWeb.CategoryUpdateRequest{}
	helpers.ReadRequestBody(req, &categoryUpdateRequest)

	categoryUpdateRequest.Id = newId

	// Send Data to service
	response, err := controller.CategoryService.Update(req.Context(), categoryUpdateRequest)
	helpers.PanicIfError(err)

	webResponse := httpResponse.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helpers.WriteToResponse(w, "application/json", webResponse)
}

func (controller *CategoryControllerImpl) Delete(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// Get Id
	categoryId := params.ByName("categoryId")
	newId := uuid.MustParse(categoryId)

	// Send Data to service
	controller.CategoryService.Delete(req.Context(), newId)

	webResponse := httpResponse.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helpers.WriteToResponse(w, "application/json", webResponse)
}

func (controller *CategoryControllerImpl) FindById(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// Get Id
	categoryId := params.ByName("categoryId")
	fmt.Println(categoryId)
	newId := uuid.MustParse(categoryId)

	// Send Data to service
	response := controller.CategoryService.FindById(req.Context(), newId)

	webResponse := httpResponse.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helpers.WriteToResponse(w, "application/json", webResponse)
}

func (controller *CategoryControllerImpl) FindAll(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// Send Data to service
	response := controller.CategoryService.FindAll(req.Context())
	webResponse := httpResponse.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}
	helpers.WriteToResponse(w, "application/json", webResponse)
}
