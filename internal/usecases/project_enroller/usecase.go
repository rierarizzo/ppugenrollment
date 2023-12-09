package project_enroller

import "ppugenrollment/internal/domain"

type DefaultEnroller struct {
	enrollmentRepo EnrollmentRepository
}

func New(enrollmentRepo EnrollmentRepository) *DefaultEnroller {
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
