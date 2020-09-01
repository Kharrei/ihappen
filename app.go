package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (i *App) Initialize() {
	connectionString := "postgres://postgres:admin007@localhost/ihappen?sslmode=disable"
	var err error
	i.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	i.Router = mux.NewRouter()

	i.initializeRoutes()
}

func (i *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8010", i.Router))
}
