package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"online-book-store/pkg/model"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type config struct {
	port string
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config config
	models model.Models
}

func (app *application) respondWithError(w http.ResponseWriter, code int, message string) {
	app.respondWithJSON(w, code, map[string]string{"error": message})
}

func (app *application) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (app *application) createCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	category := &model.Category{
		Title:       input.Title,
		Description: input.Description,
	}

	err = app.models.Categories.Insert(category)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusCreated, category)
}

func (app *application) getCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["categoryId"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	category, err := app.models.Categories.Get(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	app.respondWithJSON(w, http.StatusOK, category)
}

func (app *application) updateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["categoryId"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	category, err := app.models.Categories.Get(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	var input struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if input.Title != nil {
		category.Title = *input.Title
	}

	if input.Description != nil {
		category.Description = *input.Description
	}

	err = app.models.Categories.Update(category)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, category)
}

func (app *application) deleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["categoryId"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	err = app.models.Categories.Delete(id)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var cfg config
	flag.StringVar(&cfg.port, "port", ":8081", "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://bookstore:pa55word@localhost/bookstore?sslmode=disable", "PostgreSQL DSN")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	logger.Printf("database connection pool established")

	app := &application{
		config: cfg,
		models: model.NewModels(db),
	}

	app.run()
}

func (app *application) run() {
	r := mux.NewRouter()

	v1 := r.PathPrefix("/api/v1").Subrouter()

	v1.HandleFunc("/category", app.createCategoryHandler).Methods("POST")
	v1.HandleFunc("/category/{categoryId:[0-9]+}", app.getCategoryHandler).Methods("GET")
	v1.HandleFunc("/category/{categoryId:[0-9]+}", app.updateCategoryHandler).Methods("PUT")
	v1.HandleFunc("/category/{categoryId:[0-9]+}", app.deleteCategoryHandler).Methods("DELETE")

	log.Printf("Starting server on %s\n", app.config.port)
	err := http.ListenAndServe(app.config.port, r)
	log.Fatal(err)
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
