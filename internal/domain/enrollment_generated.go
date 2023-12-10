package domain

import "time"

type EnrollmentGenerated struct {
	ID                    int
	EnrollmentApplication int
	Schedule              string
	Project               Project
	ApprovedBy            User
	GeneratedAt           time.Time
}
