package domain

import (
	"time"

	"github.com/google/uuid"
)

type ShippingAddress struct {
  Id uuid.UUID `db:"id" json:"id"`
  Id_user uuid.UUID `db:"id_user" json:"id_user"`
  Place string `db:"place" json:"place"`
  Recipient_name string `db:"recipient_name" json:"recipient_name"`
  Recipient_phone string `db:"recipient_phone" json:"recipient_phone"`
  Address string `db:"address" json:"address"`
  Postal_code string `db:"postal_code" json:"postal_code"`
  City_or_subdistrict string `db:"city_or_subdistrict" json:"city_or_subdistrict"`
  Is_active bool `db:"is_active" json:"is_active"`
  Created_at time.Time `db:"created_at" json:"created_at"`
  Updated_at time.Time `db:"updated_at" json:"updated_at"`
}
