package main

import (
	"ModelGrader-Grader/routes"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	routes.SetupRoutes(mux)

	server := http.Server{
		Addr:    ":8081",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
