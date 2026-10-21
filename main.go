package main

import (
    "log"
    "net/http"
    "os"
)

func main() {
    // Connect to PostgreSQL
    connectDB()

    // Use port 8080 if PORT is not configured
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Register workout endpoint
    http.Handle("/workouts", http.HandlerFunc(handlerWorkout))

    // Start HTTP server
    log.Println("Server running on port", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}