package project_enroller

import "ppugenrollment/internal/domain"

type DefaultEnroller struct {
	enrollmentRepo EnrollmentRepository
}

func New(enrollmentRepo EnrollmentRepository) *DefaultEnroller {
	return &DefaultEnroller{enrollmentRepo}
}

func (p *DefaultEnroller) EnrollToProject(application *domain.EnrollmentApplication) (
	*domain.EnrollmentApplication,
	*domain.AppError) {
	defaultResponse := domain.EnrollmentApplication{
		ID:       0,
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
