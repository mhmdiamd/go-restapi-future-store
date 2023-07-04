package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/lib/helpers"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/response"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product response.CreateProductRequest) (Product, error) {
	id := helpers.GenerateUUID()
	query := "INSERT INTO product(id, product_name, stock, price, description, photo, id_category) VALUES($1, $2, $3, $4, $5, $6, $7)"
	rows, err := tx.QueryContext(ctx, query,
		id,
		product.Product_name,
		product.Stock,
		product.Price,
		product.Description,
		product.Photo,
		product.Id_category,
	)
	helpers.PanicIfError(err)
	defer rows.Close()

	productResponse := Product{
		Id:           id,
		Product_name: product.Product_name,
		Stock:        product.Stock,
		Price:        product.Price,
		Description:  product.Description,
		Photo:        product.Photo,
		Id_category:  product.Id_category,
	}

	return productResponse, nil
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId uuid.UUID) (Product, error) {
	query := "SELECT * FROM product WHERE id = $1 LIMIT 1"
	rows, err := tx.QueryContext(ctx, query, productId)
	helpers.PanicIfError(err)
	defer rows.Close()

	productResponse := Product{}

	if rows.Next() {
		err := rows.Scan(&productResponse.Id, &productResponse.Product_name, &productResponse.Stock, &productResponse.Price, &productResponse.Description, &productResponse.Id_category, &productResponse.Photo)
		helpers.PanicIfError(err)

		return productResponse, nil
	} else {
		return productResponse, errors.New("product not found")
	}
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product Product) (Product, error) {
	query := "UPDATE product SET product_name = $2, stock = $3, price = $4, description = $5, photo = $6, id_category = $7 WHERE id = $1"
	rows, err := tx.QueryContext(ctx, query,
		product.Id,
		product.Product_name,
		product.Stock,
		product.Price,
		product.Description,
		product.Photo,
		product.Id_category,
	)
	helpers.PanicIfError(err)
	defer rows.Close()

	productResponse := Product(product)

	return productResponse, nil
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, productId uuid.UUID) {
	query := "DELETE FROM product WHERE id = $1"
	rows, err := tx.QueryContext(ctx, query, productId)
	helpers.PanicIfError(err)
	defer rows.Close()
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []Product {
	query := "SELECT * FROM product LIMIT 10"
	rows, err := tx.QueryContext(ctx, query)
	helpers.PanicIfError(err)
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var data = Product{}
		err := rows.Scan(
			&data.Id,
			&data.Product_name,
			&data.Stock,
			&data.Price,
			&data.Description,
			&data.Id_category,
			&data.Photo,
		)

		helpers.PanicIfError(err)
		products = append(products, data)
	}

	return products
}
