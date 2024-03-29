package domain

import (
	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/lib"
)

type User struct {
	Id          uuid.UUID      `db:"id"`
	Name        string         `db:"name"`
	Email       string         `db:"email" validate:"required"`
	Password    string         `db:"password"`
	Birth_date  lib.NullTime   `db:"birth_date"`
	Store_name  lib.NullString `db:"store_name"`
	Description lib.NullString `db:"description"`
	Role        string         `db:"role"`
	Phone       lib.NullInt64  `db:"phone"`
	Address     lib.NullString `db:"address"`
	Photo       lib.NullString `db:"photo"`
}

type UserPhoto struct {
  Url  string `json:"url"`
}

type Seller struct {
	Id          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	Email       string         `json:"email" validate:"required"`
	Birth_date  lib.NullTime   `json:"birth_date"`
	Store_name  lib.NullString `json:"store_name"`
	Description lib.NullString `json:"description"`
	Role        string         `json:"role"`
	Phone       lib.NullInt64  `json:"phone"`
	Address     lib.NullString `json:"address"`
	Photo       lib.NullString `json:"photo"`
}

type Customer struct {
	Id          uuid.UUID      `db:"id"`
	Name        string         `db:"name"`
	Email       string         `db:"email" validate:"required"`
	Birth_date  lib.NullTime   `db:"birth_date"`
	Role        string         `db:"role"`
	Phone       lib.NullInt64  `db:"phone"`
	Address     lib.NullString `db:"address"`
	Photo       lib.NullString `db:"photo"`
}

type Authentication struct {
	Email    string `db:"email"`
	Password string `db:"password"`
}

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}
