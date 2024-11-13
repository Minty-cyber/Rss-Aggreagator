package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	// "github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello world")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not available")
	}

	router := chi.NewRouter()

	server := &http.Server{
		Handler: router,
		Addr: ":" + portString,

	}

	log.Printf("Starting server on port %v", portString)
	err := server.ListenAndServe()
	if err !=nil {
		log.Fatal(err)

	}

	fmt.Println("Port:", portString)
}