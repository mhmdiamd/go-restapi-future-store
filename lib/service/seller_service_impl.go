package service

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-restapi-future-store/exception"
	"github.com/mhmdiamd/go-restapi-future-store/helper"
	"github.com/mhmdiamd/go-restapi-future-store/lib/repository"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/dto"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type SellerServiceImpl struct {
  SellerRepositoy repository.SellerRepository
  DB *sqlx.DB
  Validator *validator.Validate
}

func NewSellerServiceImpl(sellerRepository repository.SellerRepository, db *sqlx.DB, validator *validator.Validate) SellerService {
  return &SellerServiceImpl{
    SellerRepositoy : sellerRepository,
    DB: db,
    Validator: validator,
  } 
} 

func (s *SellerServiceImpl) Update(ctx context.Context, body dto.SellerUpdateDTO) domain.Seller {
  tx := s.DB.MustBegin()

  newBody := repository.SellerUpdateBody {
    ID : body.ID,
    Name : body.Name,
    Store_name : body.Store_name,
    Birth_date : body.Birth_date,
    Address : body.Address,
    Phone : body.Phone,
    Description : body.Description,
  }

  result, err := s.SellerRepositoy.Update(ctx, tx, newBody)

  helper.PanicIfError(err)

  return result
}

func (s *SellerServiceImpl) UploadPhoto(ctx context.Context, body dto.UploadSellerPhotoDTO) string {
  fmt.Println(body)
  fileIdentity, err := helper.DecodedImage(body.File, body.User_id)

  if err != nil {
    panic(exception.NewForbiddenError(err.Error()))
  }

  newUrl := "https://" + fileIdentity.Filename + ".com"

  return newUrl

}


