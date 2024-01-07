package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/helper"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (c *CategoryRepositoryImpl) Create(ctx context.Context, tx *sqlx.Tx, name string) (Category, error) {
	id := helper.GenerateUUID()
	query := "INSERT INTO categories VALUES($1, $2)"
	rows, err := tx.QueryContext(ctx, query, id, name)
	helper.PanicIfError(err)

	defer rows.Close()

	var result = Category{
		Id:   id,
		Name: name,
	}

	return result, nil
}

func (c *CategoryRepositoryImpl) Update(ctx context.Context, tx *sqlx.Tx, category Category) Category {

	fmt.Println(category)
	query := "UPDATE categories SET name = $1 WHERE id = $2"
	rows, err := tx.QueryContext(ctx, query, category.Name, category.Id)

	helper.PanicIfError(err)
	defer rows.Close()

	return category
}

func (c *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sqlx.Tx, categoryId uuid.UUID) error {
	query := "DELETE FROM categories WHERE id = $1"
	rows, err := tx.QueryContext(ctx, query, categoryId)

	helper.PanicIfError(err)
	defer rows.Close()

	return nil
}

func (c *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sqlx.Tx, categoryId uuid.UUID) (Category, error) {
	query := "SELECT * FROM categories WHERE id = $1"
	rows, err := tx.QueryContext(ctx, query, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	var result = Category{}
	if rows.Next() {
		err := rows.Scan(&result.Id, &result.Name)
		helper.PanicIfError(err)

		return result, nil
	} else {
		return result, errors.New("category not found")
	}
}

func (c *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sqlx.Tx) []Category {
	query := "SELECT * FROM categories LIMIT 10"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []Category

	for rows.Next() {
		category := Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		categories = append(categories, category)
	}

	return categories
}
