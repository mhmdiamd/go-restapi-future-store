package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type ShippingAddressRepository interface {
  Create(ctx context.Context, tx *sqlx.Tx, body CreateShippingAddressBody) (domain.ShippingAddress, error)
  Update(ctx context.Context, tx *sqlx.Tx, body UpdateShippingAddressBody) (domain.ShippingAddress, error)
  Delete(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) error
  FindAllByIdUser(ctx context.Context, tx *sqlx.Tx, id_user uuid.UUID) []domain.ShippingAddress
  FindAll(ctx context.Context, tx *sqlx.Tx) []domain.ShippingAddress
  FindById(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) (domain.ShippingAddress, error)
  
  Activate(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) (domain.ShippingAddress, error)
}
