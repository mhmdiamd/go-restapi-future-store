package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-restapi-future-store/exception"
	"github.com/mhmdiamd/go-restapi-future-store/helper"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type CartRepositoryImpl struct {

}

func NewCartRepositoryImpl() CartRepository {
  return &CartRepositoryImpl{}
}

func (r *CartRepositoryImpl) FindById(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) (domain.Cart, error) {

  query := "SELECT id, id_user, id_product, qty FROM carts WHERE id=$1 LIMIT 1"
  row, err := tx.QueryxContext(ctx, query, id)

  helper.PanicIfError(err)

  var cart = domain.Cart{}

  if !row.Next() {
    panic(exception.NewNotFoundError("Cart not found"))
  }else {
    err := row.Scan(
      &cart.Id,
      &cart.Id_user,
      &cart.Id_product,
      &cart.Qty,
    )
  
    if err != nil {
      fmt.Println(err)
      return cart, err
    }
  }

  return cart, nil
}

func (r CartRepositoryImpl) FindAll(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) []domain.Cart {
  query := `
  SELECT 
    carts.id as id, 
    carts.id_user as user, 
    carts.id_product as id_product, 
    carts.qty as qty, 
    JSONB_BUILD_OBJECT(        
      'id', products.id,
      'name', products.name,
      'stock', products.stock,
      'price', products.price,
      'description', products.description,
      'color', products.color,
      'condition', products.condition,
      'size', products.size,
      'id_user', products.id_user,
      'id_category', products.id_category
    ) AS product 
  FROM carts INNER JOIN products ON carts.id_product = products.id 
  WHERE carts.id_user=$1
  `
  rows, err := tx.QueryxContext(ctx, query, id)

  helper.PanicIfError(err)

  var carts []domain.Cart
  
  for rows.Next() {
    cart := domain.Cart{}

    err := rows.Scan(
      &cart.Id,
      &cart.Id_user,
      &cart.Id_product,
      &cart.Qty,
      &cart.Product,
    )

    helper.PanicIfError(err)

    carts = append(carts, cart)
  }
 
  return carts
}


func (r *CartRepositoryImpl) Create(ctx context.Context, tx *sqlx.Tx, body CartCreateBody) (domain.Cart, error) {

  fmt.Println(body)

  id := helper.GenerateUUID()

  query := "INSERT INTO carts(id, id_user, id_product, qty) VALUES($1, $2, $3, $4) RETURNING id, id_user, id_product, qty"
  row := tx.QueryRowContext(ctx, query, id, body.Id_user, body.Id_product, body.Qty)

  var cart domain.Cart

  err := row.Scan(
    &cart.Id,
    &cart.Id_user,
    &cart.Id_product,
    &cart.Qty,
  )

  if err != nil {
    return cart, err
  }

  return cart, nil
}

func (r *CartRepositoryImpl) Delete(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) error {
  query := "DELETE FROM carts WHERE id=$1"
  _, err := tx.ExecContext(ctx, query, id)

  if err != nil {
    return err
  }

  return  nil
}

func (r *CartRepositoryImpl) Update(ctx context.Context, tx *sqlx.Tx, body CartUpdateBody) (domain.Cart, error) {
  query := "UPDATE carts SET qty=$1 WHERE id=$2 RETURNING id, id_user, id_product, qty"
  row := tx.QueryRowContext(ctx, query, body.Qty, body.Id)

  var newCart domain.Cart

  err := row.Scan(
    &newCart.Id,
    &newCart.Id_user,
    &newCart.Id_product,
    &newCart.Qty,
  )

  if err != nil {
    return newCart, err
  }

  return newCart, nil
}


