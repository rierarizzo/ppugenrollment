package domain

import "time"

type EnrollmentApplication struct {
	ID           int
	Student      User
	Project      Project
	Schedule     int
	ScheduleCode string
	AppliedOn    time.Time
	Status       string
}
