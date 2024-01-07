package service

import (
	"context"

	"github.com/mhmdiamd/go-restapi-future-store/lib/web/response"
)

type AuthService interface {
  SignUp(c context.Context, body response.SignUpRequestDTO) response.SignUpResponse
  SignIn(c context.Context, body response.SignInRequestDTO) response.SignInResponse
}
