package service

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/go-playground/validator/v10"
	"github.com/mhmdiamd/go-restapi-future-store/app"
	"github.com/mhmdiamd/go-restapi-future-store/exception"
	"github.com/mhmdiamd/go-restapi-future-store/helper"
	"github.com/mhmdiamd/go-restapi-future-store/lib/repository"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/response"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	DB             *sqlx.DB
	Validate       *validator.Validate
}

func NewAuthServiceImpl(r repository.AuthRepository, db *sqlx.DB, v *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: r,
		DB:             db,
		Validate:       v,
	}
}

func (service *AuthServiceImpl) SignUp(c context.Context, body response.SignUpRequestDTO) response.SignUpResponse {

	// Validation Request
	errValidation := service.Validate.Struct(body)
	helper.PanicIfError(errValidation)

	// Create database transactin
	tx := service.DB.MustBegin()

	// Rollback if have error on transaction
	defer helper.CommitOrRollback(tx)

	res, err := service.AuthRepository.SignUp(c, tx, body)
	helper.PanicIfError((err))

	result := response.SignUpResponse{
		Id:          res.Id,
		Name:        res.Name,
		Email:       res.Email,
		Birth_date:  res.Birth_date,
		Phone:       res.Phone,
		Photo:       res.Photo,
		Role:        res.Role,
		Address:     res.Address,
		Store_name:  res.Store_name,
		Description: res.Description,
	}

	return result
}

func (service *AuthServiceImpl) SignIn(c context.Context, body response.SignInRequestDTO) response.SignInResponse {

	// Validation Request
	errValidation := service.Validate.Struct(body)
	helper.PanicIfError(errValidation)

	// Create database transactin
	tx := service.DB.MustBegin()

	// Rollback if have error on transaction
	defer helper.CommitOrRollback(tx)

	res, err := service.AuthRepository.SignIn(c, tx, domain.Authentication(body))

	if err != nil {
		if err.Error() == "Unauthorized" {
			panic(exception.NewErrorUnauthorized(err.Error()))
		} else {
			helper.PanicIfError(err)
		}
	}

	convertedRes := response.SignUpResponse{
		Id:          res.Id,
		Name:        res.Name,
		Phone:       res.Phone,
		Photo:       res.Photo,
		Email:       res.Email,
		Role:        res.Role,
		Address:     res.Address,
		Birth_date:  res.Birth_date,
		Store_name:  res.Store_name,
		Description: res.Description,
	}

  // Generate Token

  // Create Payload
  tokenPayload := app.JWTPayload{
    ID : convertedRes.Id,
    Email: convertedRes.Email,
    Role: convertedRes.Role,
  }

  token, err := app.GenerateJWT(tokenPayload)
  helper.PanicIfError(err)

	result := response.SignInResponse{
		User:  convertedRes,
		Token: token,
	}

	return result
}
