package main

import (
	"encoding/json"
	"net/http"
)

func sendError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode) // Set the HTTP status code

	// Create a map for the error message
	errorResponse := map[string]string{"error": message}

	// Marshal the map to a JSON string
	jsonResponse, err := json.Marshal(errorResponse)
	if err != nil {
		// If there is an error marshaling, send a simple error message
		http.Error(w, "Error creating the error message", http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
