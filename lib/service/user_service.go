package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/dto"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/response"
)

type UserService interface {
  Update(c context.Context, body dto.UserUpdateDTO) (response.UserResponse, error)
  GetProfile(c context.Context, user_id uuid.UUID) (response.UserResponse, error)
}
