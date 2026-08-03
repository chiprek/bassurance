package main

import (
	"time"

	"github.com/google/uuid"
)

type Jobs struct {
	ID         uuid.UUID `json:"id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Name       string    `json:"name"`
	Status     string    `json:"status,omitempty"`
}

type Unit struct {
	ID            uuid.UUID `json:"id"`
	Job_ID        uuid.UUID `json:"job_id"`
	Serial_number string    `json:"serial_number"`
	Status        string    `json:"status"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
}
