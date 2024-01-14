package repository

import (
	"context"
	"errors"
	"fmt"

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
  fmt.Println(body)

  query := "UPDATE users SET name=$1, birth_date=$2, store_name=$3, description=$4, role=$5, phone=$6, address=$7 WHERE id=$8 RETURNING id, name, email, birth_date, store_name, description, role, phone, address, photo"
  row := tx.QueryRowContext(c,query, body.Name, body.Birth_date, body.Store_name, body.Description, body.Role, body.Phone, body.Address, body.ID)

  var newUser domain.User

  err := row.Scan(
    &newUser.Id,
    &newUser.Name,
    &newUser.Email,
    &newUser.Birth_date,
    &newUser.Store_name,
    &newUser.Description,
    &newUser.Role,
    &newUser.Phone,
    &newUser.Address,
    &newUser.Photo,
  )

  helper.PanicIfError(err)

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

