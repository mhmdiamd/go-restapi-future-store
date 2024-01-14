package domain

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/helper"
)

type Product struct {
  ID           uuid.UUID `db:"id" json:"id"`
  Name string `db:"name" json:"name"`
  Description  string `db:"description" json:"description"`
	Stock        int `db:"stock" json:"stock"`
	Price        int `db:"price" json:"price"`
	Color        string `db:"color" json:"color"`
  Size string `db:"size" json:"size"`
  Condition string `db:"condition" json:"condition"`
	Photo        JSONBArray `db:"photo" json:"photo"`
	Id_category  uuid.UUID `db:"id_category" json:"id_category"`
	Id_user  uuid.UUID `db:"id_user" json:"id_user"`
  Created_at time.Time `db:"created_at" json:"created_at"`
  Updated_at time.Time `db:"updated_at" json:"updated_at"`
}

type JSONBProduct Product

func (j *JSONBProduct) Scan(value interface{}) error {
	if value == nil {
		j = nil
		return nil
	}

	// Periksa apakah nilai tersebut adalah []byte
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Error: Scan Error: expected []byte, got %v", value)
	}

	// Periksa apakah bytes tersebut merepresentasikan NULL
	if string(bytes) == "null" {
		j = nil
		return nil
	}

	// Unmarshal array JSONB ke dalam tipe kustom
	var product Product
	err := json.Unmarshal(bytes, &product)
	if err != nil {
		helper.PanicIfError(err)
	}

	*j = JSONBProduct(product)
	return nil
}
