package enrollment_application_approver

import (
	"log/slog"
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
	generated, appErr := d.approvalRepo.ApproveEnrollmentApplication(applicationID, approvedBy)
	if appErr != nil {
		slog.Error(appErr.Error())
		return nil, domain.NewAppError(appErr, domain.UnexpectedError)
	}

	return generated, nil
}
