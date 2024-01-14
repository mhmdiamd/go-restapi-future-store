package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type CartRepository interface {
  Create(ctx context.Context, tx *sqlx.Tx, body CartCreateBody) (domain.Cart, error)
  Delete(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) error
  Update(ctx context.Context, tx *sqlx.Tx, body CartUpdateBody) (domain.Cart, error)
  FindById(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) (domain.Cart, error)
  FindAll(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) []domain.Cart
} 
