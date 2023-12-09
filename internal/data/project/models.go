package project

import "time"

type Model struct {
	ID          int       `db:"id"`
	Company     int       `db:"company"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Starts      time.Time `db:"starts"`
	Ends        time.Time `db:"ends"`
}
