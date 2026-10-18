package main

import (
	"encoding/json"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
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

	// now we have to run validation checks


	// validate that name is not empty
	err = validation.Validate(newWorkout.Name, validation.Required)
	if err != nil {
		writeJSONError(w, "Name not given", http.StatusBadRequest)
		return
	}

	// now inserting into postgre sql. we need to scan the rows and get the id
	// this is because the frontend might need to use it
	err = db.QueryRow(
		"INSERT INTO transactions (account_id, name, amount, transaction_date, category_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		newWorkout.ID,
		newWorkout.Name,
		newWorkout.Amount,
		newWorkout.CategoryID,
	).Scan(&newTransaction.ID)

	if err != nil {
		writeJSONError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}


}

