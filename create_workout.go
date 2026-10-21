package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func createWorkout(w http.ResponseWriter, r *http.Request) {

	// create an empty account
	var newWorkout Workout

	// first decode the request
	err := json.NewDecoder(r.Body).Decode(&newWorkout)

	if err != nil {
		writeJSONError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// ----------- name validation ----------------

	// validate that name is not empty
	err = validation.Validate(newWorkout.Name, validation.Required)
	if err != nil {
		writeJSONError(w, "Name not given", http.StatusBadRequest)
		return
	}

	// ----------- date validation ----------------

	// make sure it is not empty
	if newWorkout.Date == "" {
		writeJSONError(w, "Workout date is required", http.StatusBadRequest)
		return
	}

	// Parse the date
	// workoutDate, err := time.Parse("2006-01-02", newWorkout.Date)
	// if err != nil {
	// 	writeJSONError(w, "Invalid date format", http.StatusBadRequest)
	//	return
	// }

	// Make sure the date is not greater than today
	// now := time.Now()
	// today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	// if workoutDate.After(today) {
	//	writeJSONError(w, "Workout date can not be in the future", http.StatusBadRequest)
	//	return
	//}

	// now inserting into postgre sql. we need to scan the rows and get the id
	// this is because the frontend might need to use it
	err = db.QueryRow(
		"INSERT INTO workouts (name, workout_date, created_at, updated_at, program_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		newWorkout.Name,
		newWorkout.Date,
		time.Now(),
		newWor,
		newWorkout.Sets,
	).Scan(&newWorkout.ID)

	if err != nil {
		log.Println("Database insert error:", err)
		writeJSONError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Expense succesfull created
	w.WriteHeader(http.StatusCreated)

	// send created expense back as JSON
	json.NewEncoder(w).Encode(newWorkout)

}
