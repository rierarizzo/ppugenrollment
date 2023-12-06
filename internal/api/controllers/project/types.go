package project

import "time"

type Response struct {
	ID          int       `json:"id"`
	Company     int       `json:"company"`
	Description string    `json:"description"`
	Schedule    string    `json:"schedule"`
	Starts      time.Time `json:"starts"`
	Ends        time.Time `json:"ends"`
}
