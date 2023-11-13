package domain

import "time"

type EnrollmentGenerated struct {
	ID                  int
	EnrollmentGenerated int
	ApprovedBy          int
	GeneratedAt         time.Time
}
