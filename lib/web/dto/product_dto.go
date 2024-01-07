package dto

import (

	"github.com/google/uuid"
)

type ProductCreateDTO struct {
  Name string `db:"name"`
	Stock        int `db:"stock"`
	Price        int `db:"price"`
	Description  string `db:"description"`
	Size  string `db:"size"`
	Condition  string `db:"condition"`
	Color  string `db:"color"`
	Id_category  uuid.UUID `db:"id_category"`
  Id_user uuid.UUID `db:"id_user"`
}

