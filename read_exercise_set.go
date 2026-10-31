package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

func WorkoutExerciseResponse(w http.ResponseWriter, *http.Request) {
	var workoutExerciseResponse WorkoutExerciseResponse
	
	
	// select relevant stuff from SQL
	// you need to join

	results, err := db.Query(`
    	SELECT
        	id,
        	workout_exercise_id,
        	set_number,
        	reps,
        	weight
    	FROM exercise_sets
		`)

	if err != nil {
		writeJSONError(w, "Internal Server Error", http.StatusInternalServerError)

	if errors.Is(err, rowsAffected) {
		writeJSONError()
	}


	}





}