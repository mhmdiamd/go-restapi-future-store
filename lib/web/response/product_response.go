package response

import (
	"time"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type ProductCreateResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Stock       int       `json:"stock"`
	Price       int       `json:"price"`
	Color       string    `json:"color"`
	Photo        []domain.ProductImage   `json:"photo"`
	Size        string    `json:"size"`
	Condition   string    `json:"condition"`
	Id_category uuid.UUID `json:"id_category"`
	Id_user     uuid.UUID `json:"id_user"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

type ProductResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Stock       int       `json:"stock"`
	Price       int       `json:"price"`
	Color       string    `json:"color"`
	Photo       []domain.ProductImage    `json:"photo"`
	Size        string    `json:"size"`
	Condition   string    `json:"condition"`
  Images []domain.ProductImage `json:"image"`
	Id_category uuid.UUID `json:"id_category"`
	Id_user     uuid.UUID `json:"id_user"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

type UpdateProductRequest struct {
	Id           uuid.UUID
	Product_name string
	Stock        int
	Price        int
	Description  string
	Photo        string
	Id_category  uuid.UUID
}
