package main

import (
	"encoding/json"
	"net/http"
)

// WriteJSONError sends all API erros in a consistent JSON format.
func writeJSONError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}
