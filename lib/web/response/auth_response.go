package response

import (
	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/lib"
)

type SignUpResponse struct {
	Id          uuid.UUID      `json:"id"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	Birth_date  lib.NullTime   `json:"birth_date,omitempty"`
	Store_name  lib.NullString `json:"store_name,omitempty"`
	Description lib.NullString `json:"description,omitempty"`
	Role        string         `json:"role"`
	Phone       lib.NullInt64  `json:"phone"`
	Address     lib.NullString `json:"address,omitempty"`
	Photo       lib.NullString `json:"photo,omitempty"`
}

type SignUpRequestDTO struct {
	Name     string
	Email    string `validate:'required'`
	Password string
	Role     string
}

type SignInRequestDTO struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

type SignInResponse struct {
	User  SignUpResponse `json:"user"`
	Token string         `json:"token"`
}
