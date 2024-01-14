package repository

import "github.com/google/uuid"

type CartCreateBody struct {
  Id_user uuid.UUID `db:"id_user"`
  Id_product uuid.UUID `db:"id_product"`
  Qty int64 `db:"qty"`
}

type CartUpdateBody struct {
  Id uuid.UUID `db:"id"`
  Qty int64 `db:"qty"`
}


