package main

import (
	"log"
	"net/http"
	"os"
)


func main() {
	// Connect to PostgreSQL.
	connectDB()

	// Use 8080 
	port := os.Getenv("PORT")
	if port == " " {
		port = "8080"
}

}