package ports

import "ppugenrollment/pkg/domain"

type ProjectEnroller interface {
	EnrollToProject(application *domain.EnrollmentApplication, enrolledBy int) (*domain.EnrollmentApplication,
		*domain.AppError)
}

type EnrollmentRepository interface {
	InsertEnrollment(application *domain.EnrollmentApplication) (int, *domain.AppError)
}
