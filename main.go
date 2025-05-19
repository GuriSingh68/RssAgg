package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env") // Load .env file if because without it the env variables will not be set
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT environment variable is not set")
	}

	//Create a new router using chi
	router := chi.NewRouter()
	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%s", portString),
	}
	log.Printf("Starting server on port %s ....", portString)
	err := srv.ListenAndServe() // Start the server and handle http requests
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	fmt.Println("PORT:", portString)
}
