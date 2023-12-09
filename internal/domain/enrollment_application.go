package domain

import "time"

type EnrollmentApplication struct {
	ID        int
	Student   User
	Project   Project
	Schedule  int
	AppliedOn time.Time
}
