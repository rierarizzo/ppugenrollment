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

// EnrollToProject is a method of the DefaultEnroller type that allows enrolling a student to a project.
// It takes an enrollment application and the ID of the user performing the enrollment.
// It updates the application with the enrolled student information and creates a default response.
// It inserts the enrollment application into the enrollment repository and retrieves the last inserted ID.
// If there is an error during the insertion, it returns the default response with a new AppError.
// Finally, it sets the ID of the default response and returns it along with a nil error.
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
