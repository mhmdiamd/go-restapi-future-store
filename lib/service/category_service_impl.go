package service

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/exception"
	"github.com/mhmdiamd/go-restapi-future-store/helper"
	"github.com/mhmdiamd/go-restapi-future-store/helper/converts"
	"github.com/mhmdiamd/go-restapi-future-store/lib/repository"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/response"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sqlx.DB
	Validate           *validator.Validate
}

func NewCategoryService(repository repository.CategoryRepository, db *sqlx.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: repository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, req response.CategoryCreateRequest) (response.CategoryResponse, error) {

	// Validation Request
	errValidation := service.Validate.Struct(req)
	helper.PanicIfError(errValidation)

	// Create Transaction
	tx := service.DB.MustBegin()

	// Check the transaction after all of code is execution
	defer helper.CommitOrRollback(tx)

	response, _ := service.CategoryRepository.Create(ctx, tx, req.Name)

	fmt.Println(response)

	return converts.ConvertToCategoryResponse(response), nil
}

func (service *CategoryServiceImpl) Update(ctx context.Context, req response.CategoryUpdateRequest) (response.CategoryResponse, error) {
	errValidation := service.Validate.Struct(req)
	helper.PanicIfError(errValidation)

	// Create transaction
	tx := service.DB.MustBegin()

	response, err := service.CategoryRepository.FindById(ctx, tx, req.Id)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	response.Name = req.Name
	updatedCategory := service.CategoryRepository.Update(ctx, tx, response)

	defer helper.CommitOrRollback(tx)
	return converts.ConvertToCategoryResponse(updatedCategory), nil
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId uuid.UUID) {
	// Create Transaction
	tx := service.DB.MustBegin()

	service.CategoryRepository.Delete(ctx, tx, categoryId)
	defer helper.CommitOrRollback(tx)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId uuid.UUID) response.CategoryResponse {
	// Create Transaction
	tx := service.DB.MustBegin()
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return converts.ConvertToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []response.CategoryResponse {
	// Create Transaction
	tx := service.DB.MustBegin()

	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	fmt.Println(categories)
	return converts.ConvertToSliceCategoryResponse(categories)
}
