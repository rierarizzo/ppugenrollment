package project

import "time"

type Response struct {
	ID          int       `json:"id"`
	Company     int       `json:"company"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Starts      time.Time `json:"starts"`
	Ends        time.Time `json:"ends"`
}
