package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-restapi-future-store/exception"
	"github.com/mhmdiamd/go-restapi-future-store/helper"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type ShippingAddressRepositoryImpl struct {}

func NewShippingAddressRepositoryImpl() ShippingAddressRepository {
  return &ShippingAddressRepositoryImpl{}
}

func (r *ShippingAddressRepositoryImpl) FindAll(ctx context.Context, tx *sqlx.Tx) []domain.ShippingAddress {

  query := "SELECT * FROM shipping_addresses ORDER BY created_at DESC LIMIT 10"
  row, err := tx.QueryxContext(ctx, query)

  helper.PanicIfError(err)

  var shippingAddresses []domain.ShippingAddress

  for row.Next() {
    var shippingAddress domain.ShippingAddress

    err := row.Scan(
      &shippingAddress.Id,
      &shippingAddress.Id_user,
      &shippingAddress.Place,
      &shippingAddress.Recipient_name,
      &shippingAddress.Recipient_phone,
      &shippingAddress.Address,
      &shippingAddress.Postal_code,
      &shippingAddress.City_or_subdistrict,
      &shippingAddress.Is_active,
      &shippingAddress.Created_at,
      &shippingAddress.Updated_at,
    )

    helper.PanicIfError(err)

    shippingAddresses = append(shippingAddresses, shippingAddress)

  }

  return shippingAddresses
}


func (r *ShippingAddressRepositoryImpl) FindAllByIdUser(ctx context.Context, tx *sqlx.Tx, id_user uuid.UUID) []domain.ShippingAddress {

  query := "SELECT * FROM shipping_addresses WHERE id_user=$1 LIMIT 1"
  row, err := tx.QueryxContext(ctx, query, id_user)

  helper.PanicIfError(err)

  var shippingAddresses []domain.ShippingAddress

  for row.Next() {
    var shippingAddress domain.ShippingAddress

    err := row.Scan(
      &shippingAddress.Id,
      &shippingAddress.Id_user,
      &shippingAddress.Place,
      &shippingAddress.Recipient_name,
      &shippingAddress.Recipient_phone,
      &shippingAddress.Address,
      &shippingAddress.Postal_code,
      &shippingAddress.City_or_subdistrict,
      &shippingAddress.Is_active,
      &shippingAddress.Created_at,
      &shippingAddress.Updated_at,
    )

    helper.PanicIfError(err)

    shippingAddresses = append(shippingAddresses, shippingAddress)

  }

  return shippingAddresses
}

func (r *ShippingAddressRepositoryImpl) FindById(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) (domain.ShippingAddress, error) {

  query := "SELECT * FROM shipping_addresses WHERE id=$1 LIMIT 1"
  row, err := tx.QueryxContext(ctx, query, id)

  helper.PanicIfError(err)

  var shippingAddress domain.ShippingAddress

  if !row.Next() {
    panic(exception.NewNotFoundError("Shipping address not found"))
  }else {
    err := row.Scan(
      &shippingAddress.Id,
      &shippingAddress.Id_user,
      &shippingAddress.Place,
      &shippingAddress.Recipient_name,
      &shippingAddress.Recipient_phone,
      &shippingAddress.Address,
      &shippingAddress.Postal_code,
      &shippingAddress.City_or_subdistrict,
      &shippingAddress.Is_active,
      &shippingAddress.Created_at,
      &shippingAddress.Updated_at,
    )
  
    if err != nil {
      return shippingAddress, err
    }
  }

  return shippingAddress, nil
}

func (r *ShippingAddressRepositoryImpl) Create(ctx context.Context, tx *sqlx.Tx, body CreateShippingAddressBody) (domain.ShippingAddress, error) {
  id := helper.GenerateUUID()

  query := `
    INSERT INTO 
      shipping_addresses(id, id_user, place, recipient_name, recipient_phone, address, postal_code, city_or_subdistrict, is_active)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
    RETURNING id, id_user, place, recipient_name, recipient_phone, address, postal_code, city_or_subdistrict, is_active, created_at, updated_at
  `  

  row := tx.QueryRowContext(ctx, query, 
    id, 
    body.Id_user,
    body.Place,
    body.Recipient_name,
    body.Recipient_phone,
    body.Address,
    body.Postal_code,
    body.City_or_subdistrict,
    body.Is_active,
  )

  var shippingAddress domain.ShippingAddress

  err := row.Scan(
    &shippingAddress.Id,
    &shippingAddress.Id_user,
    &shippingAddress.Place,
    &shippingAddress.Recipient_name,
    &shippingAddress.Recipient_phone,
    &shippingAddress.Address,
    &shippingAddress.Postal_code,
    &shippingAddress.City_or_subdistrict,
    &shippingAddress.Is_active,
    &shippingAddress.Created_at,
    &shippingAddress.Updated_at,
  )

  if err != nil {
    return shippingAddress, err
  }

  return shippingAddress, nil
}

func (r *ShippingAddressRepositoryImpl) Update(ctx context.Context, tx *sqlx.Tx, body UpdateShippingAddressBody) (domain.ShippingAddress, error) {
  query := `
    UPDATE 
      shipping_addresses 
    SET 
      id=$1, 
      place=$2, 
      recipient_name=$3, 
      recipient_phone=$4, 
      address=$5, 
      postal_code=$6, 
      city_or_subdistrict=$7, 
      is_active=$8
    RETURNING 
      id, 
      id_user, 
      place, 
      recipient_name, 
      recipient_phone, 
      address, 
      postal_code, 
      city_or_subdistrict, 
      is_active, 
      created_at, 
      updated_at
  `  

  row := tx.QueryRowContext(ctx, query, 
    body.Id, 
    body.Place,
    body.Recipient_name,
    body.Recipient_phone,
    body.Address,
    body.Postal_code,
    body.City_or_subdistrict,
    body.Is_active,
  )

  var shippingAddress domain.ShippingAddress

  err := row.Scan(
    &shippingAddress.Id,
    &shippingAddress.Id_user,
    &shippingAddress.Place,
    &shippingAddress.Recipient_name,
    &shippingAddress.Recipient_phone,
    &shippingAddress.Address,
    &shippingAddress.Postal_code,
    &shippingAddress.City_or_subdistrict,
    &shippingAddress.Is_active,
    &shippingAddress.Created_at,
    &shippingAddress.Updated_at,
  )

  if err != nil {
    return shippingAddress, err
  }

  return shippingAddress, nil
}

func (r *ShippingAddressRepositoryImpl) Delete(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) error {
  query := `DELETE FROM shipping_addresses WHERE id=$1`  

  _, err := tx.ExecContext(ctx, query, id)


  if err != nil {
    return err
  }

  return nil
}

func (r *ShippingAddressRepositoryImpl) Activate(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) (domain.ShippingAddress, error) {
  query := `
    UPDATE 
      shipping_addresses 
    SET 
      is_active='true'
    WHERE id=$1
    RETURNING 
      id, 
      id_user, 
      place, 
      recipient_name, 
      recipient_phone, 
      address, 
      postal_code, 
      city_or_subdistrict, 
      is_active, 
      created_at, 
      updated_at
  `  

  row := tx.QueryRowContext(ctx, query, id)

  var shippingAddress domain.ShippingAddress

  err := row.Scan(
    &shippingAddress.Id,
    &shippingAddress.Id_user,
    &shippingAddress.Place,
    &shippingAddress.Recipient_name,
    &shippingAddress.Recipient_phone,
    &shippingAddress.Address,
    &shippingAddress.Postal_code,
    &shippingAddress.City_or_subdistrict,
    &shippingAddress.Is_active,
    &shippingAddress.Created_at,
    &shippingAddress.Updated_at,
  )

  if err != nil {
    return shippingAddress, err
  }

  return shippingAddress, nil
}
