package dto

import (
	"mime/multipart"

	"github.com/google/uuid"
)


type FileHandler struct {
  FileHeader *multipart.FileHeader
  File multipart.File
}

type CreateProductImageDto struct {
  File FileHandler
  User_id uuid.UUID
  Product_id string 
}

type CreateProductBody struct {
  Name string `db:"name"`
  Product_id string `db:"product_id"`
  Url string `db:"url"`
}
