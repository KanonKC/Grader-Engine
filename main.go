package main

import (
	controllers_grader "ModelGrader-Grader/controllers/grader"
	services_grader "ModelGrader-Grader/services/grader"
	services_sandbox "ModelGrader-Grader/services/sandbox"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!\n"))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the server!\n"))
}

func main() {
	sandboxSvc := services_sandbox.New(8)
	sandboxSvc.Init()
	graderSvc := services_grader.New(sandboxSvc)
	graderCtrl := controllers_grader.New(graderSvc)
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/", rootHandler)
	// Make post method
	mux.HandleFunc("/output", graderCtrl.GenerateOutput)

	server := http.Server{
		Addr:    ":8081",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
