package domain

import "time"

type EnrollmentGenerated struct {
	ID                    int
	EnrollmentApplication int
	Project               Project
	ApprovedBy            Approver
	GeneratedAt           time.Time
}
