package main

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *sql.DB

// First I define the data structure db as an sql variable DB
func connectDB() {
	// Load .env locally. Render has own env.
	if err := godotenv.Load(); err != nil { // this looks for a file called .env
		log.Println("No .env file found; using environment variables")
	}

	// Database URL assignment needs to be in the function connectDB
	// Read the complete PostgreSQL connection URL
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	// Connect to PostgreSQL. This is where we initialize the details of the database from the database URL using pgx.
	var err error // creating a new data type error called err. we are creating a new one instead of just using var, err like befrore because we need our old db assignment to remain
	db, err = sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatal("failed to open databaseL ", err)
	}

	// Verify that PostgreSQL is reachable
	if err = db.Ping(); err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
}
