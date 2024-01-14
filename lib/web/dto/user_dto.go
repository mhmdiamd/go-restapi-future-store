package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserUpdateDTO struct {
	ID          uuid.UUID  `db:"id"`
  Name        string     `db:"name"`
	Birth_date  time.Time  `db:"birth_date"`
	Store_name  string     `db:"store_name"`
	Description string     `db:"description"`
	Role        string     `db:"role"`
	Phone       int64      `db:"phone"`
	Address     string     `db:"address"`
}

type SellerUpdateDTO struct {
	ID          uuid.UUID  `db:"id"`
  Name        string         `db:"name"`
	Birth_date  time.Time  `db:"birth_date"`
	Store_name  string `db:"store_name"`
	Description string `db:"description"`
	Phone       int64  `db:"phone"`
	Address     string `db:"address"`
}

type CustomterUpdateDto struct {
	ID          uuid.UUID      `db:"name"`
  Name        string         `db:"name"`
	Birth_date  time.Time  `db:"birth_date"`
	Description string `db:"description"`
	Phone       int64  `db:"phone"`
	Address     string `db:"address"`
}


