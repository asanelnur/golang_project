package main

import (
	"database/sql"
	"flag"
	"online-book-store/pkg/book/model"
	"online-book-store/pkg/jsonlog"
	"online-book-store/pkg/vcs"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var (
	version = vcs.Version()
)

type config struct {
	port int
	env  string
	fill bool
	db   struct {
		dsn string
	}
}

type application struct {
	config config
	models model.Models
	logger *jsonlog.Logger
	wg     sync.WaitGroup
}

func main() {
	var cfg config
	flag.BoolVar(&cfg.fill, "fill", false, "Fill database with dummy data")
	flag.IntVar(&cfg.port, "port", 8081, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://bookstore:pa55word@localhost/bookstore?sslmode=disable", "PostgreSQL DSN")
	flag.Parse()

	logger := jsonlog.NewLogger(os.Stdout, jsonlog.LevelInfo)

	db, err := openDB(cfg)
	if err != nil {
		logger.PrintError(err, nil)
		return
	}

	defer func() {
		if err := db.Close(); err != nil {
			logger.PrintFatal(err, nil)
		}
	}()

	logger.PrintInfo("database connection pool established", nil)

	app := &application{
		config: cfg,
		models: model.NewModels(db),
		logger: logger,
	}

	if err := app.serve(); err != nil {
		logger.PrintFatal(err, nil)
	}
}

// func (app *application) run() {
// 	r := mux.NewRouter()

// 	v1 := r.PathPrefix("/api/v1").Subrouter()

// 	v1.HandleFunc("/book", app.getBookList).Methods("GET")
// 	v1.HandleFunc("/category", app.createCategoryHandler).Methods("POST")
// 	v1.HandleFunc("/category/{categoryId:[0-9]+}", app.getCategoryHandler).Methods("GET")
// 	v1.HandleFunc("/category/{categoryId:[0-9]+}", app.updateCategoryHandler).Methods("PUT")
// 	v1.HandleFunc("/category/{categoryId:[0-9]+}", app.deleteCategoryHandler).Methods("DELETE")
// 	v1.HandleFunc("/users", app.registerUserHandler).Methods("POST")
// 	v1.HandleFunc("/users/activated", app.activateUserHandler).Methods("PUT")
// 	v1.HandleFunc("/tokens/authentication", app.createAuthenticationTokenHandler).Methods("POST")

// 	log.Printf("Starting server on %s\n", app.config.port)
// 	err := http.ListenAndServe(app.config.port, r)
// 	log.Fatal(err)
// }

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
