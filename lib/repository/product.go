package repository

import "github.com/google/uuid"

type Product struct {
	Id           uuid.UUID
	Product_name string
	Stock        int
	Price        int
	Description  string
	Photo        string
	Id_category  uuid.UUID
}
