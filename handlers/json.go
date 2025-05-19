package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	object "github.com/GuriSingh68/rssagg/Object"
)

// Respond with a json response
func respondWithJson(w http.ResponseWriter, status int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error marshalling json response:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
	log.Println("Response:", string(data))
}

// Respond with error
func respondWithError(w http.ResponseWriter, status int, message string) {
	if status >= 500 {
		log.Println("Error:", message)
	}
	respondWithJson(w, status, object.ErrorResponse{Error: message})
}
