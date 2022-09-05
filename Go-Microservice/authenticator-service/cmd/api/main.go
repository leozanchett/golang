package main

import (
	"database/sql"
	"log"
	"net/http"

	"authenticator/data"
)

const WEBPORT = "81"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authenticator service on port:", WEBPORT)

	// TODO: connect to DB

	// set up config
	app := Config{}

	srv := &http.Server{
		Addr:    ":" + WEBPORT,
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
