package response

import "github.com/google/uuid"

type CreateProductRequest struct {
	Product_name string
	Stock        int
	Price        int
	Description  string
	Photo        string
	Id_category  uuid.UUID
}
