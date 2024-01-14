package repository

import (
	"time"

	"github.com/google/uuid"
)


type SellerUpdateBody struct {
	ID          uuid.UUID       `db:"id"`
  Name        string         `db:"name"`
	Birth_date  time.Time    `db:"birth_date"`
	Store_name  string    `db:"store_name"`
	Description string    `db:"description"`
	Phone       int64     `db:"phone"`
	Address     string    `db:"address"`
}

type SellerPhoto struct {
  Url string  `json:"url"`
}
