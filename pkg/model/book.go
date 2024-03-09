package model

import (
	"database/sql"
	"log"
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
