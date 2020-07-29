package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App main API
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Init method for the core application
func (a *App) Init(user, password, dbname string) {}

// Run application on given address
func (a *App) Run(addr string) {}
