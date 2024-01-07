package domain

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
  ID           uuid.UUID `db:"id"`
	Name string `db:"name"`
	Description  string `db:"description"`
	Stock        int `db:"stock"`
	Price        int `db:"price"`
	Color        string `db:"color"`
  Size string `db:"size"`
  Condition string `db:"condition"`
	Photo        JSONBArray `db:"photo"`
	Id_category  uuid.UUID `db:"id_category"`
	Id_user  uuid.UUID `db:"id_user"`
  Created_at time.Time `db:"created_at"`
  Updated_at time.Time `db:"updated_at"`
}
