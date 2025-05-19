package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GuriSingh68/rssagg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
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
	//Add CORS configuration
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	//Hooking up the routes
	v1Router := chi.NewRouter()
	router.Mount("/v1", v1Router)                       // Mount the v1 router to the main router the request will be like /v1/ready
	v1Router.Get("/error", handlers.ErrorHandler)       // Error handler endpoint
	v1Router.Get("/healthz", handlers.ReadinessHandler) // Health check endpoint

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
