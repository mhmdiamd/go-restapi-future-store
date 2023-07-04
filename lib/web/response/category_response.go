package response

import "github.com/google/uuid"

type CategoryResponse struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
