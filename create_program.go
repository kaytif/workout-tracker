package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func createProgram(w http.ResponseWriter, r *http.Request) {

	// Create Program var to store incoming program in
	var program Program
	
	// first we need to decode the program
	err := json.NewDecoder(r.Body).Decode(&program)

	// check for any json errors
	if err != nil {
		writeJSONError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// ------ name validation --------------
	
	// 1. validate that name is not empty
	err = validation.Validate(
		strings.TrimSpace(program.Name),
		validation.Required,
	)

	if err != nil {
		writeJSONError(w, "Name not given", http.StatusBadRequest)
		return
	}

	err = db.QueryRow(
		`INSERT 
		INTO programs
		(name)
		VALUES ($1)
		RETURNING id`,
		program.Name,
	).Scan(&program.ID)

	// id validation not needed since postgresql generates it for us.

	if err != nil {
		log.Println("Database insert error:", err)
		writeJSONError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// set the header 
	w.Header().Set("content-Type", "application/json")

	// Workout successfully created
	w.WriteHeader(http.StatusCreated)

	// Send created workout back as JSON
	json.NewEncoder(w).Encode(program)
	}

	