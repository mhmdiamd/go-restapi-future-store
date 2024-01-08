package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mhmdiamd/go-restapi-future-store/helper"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/dto"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type ImageRepositoryImpl struct {
}

func NewImageRepositoryImpl() ImageRepository {
  return &ImageRepositoryImpl{}
}

func (r *ImageRepositoryImpl) UploadProductImage(ctx context.Context, tx *sqlx.Tx, body dto.CreateProductBody) (domain.ProductImage, error) {
  id := helper.GenerateUUID()

  fmt.Println(body)
  
	query := `INSERT INTO images_product VALUES($1, $2, $3, $4) RETURNING id, name, url, product_id`
  row := tx.QueryRowContext(ctx, query, id, body.Name, body.Url, body.Product_id)

  var newProductImage = domain.ProductImage{}

  err := row.Scan(
    &newProductImage.Id,
    &newProductImage.Name,
    &newProductImage.Product_id,
    &newProductImage.Url,
  )

  if err != nil {
    return newProductImage, err
  }

	return newProductImage, nil
}

func (r *ImageRepositoryImpl) UpdateProductImage(ctx context.Context, tx *sqlx.Tx, body dto.UpdateProductBody) (domain.ProductImage, error) {

  query := `UPDATE images_product SET name=$2 WHERE id=$1 RETURNING id, name, url, product_id`
  row := tx.QueryRowContext(ctx, query, body.Id_product_image, body.Name)

  var newProductImage = domain.ProductImage{}

  err := row.Scan(
    &newProductImage.Id,
    &newProductImage.Name,
    &newProductImage.Product_id,
    &newProductImage.Url,
  )

  if err != nil {
    return newProductImage, err
  }

	return newProductImage, nil
}

func (r *ImageRepositoryImpl) GetAllByIdProduct(ctx context.Context, tx *sqlx.Tx, id_product uuid.UUID) []domain.ProductImage {
  query := "SELECT * FROM images_product WHERE id_product=$1"
  rows, err := tx.QueryContext(ctx, query, id_product)

  helper.PanicIfError(err)

  var images []domain.ProductImage

  for rows.Next() {
    image := domain.ProductImage{}

    err := rows.Scan(
      &image.Id,
      &image.Name,
      &image.Url,
    )

    helper.PanicIfError(err)

    images = append(images, image)
  }

  return images
}











