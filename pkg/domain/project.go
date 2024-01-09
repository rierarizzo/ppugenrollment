package domain

import "time"

type Project struct {
	ID          int
	Company     Company
	Name        string
	Description string
	Schedules   []string
	Starts      time.Time
	Ends        time.Time
}
