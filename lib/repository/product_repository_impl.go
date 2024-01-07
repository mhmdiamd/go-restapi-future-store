package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/helper"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/dto"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sqlx.Tx, body dto.ProductCreateDTO) (domain.Product, error) {
	id := helper.GenerateUUID()
  query := `
    INSERT INTO products(id, name, stock, price, description, color, size, condition, id_category, id_user)
    VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
    RETURNING id, name, stock, price, description, color, size, condition, id_category, id_user, created_at, updated_at
  `
	result := tx.QueryRowxContext(ctx, query, 
    id, body.Name, body.Stock, body.Price, body.Description,
    body.Color, body.Size, body.Condition, body.Id_category, body.Id_user)

  var newProduct domain.Product

  err := result.Scan(
  &newProduct.ID, &newProduct.Name, &newProduct.Stock, &newProduct.Price,
    &newProduct.Description, &newProduct.Color, &newProduct.Size,
    &newProduct.Condition, &newProduct.Id_category, &newProduct.Id_user, &newProduct.Created_at, &newProduct.Updated_at,
  )

	helper.PanicIfError(err)

  fmt.Println(newProduct)
	return newProduct, nil
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sqlx.Tx, productId uuid.UUID) (domain.Product, error) {
  var result domain.Product

  query := "SELECT * FROM products WHERE id = $1 LIMIT 1"
	err := tx.GetContext(ctx, &result, query)
	helper.PanicIfError(err)

  if err != nil {
    return result, errors.New("Product not found")
  }

  return result, nil
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sqlx.Tx, product Product) domain.Product {
  var newProduct domain.Product
  
  query := "UPDATE products SET product_name = $2, stock = $3, price = $4, description = $5, photo = $6, id_category = $7 WHERE id=$1"
	err := tx.GetContext(ctx, &newProduct, query,
		product.Id,
		product.Product_name,
		product.Stock,
		product.Price,
		product.Description,
		product.Photo,
		product.Id_category,
	)

	helper.PanicIfError(err)

	return newProduct
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sqlx.Tx, productId uuid.UUID) {
	query := "DELETE FROM products WHERE id = $1"
	rows, err := tx.QueryContext(ctx, query, productId)
	helper.PanicIfError(err)
	defer rows.Close()
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sqlx.Tx) []domain.Product {
  query := `
    SELECT 
      products.*, 
      CASE
        WHEN COUNT(images_product.id) > 0 THEN
            JSONB_AGG(
                JSONB_BUILD_OBJECT(
                    'id', images_product.id,
                    'name', images_product.name,
                    'url', images_product.url,
                    'product_id', products.id
                )
            )
        ELSE
            '[]'::jsonb
    END AS photo
    FROM products LEFT JOIN images_product ON products.id = images_product.product_id
    GROUP BY products.id 
    LIMIT 10
  `
	rows, err := tx.QueryxContext(ctx, query)
	helper.PanicIfError(err)

  var products []domain.Product

  for rows.Next() {
    product := domain.Product{}

    err := rows.Scan(
      &product.ID,
      &product.Name,
      &product.Description,
      &product.Price,
      &product.Color,
      &product.Size,
      &product.Stock,
      &product.Id_category,
      &product.Id_user,
      &product.Condition,
      &product.Created_at,
      &product.Updated_at,
      &product.Photo,
    )

    if err != nil {
      helper.PanicIfError(err)
    }

    products = append(products, product)
  }

	return products
}

func (r *ProductRepositoryImpl) FindAllByIdSeller(ctx context.Context, tx *sqlx.Tx, id_seller uuid.UUID) []domain.Product {
  
  query := `
    SELECT 
      products.*, 
      CASE
        WHEN COUNT(images_product.id) > 0 THEN
            JSONB_AGG(
                JSONB_BUILD_OBJECT(
                    'id', images_product.id,
                    'name', images_product.name,
                    'url', images_product.url,
                    'product_id', products.id
                )
            )
        ELSE
            '[]'::jsonb
    END AS photo
    FROM products LEFT JOIN images_product ON products.id = images_product.product_id
    WHERE id_user=$1 
    GROUP BY products.id 
    LIMIT 10
  `
	rows, err := tx.QueryxContext(ctx, query, id_seller)
	helper.PanicIfError(err)

  var products []domain.Product

  for rows.Next() {
    product := domain.Product{}

    err := rows.Scan(
      &product.ID,
      &product.Name,
      &product.Description,
      &product.Price,
      &product.Color,
      &product.Size,
      &product.Stock,
      &product.Id_category,
      &product.Id_user,
      &product.Condition,
      &product.Created_at,
      &product.Updated_at,
      &product.Photo,
    )

    if err != nil {
      helper.PanicIfError(err)
    }

    products = append(products, product)
  }

	return products
}


