package repository

import "github.com/google/uuid"

type CreateShippingAddressBody struct {
  Id_user uuid.UUID `db:"id_user"`
  Place string `db:"place"`
  Recipient_name string `db:"recipient_name"`
  Recipient_phone string `db:"recipient_phone"`
  Address string `db:"address"`
  Postal_code string `db:"postal_code"`
  City_or_subdistrict string `db:"city_or_subdistrict"`
  Is_active bool `db:"is_active"`
}

type UpdateShippingAddressBody struct {
  Id uuid.UUID `db:"id"`
  Place string `db:"place"`
  Recipient_name string `db:"recipient_name"`
  Recipient_phone string `db:"recipient_phone"`
  Address string `db:"address"`
  Postal_code string `db:"postal_code"`
  City_or_subdistrict string `db:"city_or_subdistrict"`
  Is_active bool `db:"is_active"`
}
