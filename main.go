package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not available")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(
			cors.Options{
			
			AllowedOrigins:   []string{"https://*", "http://*"},
			// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
	  }))

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