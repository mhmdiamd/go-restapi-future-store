package web

import "github.com/google/uuid"

type CategoryUpdateRequest struct {
	Id   uuid.UUID
	Name string
}
