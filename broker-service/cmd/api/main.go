package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "8081"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting server on %s\n", port)

	srv := &http.Server{Addr: fmt.Sprintf(":%s", port), Handler: app.routes()}

	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}

}
