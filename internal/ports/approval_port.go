package ports

import "ppugenrollment/pkg/domain"

type EnrollmentApprover interface {
	ApproveEnrollmentApplication(applicationID, approvedBy int, observation string) (*domain.EnrollmentGenerated, *domain.AppError)
}

type ApprovalRepository interface {
	InsertEnrollmentApproval(applicationID, approvedBy int, observation string) (*domain.EnrollmentGenerated, *domain.AppError)
}
