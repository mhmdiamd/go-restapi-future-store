package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/exceptions"
	"github.com/mhmdiamd/go-restapi-future-store/helpers"

	web "github.com/mhmdiamd/go-restapi-future-store/model/web/category"
	"github.com/mhmdiamd/go-restapi-future-store/src/category/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           validator.Validate
}

func NewCategoryService(repository repository.CategoryRepository, db *sql.DB, validate validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: repository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, req web.CategoryCreateRequest) (web.CategoryResponse, error) {

	// Validation Request
	errValidation := service.Validate.Struct(req)
	helpers.PanicIfError(errValidation)

	// Create Transaction
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	// Check the transaction after all of code is execution
	defer helpers.CommitOrRollback(tx)

	response, _ := service.CategoryRepository.Create(ctx, tx, req.Name)

	fmt.Println(response)

	return helpers.ConvertToCategoryResponse(response), nil
}

func (service *CategoryServiceImpl) Update(ctx context.Context, req web.CategoryUpdateRequest) (web.CategoryResponse, error) {
	errValidation := service.Validate.Struct(req)
	helpers.PanicIfError(errValidation)

	// Create transaction
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	response, err := service.CategoryRepository.FindById(ctx, tx, req.Id)
	helpers.PanicIfError(err)

	response.Name = req.Name
	updatedCategory := service.CategoryRepository.Update(ctx, tx, response)

	defer helpers.CommitOrRollback(tx)
	return helpers.ConvertToCategoryResponse(updatedCategory), nil
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId uuid.UUID) {
	// Create Transaction
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	service.CategoryRepository.Delete(ctx, tx, categoryId)
	defer helpers.CommitOrRollback(tx)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId uuid.UUID) web.CategoryResponse {
	// Create Transaction
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	return helpers.ConvertToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	// Create Transaction
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helpers.ConvertToSliceCategoryResponse(categories)
}
