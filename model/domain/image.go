package domain

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/helper"
)

type ProductImage struct {
  Id uuid.UUID `db:"id" json:"id"`
  Name string `db:"name" json:"name"`
  Url string `db:"url" json:"url"`
  Product_id string `db:"product_id" json:"product_id"`
}

type JSONBArray []ProductImage

func (j *JSONBArray) Scan(value interface{}) error {
  if value == nil {
    *j = nil
    return nil
  }

  // fmt.Println(value)

  // check if the value is a []byte
  bytes, ok := value.([]byte)
  if !ok {
    return fmt.Errorf("Error: Scan Error : expected []byte got ", value)
  }

  // Check if the bytes represent NULL
  if string(bytes) == "null" {
      *j = nil
      return nil
  }

  // Unmarshal then JSONB array into the custom type
  err := json.Unmarshal(bytes, j)
  helper.PanicIfError(err)

  return nil
}


