package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "81"

type Config struct{}

func main() {
	app := Config{}
	log.Println("Starting the broker:", webPort)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
