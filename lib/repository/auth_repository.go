package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/response"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type AuthRepository interface {
	SignUp(c context.Context, tx *sqlx.Tx, body response.SignUpRequestDTO) (domain.User, error)
	SignIn(c context.Context, tx *sqlx.Tx, body domain.Authentication) (domain.User, error)
}
