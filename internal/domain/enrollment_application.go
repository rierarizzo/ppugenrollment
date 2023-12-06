package domain

import "time"

type EnrollmentApplication struct {
	ID        int
	Student   Student
	Project   Project
	Schedule  int
	AppliedOn time.Time
}
