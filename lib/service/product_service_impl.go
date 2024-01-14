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
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/dto"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/response"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type ProductServiceImpl struct {
	repository repository.ProductRepository
	DB         *sqlx.DB
	Validate   *validator.Validate
}

func NewProductServiceImpl(productRepository repository.ProductRepository, Db *sqlx.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		repository: productRepository,
		DB:         Db,
		Validate:   validate,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, product dto.ProductCreateDTO) response.ProductCreateResponse {

	// do Validation First
	validatedErr := service.Validate.Struct(product)
	helper.PanicIfError(validatedErr)

	// Create Transaction First
	tx := service.DB.MustBegin()

	defer helper.CommitOrRollback(tx)

	data, err := service.repository.Save(ctx, tx, product)
	helper.PanicIfError(err)

  response := response.ProductCreateResponse{
    ID: data.ID,
    Name: data.Name,
    Condition: data.Condition,
    Size: data.Size,
    Color: data.Color,
    Photo: data.Photo,
    Description: data.Description,
    Price: data.Price,
    Stock: data.Price,
    Id_category: data.Id_category,
    Id_user: data.Id_user,
    Created_at: data.Created_at,
    Updated_at: data.Updated_at,
  }
  

	return response
}

func (service *ProductServiceImpl) FindById(ctx context.Context, productId uuid.UUID) domain.Product {
	// Create transactions
	tx := service.DB.MustBegin()

	defer helper.CommitOrRollback(tx)

	product, err := service.repository.FindById(ctx, tx, productId)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return product
}

func (service *ProductServiceImpl) Update(ctx context.Context, product response.UpdateProductRequest) domain.Product {

	// Create Transactions
	tx := service.DB.MustBegin()

	defer helper.CommitOrRollback(tx)

	_, errFind := service.repository.FindById(ctx, tx, product.Id)

	if errFind != nil {
		panic(exception.NewNotFoundError(errFind.Error()))
	}

	fmt.Println(product.Id)

	productUpdated := service.repository.Update(ctx, tx, repository.Product(product))

	return productUpdated
}

func (service *ProductServiceImpl) Delete(ctx context.Context, productId uuid.UUID) {

	// Create Transactions
	tx := service.DB.MustBegin()

	service.repository.Delete(ctx, tx, productId)
	defer helper.CommitOrRollback(tx)
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []response.ProductResponse {

	// Create Transactions
	tx := service.DB.MustBegin()

  products := service.repository.FindAll(ctx, tx)

	return converts.ConvertToSliceProductResponse(products)
}

func (service *ProductServiceImpl) FindAllByIdSeller(ctx context.Context, id_seller uuid.UUID) []response.ProductResponse {

  tx := service.DB.MustBegin()

  result := service.repository.FindAllByIdSeller(ctx, tx, id_seller)

  return converts.ConvertToSliceProductResponse(result)
}



