package usecases

import (
	"log/slog"
	"ppugenrollment/internal/ports"
	"ppugenrollment/pkg/domain"
)

type DefaultEnrollmentApprover struct {
	approvalRepo ports.ApprovalRepository
}

func NewEnrollmentApprover(approvalRepo ports.ApprovalRepository) *DefaultEnrollmentApprover {
	return &DefaultEnrollmentApprover{approvalRepo}
}

// ApproveEnrollmentApplication approves an enrollment application with the given application ID by the specified approver ID.
// It returns the approved enrollment generated data and any error that occurred during the approval process.
func (d *DefaultEnrollmentApprover) ApproveEnrollmentApplication(applicationID, approvedBy int) (
	*domain.EnrollmentGenerated,
	*domain.AppError) {
	generated, appErr := d.approvalRepo.InsertEnrollmentApproval(applicationID, approvedBy)

	if appErr != nil {
		slog.Error(appErr.Error())
		return nil, domain.NewAppErrorWithType(domain.UnexpectedError)
	}

	return generated, nil
}
