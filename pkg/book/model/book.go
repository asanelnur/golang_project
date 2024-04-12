package model

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Book struct {
	Id          string `json:"id"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	Author      string `json:"author"`
}

type BookModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (m BookModel) GetAll(title string, from, to int, filters Filters) ([]*Book, Metadata, error) {

	// Retrieve all book items from the database.
	query := fmt.Sprintf(
		`
		SELECT count(*) OVER(), id, created_at, updated_at, title, description, price, category, author
		FROM books
		WHERE (LOWER(title) = LOWER($1) OR $1 = '')
		AND (price >= $2 OR $2 = 0)
		AND (price <= $3 OR $3 = 0)
		ORDER BY %s %s, id ASC
		LIMIT $4 OFFSET $5
		`,
		filters.sortColumn(), filters.sortDirection())

	// Create a context with a 3-second timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Organize our four placeholder parameter values in a slice.
	args := []interface{}{title, from, to, filters.limit(), filters.offset()}

	// log.Println(query, title, from, to, filters.limit(), filters.offset())
	// Use QueryContext to execute the query. This returns a sql.Rows result set containing
	// the result.
	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	// Importantly, defer a call to rows.Close() to ensure that the result set is closed
	// before GetAll returns.
	defer func() {
		if err := rows.Close(); err != nil {
			m.ErrorLog.Println(err)
		}
	}()

	// Declare a totalRecords variable
	totalRecords := 0

	var books []*Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&totalRecords, &book.Id, &book.CreatedAt, &book.UpdatedAt, &book.Title, &book.Description, &book.Price, &book.Category, &book.Author)
		if err != nil {
			return nil, Metadata{}, err
		}

		// Add the Movie struct to the slice
		books = append(books, &book)
	}

	// When the rows.Next() loop has finished, call rows.Err() to retrieve any error
	// that was encountered during the iteration.
	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	// Generate a Metadata struct, passing in the total record count and pagination parameters
	// from the client.
	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	// If everything went OK, then return the slice of the movies and metadata.
	return books, metadata, nil
}
