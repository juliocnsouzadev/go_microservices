package main

import (
	"broker/cmd/api/config"
	"fmt"
	"log"
	"net/http"
)

const (
	webPort = "8080"
)

func main() {
	app := config.Config{}

	log.Printf("Starting broker service on port %s\n", webPort)

	//define http server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.Routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
