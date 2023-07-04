package response

import "github.com/google/uuid"

type UpdateProductRequest struct {
	Id           uuid.UUID
	Product_name string
	Stock        int
	Price        int
	Description  string
	Photo        string
	Id_category  uuid.UUID
}
