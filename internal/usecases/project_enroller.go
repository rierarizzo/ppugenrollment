package usecases

import (
	"log/slog"
	"ppugenrollment/internal/ports"
	"ppugenrollment/pkg/domain"
)

type DefaultProjectEnroller struct {
	enrollmentRepo ports.EnrollmentRepository
}

func NewProjectEnroller(enrollmentRepo ports.EnrollmentRepository) *DefaultProjectEnroller {
	return &DefaultProjectEnroller{enrollmentRepo}
}

func (p *DefaultProjectEnroller) EnrollToProject(application *domain.EnrollmentApplication,
	enrolledBy int) (*domain.EnrollmentApplication, *domain.AppError) {
	application.Student = domain.User{ID: enrolledBy}

	lastID, appErr := p.enrollmentRepo.InsertEnrollment(application)

	if appErr != nil {
		slog.Error(appErr.Error())
		return nil, domain.NewAppErrorWithType(domain.UnexpectedError)
	}

	completeApplication := domain.EnrollmentApplication{
		ID:       lastID,
		Student:  application.Student,
		Project:  application.Project,
		Schedule: application.Schedule,
	}

	return &completeApplication, nil
}

func (d *DefaultProjectEnroller) GetEnrollmentApplications() ([]domain.EnrollmentApplication, *domain.AppError) {
	applications, appErr := d.enrollmentRepo.SelectEnrollmentApplications()

	if appErr != nil {
		slog.Error(appErr.Error())
		return nil, domain.NewAppErrorWithType(domain.UnexpectedError)
	}

	return applications, nil
}
