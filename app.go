package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
)

// App main API
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Init method for the core application
func (a *App) Init(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable",
			username, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

// Run application on given address
func (a *App) Run(addr string) {}
