package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-restapi-future-store/helper"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/dto"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type UserRepositoryImpl struct {

}

func NewUserRepositoryImpl() UserRepository {
  return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) Update(c context.Context, tx *sqlx.Tx, body dto.UserUpdateDTO) (domain.User, error) {

  var newUser domain.User
  query := "UPDATE users SET name=$1, birth_date=$2, store_name=$2, description=$3, role=$4, phone=$5, address=$6, photo=$7"

  err := tx.GetContext(c, &newUser, query, body.Name, body.Birth_date, body.Store_name, body.Role, body.Phone, body.Address, body.Photo, body.Description)

  if err != nil {
    helper.PanicIfError(err)
  }

  return newUser, nil

}


func (r *UserRepositoryImpl) GetProfile(c context.Context, tx *sqlx.Tx, user_id uuid.UUID) (domain.User, error) {
  
  var user domain.User
  query := "SELECT * FROM users where id=$1"

  err := tx.GetContext(c, &user, query, user_id)

  if err != nil {
    return user, errors.New("User not found")
  }

  return user, nil
}

