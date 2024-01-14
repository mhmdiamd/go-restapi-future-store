package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-restapi-future-store/helper"
	"github.com/mhmdiamd/go-restapi-future-store/lib/repository"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type CartServiceImpl struct {
  CartRepository repository.CartRepository 
  Db *sqlx.DB
  Validator *validator.Validate
}

func NewCartServiceImpl(cartRepository repository.CartRepository, DB *sqlx.DB, validator *validator.Validate) CartService {
  return &CartServiceImpl{
    CartRepository : cartRepository,
    Db: DB,
    Validator: validator,
  }
}

func (r *CartServiceImpl) Create(ctx context.Context, body repository.CartCreateBody) domain.Cart {

  tx := r.Db.MustBegin()
  defer helper.CommitOrRollback(tx)

  tx2 := r.Db.MustBegin()
  defer helper.CommitOrRollback(tx2)

  result, err := r.CartRepository.Create(ctx, tx, body)

  helper.PanicIfError(err)

  return result
}

func (r *CartServiceImpl) FindAll(ctx context.Context, id_user uuid.UUID) []domain.Cart {
  tx := r.Db.MustBegin()
  defer helper.CommitOrRollback(tx)

  result := r.CartRepository.FindAll(ctx, tx, id_user) 

  return result
}

func (r *CartServiceImpl) Delete(ctx context.Context, id uuid.UUID) string {
  tx := r.Db.MustBegin()
  defer helper.CommitOrRollback(tx)

  err := r.CartRepository.Delete(ctx, tx, id)
  helper.PanicIfError(err)

  return "Success delete cart"
}

func (r *CartServiceImpl) Increment(ctx context.Context, id uuid.UUID) domain.Cart {
  tx := r.Db.MustBegin()
  defer helper.CommitOrRollback(tx)

  tx2 := r.Db.MustBegin()
  defer helper.CommitOrRollback(tx2)

  // Find first
  cart, err := r.CartRepository.FindById(ctx, tx2, id)
  helper.PanicIfError(err)

  newQty := cart.Qty + 1
  newData := repository.CartUpdateBody{
    Id: id,
    Qty: newQty,
  }

  result, err := r.CartRepository.Update(ctx, tx, newData)
  helper.PanicIfError(err)

  return result
}

func (r *CartServiceImpl) Decrement(ctx context.Context, id uuid.UUID) domain.Cart {
  tx := r.Db.MustBegin()
  defer helper.CommitOrRollback(tx)

  tx2 := r.Db.MustBegin()
  defer helper.CommitOrRollback(tx2)

  tx3 := r.Db.MustBegin()
  defer helper.CommitOrRollback(tx3)

  // Find first
  cart, err := r.CartRepository.FindById(ctx, tx, id)
  helper.PanicIfError(err)

  if cart.Qty == 1 {
    var _ = r.CartRepository.Delete(ctx, tx2, id)
    return domain.Cart{}
  }

  newQty := cart.Qty - 1
  newData := repository.CartUpdateBody{
    Id: id,
    Qty: newQty,
  }

  result, err := r.CartRepository.Update(ctx, tx3, newData)
  helper.PanicIfError(err)

  return result
}


