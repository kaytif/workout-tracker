package main

import "net/http"

// workoutHandler sends each HTTP method to the correct handler
func handlerWorkout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		getWorkout(w, r)

	case http.MethodPost:
		createWorkout(w, r)

	default:
		writeJSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}