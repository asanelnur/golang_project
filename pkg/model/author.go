package model

import (
	"database/sql"
	"log"
)

type Author struct {
	Id        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Name      string `json:"name"`
	Info      string `json:"info"`
	Age       int    `json:"age"`
}

type AuthorModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}
