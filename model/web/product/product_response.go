package web

import "github.com/google/uuid"

type ProductResponse struct {
	Id           uuid.UUID `json:"id"`
	Product_name string    `json:"product_name"`
	Stock        int       `json:"stock"`
	Price        int       `json:"price"`
	Description  string    `json:"description"`
	Photo        string    `json:"photo"`
	Id_category  uuid.UUID `json:"id_category"`
}
