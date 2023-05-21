package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/helpers"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (c *CategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, name string) (domain.Category, error) {
	id := helpers.GenerateUUID()
	query := "INSERT INTO category VALUES($1, $2)"
	rows, err := tx.QueryContext(ctx, query, id, name)
	helpers.PanicIfError(err)

	defer rows.Close()

	var result = domain.Category{
		Id:   id,
		Name: name,
	}

	return result, nil
}

func (c *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {

	fmt.Println(category)
	query := "UPDATE category SET name = $1 WHERE id = $2"
	rows, err := tx.QueryContext(ctx, query, category.Name, category.Id)

	helpers.PanicIfError(err)
	defer rows.Close()

	return category
}

func (c *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, categoryId uuid.UUID) error {
	query := "DELETE FROM category WHERE id = $1"
	rows, err := tx.QueryContext(ctx, query, categoryId)

	helpers.PanicIfError(err)
	defer rows.Close()

	return nil
}

func (c *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId uuid.UUID) (domain.Category, error) {
	query := "SELECT * FROM category WHERE id = $1"
	rows, err := tx.QueryContext(ctx, query, categoryId)
	helpers.PanicIfError(err)
	defer rows.Close()

	var result = domain.Category{}
	if rows.Next() {
		err := rows.Scan(&result.Id, &result.Name)
		helpers.PanicIfError(err)

		return result, nil
	} else {
		return result, errors.New("category not found")
	}
}

func (c *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	query := "SELECT * FROM category LIMIT 10"
	rows, err := tx.QueryContext(ctx, query)
	helpers.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category

	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helpers.PanicIfError(err)

		categories = append(categories, category)
	}

	return categories
}
