package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"authenticator/data"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const WEBPORT = "3000"

var COUNTS int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authenticator service on port:", WEBPORT)

	// TODO: connect to DB
	conn := connectToDB()
	if conn == nil {
		log.Panic("Could not connect to Postgres. Exiting...")
	}

	// set up config
	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", WEBPORT),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready...:", err)
			COUNTS++
		} else {
			log.Println("Postgres is ready!")
			return connection
		}
		if COUNTS > 10 {
			log.Println("Postgres not ready after 10 tries. Exiting...")
			return nil
		}

		log.Println("Waiting 2 seconds before trying again...")
		time.Sleep(2 * time.Second)
		continue
	}
}
