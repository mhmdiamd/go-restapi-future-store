package response

import "github.com/google/uuid"

type CategoryResponse struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type CategoryUpdateRequest struct {
	Id   uuid.UUID
	Name string
}
type CategoryCreateRequest struct {
	Name string `validate:"required"`
}
