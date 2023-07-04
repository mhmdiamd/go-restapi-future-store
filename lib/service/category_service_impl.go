package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/lib/exceptions"
	"github.com/mhmdiamd/go-restapi-future-store/lib/helpers"
	"github.com/mhmdiamd/go-restapi-future-store/lib/helpers/converts"
	"github.com/mhmdiamd/go-restapi-future-store/lib/repository"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/response"
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

func (service *CategoryServiceImpl) Create(ctx context.Context, req response.CategoryCreateRequest) (response.CategoryResponse, error) {

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

	return converts.ConvertToCategoryResponse(response), nil
}

func (service *CategoryServiceImpl) Update(ctx context.Context, req response.CategoryUpdateRequest) (response.CategoryResponse, error) {
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
	return converts.ConvertToCategoryResponse(updatedCategory), nil
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId uuid.UUID) {
	// Create Transaction
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	service.CategoryRepository.Delete(ctx, tx, categoryId)
	defer helpers.CommitOrRollback(tx)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId uuid.UUID) response.CategoryResponse {
	// Create Transaction
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	return converts.ConvertToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []response.CategoryResponse {
	// Create Transaction
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return converts.ConvertToSliceCategoryResponse(categories)
}
