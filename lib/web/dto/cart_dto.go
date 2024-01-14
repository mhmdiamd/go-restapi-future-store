package dto

import "github.com/google/uuid"

type IncrementCart struct {
  Id_cart uuid.UUID
}

type DecrementCart struct {
  Id_cart uuid.UUID
}
