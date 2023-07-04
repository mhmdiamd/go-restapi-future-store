package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/lib/helpers"
	"github.com/mhmdiamd/go-restapi-future-store/lib/helpers/converts"
	"github.com/mhmdiamd/go-restapi-future-store/lib/repository"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/response"
)

type ProductServiceImpl struct {
	repository repository.ProductRepository
	DB         *sql.DB
	Validate   validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, Db *sql.DB, validate validator.Validate) ProductService {
	return &ProductServiceImpl{
		repository: productRepository,
		DB:         Db,
		Validate:   validate,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, product response.ProductCreateRequest) (response.ProductResponse, error) {

	// do Validation First
	validatedErr := service.Validate.Struct(product)
	helpers.PanicIfError(validatedErr)

	// Create Transaction First
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	response, err := service.repository.Save(ctx, tx, product)
	helpers.PanicIfError(err)

	return converts.ConvertToProductResponse(response), nil
}

func (service *ProductServiceImpl) FindById(ctx context.Context, productId uuid.UUID) response.ProductResponse {
	// Create transactions
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	product, err := service.repository.FindById(ctx, tx, productId)
	helpers.PanicIfError(err)

	return response.ProductResponse(product)
}

func (service *ProductServiceImpl) Update(ctx context.Context, product response.UpdateProductRequest) (response.ProductResponse, error) {

	// Create Transactions
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	productUpdated, err := service.repository.Update(ctx, tx, repository.Product(product))
	helpers.PanicIfError(err)

	return converts.ConvertToProductResponse(productUpdated), nil
}

func (service *ProductServiceImpl) Delete(ctx context.Context, productId uuid.UUID) {

	// Create Transactions
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	service.repository.Delete(ctx, tx, productId)
	defer helpers.CommitOrRollback(tx)
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []response.ProductResponse {

	// Create Transactions
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	products := service.repository.FindAll(ctx, tx)

	return converts.ConvertToSliceProductResponse(products)
}
