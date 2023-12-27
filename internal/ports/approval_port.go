package ports

import "ppugenrollment/pkg/domain"

type EnrollmentApprover interface {
	ApproveEnrollmentApplication(applicationID, approvedBy int) (*domain.EnrollmentGenerated, *domain.AppError)
}

type ApprovalRepository interface {
	InsertEnrollmentApproval(applicationID, approvedBy int) (*domain.EnrollmentGenerated, *domain.AppError)
}
