package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/mhmdiamd/go-restapi-future-store/helper"
	"github.com/mhmdiamd/go-restapi-future-store/lib/service"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/response"
	"github.com/mhmdiamd/go-restapi-future-store/model/web"
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
	categoryCreateRequest := response.CategoryCreateRequest{}

	// Encode data json dan memasukanya ke dalam variable categoryCreateRequest
  err := decoder.Decode(&categoryCreateRequest)
	helper.PanicIfError(err)

	// Kirim data ke service
	response, err := controller.CategoryService.Create(req.Context(), categoryCreateRequest)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	// Write to Response
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) Update(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// Get Id
	categoryId := params.ByName("id")
	newId := uuid.MustParse(categoryId)

	// Decode request body to json
	categoryUpdateRequest := response.CategoryUpdateRequest{}
	helper.ReadRequestBody(req, &categoryUpdateRequest)

	categoryUpdateRequest.Id = newId

	// Send Data to service
	response, err := controller.CategoryService.Update(req.Context(), categoryUpdateRequest)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) Delete(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// Get Id
	categoryId := params.ByName("id")
	newId := uuid.MustParse(categoryId)

	// Send Data to service
	controller.CategoryService.Delete(req.Context(), newId)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) FindById(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// Get Id
	categoryId := params.ByName("id")
	fmt.Println(categoryId)
	newId := uuid.MustParse(categoryId)

	// Send Data to service
	response := controller.CategoryService.FindById(req.Context(), newId)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// Send Data to service
	response := controller.CategoryService.FindAll(req.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}
	helper.WriteToResponseBody(w, webResponse)
}
