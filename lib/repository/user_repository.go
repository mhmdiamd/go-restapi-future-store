package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/dto"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type UserRepository interface {
  GetProfile(c context.Context, tx *sqlx.Tx, id_user uuid.UUID) (domain.User, error)
  Update(c context.Context, tx *sqlx.Tx, body dto.UserUpdateDTO) (domain.User, error) 
}
