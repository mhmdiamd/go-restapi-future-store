package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type SellerRepository interface {
  Update(c context.Context, tx *sqlx.Tx, body SellerUpdateBody) (domain.Seller, error) 
}
