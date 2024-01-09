package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-restapi-future-store/exception"
	"github.com/mhmdiamd/go-restapi-future-store/helper"
	"github.com/mhmdiamd/go-restapi-future-store/lib/repository"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/dto"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type ImageServiceImpl struct {
  imageRepository repository.ImageRepository
  DB *sqlx.DB
  Validator *validator.Validate
}

func NewImageServiceImpl(repository repository.ImageRepository, db *sqlx.DB, validate *validator.Validate) ImageService {
  return &ImageServiceImpl{
    imageRepository : repository,
    DB : db,
    Validator : validate,
  }
}


func (s *ImageServiceImpl) FindProductImageById(ctx context.Context, Id_product_image uuid.UUID) domain.ProductImage {

  tx := s.DB.MustBegin()
  defer helper.CommitOrRollback(tx)

  result, err := s.imageRepository.FindProductImageById(ctx, tx, Id_product_image)

  if err != nil {
    panic(exception.NewForbiddenError(err.Error()))
  }

 return result
}

func (s *ImageServiceImpl) UploadProductImage(ctx context.Context, data dto.CreateProductImageDto) domain.ProductImage {
  url := "https://example.com"

  fileIdentity, err := helper.DecodedImage(data.File, data.User_id)

  if err != nil {
    panic(exception.NewForbiddenError(err.Error()))
  }

  body := dto.CreateProductBody{
    Name: fileIdentity.Filename,
    Product_id : data.Product_id,
    Url: url,
  }

  tx := s.DB.MustBegin()
  defer tx.Commit()

  result, err := s.imageRepository.UploadProductImage(ctx, tx, body)

  if err != nil {
    panic(exception.NewForbiddenError(err.Error()))
  }

 return result
}

func (s *ImageServiceImpl) UpdateProductImage(ctx context.Context, data dto.UpdateProductImageDto) domain.ProductImage {
  fileIdentity, err := helper.DecodedImage(data.File, data.User_id)

  if err != nil {
    panic(exception.NewForbiddenError(err.Error()))
  }

  body := dto.UpdateProductBody{
    Name: fileIdentity.Filename,
    Id_product_image : data.Id_product_image,
  }

  tx := s.DB.MustBegin()
  defer helper.CommitOrRollback(tx)

  result, err := s.imageRepository.UpdateProductImage(ctx, tx, body)

  if err != nil {
    panic(exception.NewForbiddenError(err.Error()))
  }

 return result
}

func (s *ImageServiceImpl) DeleteProductImage(ctx context.Context, Id_product_image uuid.UUID) string {

  tx := s.DB.MustBegin()
  defer helper.CommitOrRollback(tx)

  // Check is the product exist
  s.FindProductImageById(ctx, Id_product_image)

  err := s.imageRepository.DeleteProductImage(ctx, tx, Id_product_image)

  if err != nil {
    helper.PanicIfError(err)
  }

  return "Success delete product image"
}


func (s *ImageServiceImpl) GetAllByIdProduct(ctx context.Context, id_product uuid.UUID) []domain.ProductImage {

  tx := s.DB.MustBegin()

  images := s.imageRepository.GetAllByIdProduct(ctx, tx, id_product)
  return images
}
