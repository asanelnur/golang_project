package model

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type Category struct {
	Id          string `json:"id"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CategoryModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (m CategoryModel) Insert(category *Category) error {
	query := `
		INSERT INTO categories (title, description) 
		VALUES ($1, $2) 
		RETURNING id, created_at, updated_at
		`
	args := []interface{}{category.Title, category.Description}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&category.Id, &category.CreatedAt, &category.UpdatedAt)
}

func (m CategoryModel) Get(id int) (*Category, error) {
	query := `
		SELECT id, created_at, updated_at, title, description
		FROM categories
		WHERE id = $1
		`
	var category Category
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&category.Id, &category.CreatedAt, &category.UpdatedAt, &category.Title, &category.Description)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (m CategoryModel) Update(category *Category) error {
	query := `
		UPDATE categories
		SET title = $1, description = $2
		WHERE id = $3
		RETURNING updated_at
		`
	args := []interface{}{category.Title, category.Description, category.Id}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&category.UpdatedAt)
}

func (m CategoryModel) Delete(id int) error {
	query := `
		DELETE FROM categories
		WHERE id = $1
		`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, id)
	return err
}
