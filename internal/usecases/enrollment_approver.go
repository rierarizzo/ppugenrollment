package usecases

import (
	"log/slog"
	"ppugenrollment/internal/domain"
	"ppugenrollment/internal/ports"
)

type DefaultApprover struct {
	approvalRepo ports.ApprovalRepository
}

func NewEnrollmentApprover(approvalRepo ports.ApprovalRepository) *DefaultApprover {
	return &DefaultApprover{approvalRepo}
}

// ApproveEnrollmentApplication approves an enrollment application with the given application ID by the specified approver ID.
// It returns the approved enrollment generated data and any error that occurred during the approval process.
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
