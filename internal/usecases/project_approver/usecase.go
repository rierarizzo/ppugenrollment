package project_approver

import (
	"ppugenrollment/internal/domain"
)

type DefaultApprover struct {
	approvalRepo ApprovalRepository
}

func New(approvalRepo ApprovalRepository) *DefaultApprover {
	return &DefaultApprover{approvalRepo}
}

func (d *DefaultApprover) ApproveEnrollmentApplication(applicationID, approvedBy int) (
	*domain.EnrollmentGenerated,
	*domain.AppError) {
	generatedID, appErr := d.approvalRepo.ApproveEnrollmentApplication(applicationID, approvedBy)
	if appErr != nil {
		return nil, domain.NewAppError(appErr, domain.UnexpectedError)
	}

	generated, appErr := d.approvalRepo.SelectEnrollmentGenerated(generatedID)
	if appErr != nil {
		return nil, domain.NewAppError(appErr, domain.UnexpectedError)
	}

	return generated, nil
}
