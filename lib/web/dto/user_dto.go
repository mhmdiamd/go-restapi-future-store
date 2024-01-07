package dto

import (
	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/lib"
)

type UserUpdateDTO struct {
	ID          uuid.UUID         `db:"name"`
  Name        string         `db:"name"`
	Birth_date  lib.NullTime   `db:"birth_date"`
	Store_name  lib.NullString `db:"store_name"`
	Description lib.NullString `db:"description"`
	Role        string         `db:"role"`
	Phone       lib.NullInt64  `db:"phone"`
	Address     lib.NullString `db:"address"`
	Photo       lib.NullString `db:"photo"`
}


