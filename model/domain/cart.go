package domain

import (
	"github.com/google/uuid"
)

type Cart struct {
  Id uuid.UUID `db:"id" json:"id"`
  Id_user uuid.UUID `db:"id_user" json:"id_user"`
  Id_product uuid.UUID `db:"id_product" json:"id_product"`
  Qty int64 `db:"qty" json:"qty"`
  Product JSONBProduct`db:"product" json:"product"`
}


