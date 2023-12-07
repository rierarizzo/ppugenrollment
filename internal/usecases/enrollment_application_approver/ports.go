package enrollment_application_approver

import "ppugenrollment/internal/domain"

type Approver interface {
	ApproveEnrollmentApplication(applicationID, approvedBy int) (*domain.EnrollmentGenerated, *domain.AppError)
}

type ApprovalRepository interface {
	ApproveEnrollmentApplication(applicationID, approvedBy int) (int, *domain.AppError)
	SelectEnrollmentGenerated(generatedID int) (*domain.EnrollmentGenerated, *domain.AppError)
}
