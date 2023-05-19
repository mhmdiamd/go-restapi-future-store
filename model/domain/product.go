package domain

import "github.com/google/uuid"

type Product struct {
	Id          uuid.UUID
	ProductName string
	Stock       int
	Price       int
	Description string
	Photo       string
	Id_category uuid.UUID
}
