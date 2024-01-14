package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-restapi-future-store/helper"
	"github.com/mhmdiamd/go-restapi-future-store/lib/repository"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/dto"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type ShippingAddressServiceImpl struct {
  shippingAddressRepository repository.ShippingAddressRepository
  DB *sqlx.DB
  Validator *validator.Validate
}

func NewShippingAddressRepositoryImpl(r repository.ShippingAddressRepository, db *sqlx.DB, validator *validator.Validate) ShippingAddressService {
  return &ShippingAddressServiceImpl{
    shippingAddressRepository: r,
    DB: db,
    Validator: validator,
  }
}

func (s *ShippingAddressServiceImpl) FindAll(ctx context.Context) []domain.ShippingAddress {
  tx := s.DB.MustBegin()
  defer helper.CommitOrRollback(tx)
  
  result := s.shippingAddressRepository.FindAll(ctx, tx)
  return result
}


func (s *ShippingAddressServiceImpl) FindAllByIdUser(ctx context.Context, id_user uuid.UUID) []domain.ShippingAddress {
  tx := s.DB.MustBegin()
  defer helper.CommitOrRollback(tx)

  result := s.shippingAddressRepository.FindAllByIdUser(ctx, tx, id_user)
  return result
}

func (s *ShippingAddressServiceImpl) FindById(ctx context.Context, id uuid.UUID) domain.ShippingAddress {
  tx := s.DB.MustBegin()
  defer helper.CommitOrRollback(tx)
  
  result, err := s.shippingAddressRepository.FindById(ctx, tx, id)
  helper.PanicIfError(err)

  return result
}

func (s *ShippingAddressServiceImpl) Create(ctx context.Context, body dto.CreateShippingAddressDTO) domain.ShippingAddress {
  tx := s.DB.MustBegin()
  defer helper.CommitOrRollback(tx)

  convertedBody := repository.CreateShippingAddressBody(body)
  
  result, err := s.shippingAddressRepository.Create(ctx, tx, convertedBody)
  helper.PanicIfError(err)

  return result
}

func (s *ShippingAddressServiceImpl) Delete(ctx context.Context, id uuid.UUID) string {
  tx := s.DB.MustBegin()
  defer helper.CommitOrRollback(tx)

  // Find first
  _, err := s.shippingAddressRepository.FindById(ctx, tx, id)
  helper.PanicIfError(err)

  tx2 := s.DB.MustBegin()
  defer helper.CommitOrRollback(tx2)

  err = s.shippingAddressRepository.Delete(ctx, tx2, id)
  helper.PanicIfError(err)

  return "Success delete shipping address"
}

func (s *ShippingAddressServiceImpl) Update(ctx context.Context, body dto.UpdateShippingAddressDTO ) domain.ShippingAddress {
  tx := s.DB.MustBegin()
  defer helper.CommitOrRollback(tx)

  // Find first
  _, err := s.shippingAddressRepository.FindById(ctx, tx, body.Id)
  helper.PanicIfError(err)

  tx2 := s.DB.MustBegin()
  defer helper.CommitOrRollback(tx2)

  convertedBody := repository.UpdateShippingAddressBody(body)

  result, err := s.shippingAddressRepository.Update(ctx, tx2, convertedBody)
  helper.PanicIfError(err)

  return result
}

func (s *ShippingAddressServiceImpl) Activate(ctx context.Context, id uuid.UUID) domain.ShippingAddress {
  tx := s.DB.MustBegin()
  defer helper.CommitOrRollback(tx)

  // Find first
  _, err := s.shippingAddressRepository.FindById(ctx, tx, id)
  helper.PanicIfError(err)

  tx2 := s.DB.MustBegin()
  defer helper.CommitOrRollback(tx2)

  result, err := s.shippingAddressRepository.Activate(ctx, tx2, id)
  helper.PanicIfError(err)

  return result
}
