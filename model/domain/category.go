package domain

import "github.com/google/uuid"

type Category struct {
	Id   uuid.UUID
	Name string
}
