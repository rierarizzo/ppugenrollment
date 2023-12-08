package domain

import "time"

type Project struct {
	ID          int
	Company     Company
	Description string
	Starts      time.Time
	Ends        time.Time
}
