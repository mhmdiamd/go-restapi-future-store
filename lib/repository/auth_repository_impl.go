package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-restapi-future-store/helper"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/response"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
	"golang.org/x/crypto/bcrypt"
)

var (
  UnauthorizedError = errors.New("Unauthorized")
)

type AuthRepositoryImpl struct {
}

func NewAuthRepositoryImpl() AuthRepository {
	return &AuthRepositoryImpl{}
}

func (r *AuthRepositoryImpl) SignUp(c context.Context, tx *sqlx.Tx, body response.SignUpRequestDTO) (domain.User, error) {
	id := helper.GenerateUUID()
	hashedPassword, err := helper.GenerateHashPassword(body.Password)
	if err != nil {
		helper.PanicIfError(err)
	}

	query := "INSERT INTO users(id, name, email, password, role) VALUES($1, $2, $3, $4, $5) RETURNING id, name, email, role"
	row := tx.QueryRowContext(c, query, id, body.Name, body.Email, hashedPassword, body.Role)

	result := domain.User{}
	err = row.Scan(&result.Id, &result.Name, &result.Email, &result.Role)

	helper.PanicIfError(err)

	fmt.Println(result)

	return result, nil
}

func (r *AuthRepositoryImpl) SignIn(c context.Context, tx *sqlx.Tx, body domain.Authentication) (domain.User, error) {
	email := body.Email
  password := []byte(body.Password)

	query := `SELECT * FROM users WHERE email=$1 LIMIT 1`

	var user domain.User
	err := tx.GetContext(c, &user, query, email)

  if err != nil {
		return user, UnauthorizedError
	}

  // Check password is match or not
  hashedPassword := &user.Password
  err = bcrypt.CompareHashAndPassword([]byte(*hashedPassword), password)

  if err != nil {
		return user, UnauthorizedError
	}

	helper.PanicIfError(err)

	return user, nil
}
