package controller

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/mhmdiamd/go-restapi-future-store/helper"
	"github.com/mhmdiamd/go-restapi-future-store/lib/service"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/dto"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
	"github.com/mhmdiamd/go-restapi-future-store/model/web"
)

type SellerControllerImpl struct {
  SellerService service.SellerService
}

func NewSellerControllerImpl(sellerService service.SellerService) SellerController {
  return &SellerControllerImpl{
    SellerService: sellerService,
  }
}

func (c *SellerControllerImpl) Update(w http.ResponseWriter, req *http.Request, p httprouter.Params) {

  userId := p.ByName("id")
  parsedId := uuid.MustParse(userId)

  var body dto.SellerUpdateDTO

  body.ID = parsedId
  helper.ReadRequestBody(req, &body)

  fmt.Println(body)

  result := c.SellerService.Update(req.Context(), body)

  webResponse := web.WebResponse{
    Status: "Success",
    Code: http.StatusOK,
    Data: result,
  }

  helper.WriteToResponseBody(w, webResponse)
}

func (c *SellerControllerImpl) UploadPhoto(w http.ResponseWriter, req *http.Request, params httprouter.Params) {

  file, fileHeader, err := req.FormFile("photo")
  userId := req.Header.Get("id_user")

  parsedId := uuid.MustParse(userId)

  helper.PanicIfError(err)

  fileHandler := dto.FileHandler{
    File: file,
    FileHeader: fileHeader,
  }

  body := dto.UploadSellerPhotoDTO{
    File: fileHandler,
    User_id : parsedId,
  }

  result := c.SellerService.UploadPhoto(req.Context(), body)

  newData := domain.UserPhoto {
    Url: result,
  }

  webResponse := web.WebResponse {
    Status: "Success",
    Code: http.StatusOK,
    Data: newData,
  }

  helper.WriteToResponseBody(w, webResponse)

}
