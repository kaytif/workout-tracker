package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func getWorkout(w http.ResponseWriter, r *http.Request) {

	idString := r.URL.Query().Get("id")
	
	// convert from string to integer
	id, err := strconv.Atoi(idString)

	// Ask postgresql for all account
	rows, err := db.Query(
		`SELECT 
		workouts.name, 
		workouts.id, 
		workouts.date,
		workouts.set, 
		workouts.reps,
		FROM workouts
		WHERE DELETED_AT IS NULL
		AND workouts.id = $1`,
		id,
	)

	if err != nil {
		writeJSONError(w, "Internal server error", http.StatusInternalServerError)
	}

	defer rows.Close()

	// Start with an empty slice to empty empty results []
	workouts := []Workout{}

	// Move throw each row postgressql; returned
	for rows.Next() {
		var workoutStore Workout
		err := rows.Scan(&workoutStore.ID, &workoutStore.Name, &workoutStore.Sets, &workoutStore.Reps, &workoutStore.Date)
		if err != nil {
			writeJSONError(w, "Internal server error", http.StatusInternalServerError)
			return
		
		}

		// Add this account to our list.
		workouts = append(workouts, workoutStore)
	}


	// Check whether any error occurred while iterating through any rows
	if err := rows.Err(); err != nil {
		writeJSONError(w, "Internal server error", http.StatusInternalServerError)
		return	
	}

	// send completed response back to client
	json.NewEncoder(w).Encode(workouts)
}

