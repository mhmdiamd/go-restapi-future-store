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
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/response"
)

type UserServiceImpl struct {
  repository repository.UserRepository
  DB *sqlx.DB
  Validate *validator.Validate
}

func NewUserServiceImpl(repository repository.UserRepository, db *sqlx.DB, v *validator.Validate) UserService {
  return &UserServiceImpl{
    repository: repository,
    DB: db,
    Validate: v,
  }
}

func (s *UserServiceImpl) Update(c context.Context, body dto.UserUpdateDTO) (response.UserResponse, error) {

  tx := s.DB.MustBegin()
  defer helper.CommitOrRollback(tx)

  _, err := s.repository.GetProfile(c, tx, body.ID)

  if err != nil {
    if err.Error() == "User not found"{
      panic(exception.NewNotFoundError(err.Error()))
    }
    helper.PanicIfError(err)
  }

  result, err := s.repository.Update(c, tx, body)

  if err != nil {
    helper.PanicIfError(err)
  }

  newBody := response.UserResponse{
    Id: result.Id,
    Name: result.Name,
    Email: result.Email,
    Role: result.Role,
    Birth_date: result.Birth_date,
    Store_name: result.Store_name,
    Address: result.Address,
    Phone: result.Phone,
    Photo: result.Photo,
    Description: result.Description,
  }

  return newBody, nil
}

func (s *UserServiceImpl) GetProfile(c context.Context, user_id uuid.UUID) (response.UserResponse, error) {
tx := s.DB.MustBegin()
defer helper.CommitOrRollback(tx)

result, err := s.repository.GetProfile(c, tx, user_id)

if err != nil {
  if err.Error() == "User not found"{
    panic(exception.NewNotFoundError(err.Error()))
  }
  helper.PanicIfError(err)
}

newBody := response.UserResponse{
  Id: result.Id,
  Name: result.Name,
  Email: result.Email,
  Role: result.Role,
  Birth_date: result.Birth_date,
  Store_name: result.Store_name,
  Address: result.Address,
  Phone: result.Phone,
  Photo: result.Photo,
  Description: result.Description,
}

return newBody, nil

}
