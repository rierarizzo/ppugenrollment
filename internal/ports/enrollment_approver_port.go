package ports

import "ppugenrollment/internal/domain"

type Approver interface {
	ApproveEnrollmentApplication(applicationID, approvedBy int) (*domain.EnrollmentGenerated, *domain.AppError)
}

type ApprovalRepository interface {
	ApproveEnrollmentApplication(applicationID, approvedBy int) (*domain.EnrollmentGenerated, *domain.AppError)
}