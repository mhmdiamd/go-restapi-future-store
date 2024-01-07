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

type ImageControllerImpl struct {
  ProductImageService service.ImageService
}

func NewImageControllerImpl(imageService service.ImageService) ImageController {
  return &ImageControllerImpl{
    ProductImageService: imageService,
  }
}

func (c *ImageControllerImpl) UploadProductImage(w http.ResponseWriter, req *http.Request, params httprouter.Params) {

  userId := req.Header.Get("id_user")
  parsedUuid := uuid.MustParse(userId)
  
  file, fileHeader, err := req.FormFile("photo")
  id_product := req.FormValue("id_product")

  helper.PanicIfError(err)

  fileHandler := dto.FileHandler{
    FileHeader: fileHeader,
    File: file,
  }

  body := dto.CreateProductImageDto{
    File: fileHandler,
    Product_id: id_product, 
    User_id: parsedUuid,
  }

  result := c.ProductImageService.UploadProductImage(req.Context(), body)

  webResponse := web.WebResponse {
    Status: "Success",
    Code: http.StatusOK,
    Data: result,
  }

  helper.WriteToResponseBody(w, webResponse)
}
