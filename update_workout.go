package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func updateWorkout(w http.ResponseWriter, r *http.Request) {

	/// ----- ID validation first ----

	idString := r.URL.Query().Get("id")

	// and does not have any issues
	// we do that by converting to string
	// then we convert from string to integer

	// use the ozzo package to conduct usual id validation checks:

	// 1. no id given
	if err := validation.Validate(
		idString,
		validation.Required.Error("ID is required"),
	); err != nil {
		writeJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 2. id is invalid i.e. like non integer
	id, err := strconv.Atoi(idString)
	if err != nil {
		writeJSONError(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// 3. id is not 0 or negative
	if err := validation.Validate(
		id,
		validation.Min(1).Error("ID must be greater than 0"),
	); err != nil {
		writeJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// ----- Decoding the input ------

	// if all the id checks pass, then create a struct and decode the id into the struct
	var workout Workout

	err = json.NewDecoder(r.Body).Decode(&workout)
	if err != nil {
		writeJSONError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// ----- Checking if workout actually exists ------

	// now we check that the workout actually exists based on the workout id
	err = db.QueryRow(
		`SELECT id FROM workouts WHERE id = 1 and deleted_at IS NULL`,
		id,
	).Scan(&workout.ID)

	// check

	if errors.Is(err, sql.ErrNoRows) {
		writeJSONError(w, "Workout does not exist", http.StatusNotFound)
		return
	}

	// ------ Checking the updated workout like we would a new workout entry ---
	// ----------- name validation ----------------

	// validate that name is not empty
	err = validation.Validate(workout.Name, validation.Required)
	if err != nil {
		writeJSONError(w, "Name not given", http.StatusBadRequest)
		return
	}

	// ----------- date validation ----------------

	// make sure it is not empty
	if workout.Date == "" {
		writeJSONError(w, "Workout date is required", http.StatusBadRequest)
		return
	}

	// Parse the date
	workoutDate, err := time.Parse("2006-01-02", workout.Date)
	if err != nil {
		writeJSONError(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	// Make sure the date is not greater than today
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	if workoutDate.After(today) {
		writeJSONError(w, "Workout date can not be in the future", http.StatusBadRequest)
		return
	}

	// once all is said and done, let's update the workout
	result, err := db.Exec(
		`UPDATE workouts 
		SET 
		name = $1,
		workout_date = $2, 
		created_at = $3,
		updated_at = $4,
		sets = $5,
		reps = $6`,
		workout.Name,
		workout.Date,
		workout.CreatedAt,
		time.Now(),
		workout.Sets,
		workout.Reps,
	)

	if err != nil {
		writeJSONError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Check how many rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		writeJSONError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// no matching workout existed
	if rowsAffected == 0 {
		writeJSONError(w, "Workout not found", http.StatusNotFound)
		return
	}

	// keep the ID from the url
	workout.ID = id

	// send updated transaction back as json
	json.NewEncoder(w).Encode(workout)
}