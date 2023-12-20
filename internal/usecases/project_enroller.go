package usecases

import (
	"ppugenrollment/internal/domain"
	"ppugenrollment/internal/ports"
)

type DefaultEnroller struct {
	enrollmentRepo ports.EnrollmentRepository
}

func NewProjectEnroller(enrollmentRepo ports.EnrollmentRepository) *DefaultEnroller {
	return &DefaultEnroller{enrollmentRepo}
}

func (p *DefaultEnroller) EnrollToProject(application *domain.EnrollmentApplication, enrolledBy int) (
	*domain.EnrollmentApplication,
	*domain.AppError) {
	application.Student = domain.User{ID: enrolledBy}

	defaultResponse := domain.EnrollmentApplication{
		Student:  application.Student,
		Project:  application.Project,
		Schedule: application.Schedule,
	}

	lastID, appErr := p.enrollmentRepo.InsertEnrollment(application)
	if appErr != nil {
		return &defaultResponse, domain.NewAppError(appErr, domain.UnexpectedError)
	}

	defaultResponse.ID = lastID

	return &defaultResponse, nil
}
