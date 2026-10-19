package main

import (
	"time"
)

type Workout struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Date      string     `json:"date"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Sets      int        `json:"sets"`
	Reps      int        `json:"reps"`
}
