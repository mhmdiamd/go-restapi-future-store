package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/helpers"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
	webProduct "github.com/mhmdiamd/go-restapi-future-store/model/web/product"
	"github.com/mhmdiamd/go-restapi-future-store/src/product/repository"
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

func (service *ProductServiceImpl) Create(ctx context.Context, product webProduct.CreateProductRequest) (webProduct.ProductResponse, error) {

	// do Validation First
	validatedErr := service.Validate.Struct(product)
	helpers.PanicIfError(validatedErr)

	// Create Transaction First
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	response, err := service.repository.Save(ctx, tx, product)
	helpers.PanicIfError(err)

	return helpers.ConvertToProductResponse(response), nil
}

func (service *ProductServiceImpl) FindById(ctx context.Context, productId uuid.UUID) webProduct.ProductResponse {
	// Create transactions
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	product, err := service.repository.FindById(ctx, tx, productId)
	helpers.PanicIfError(err)

	return webProduct.ProductResponse(product)
}

func (service *ProductServiceImpl) Update(ctx context.Context, product webProduct.UpdateProductRequest) (webProduct.ProductResponse, error) {

	// Create Transactions
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	productUpdated, err := service.repository.Update(ctx, tx, domain.Product(product))
	helpers.PanicIfError(err)

	return helpers.ConvertToProductResponse(productUpdated), nil
}

func (service *ProductServiceImpl) Delete(ctx context.Context, productId uuid.UUID) {

	// Create Transactions
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	service.repository.Delete(ctx, tx, productId)
	defer helpers.CommitOrRollback(tx)
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []webProduct.ProductResponse {

	// Create Transactions
	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	products := service.repository.FindAll(ctx, tx)

	return helpers.ConvertToSliceProductResponse(products)
}
