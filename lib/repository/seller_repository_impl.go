package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-restapi-future-store/helper"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type SellerRepositoryImpl struct {

}

func NewSellerRepositoryImpl() SellerRepository {
  return &SellerRepositoryImpl{}
}

func (r *SellerRepositoryImpl) Update(ctx context.Context, tx *sqlx.Tx, body SellerUpdateBody) (domain.Seller, error) {
 
  var newSeller domain.Seller

  fmt.Println(body)

  query := `UPDATE users SET name=$1, description=$2, phone=$3, address=$4, birth_date=$5, store_name=$6 WHERE id=$7 RETURNING id, name, email, birth_date, store_name, description, role, phone, address, photo`
  row := tx.QueryRowContext(ctx,query, body.Name, body.Description, body.Phone, body.Address, body.Birth_date, body.Store_name, body.ID)

  err := row.Scan(
    &newSeller.Id,
    &newSeller.Name,
    &newSeller.Email,
    &newSeller.Birth_date,
    &newSeller.Store_name,
    &newSeller.Description,
    &newSeller.Role,
    &newSeller.Phone,
    &newSeller.Address,
    &newSeller.Photo,
  )

  helper.PanicIfError(err)

  return newSeller, nil

}
