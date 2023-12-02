package project_enroller

import "ppugenrollment/internal/domain"

type Enroller interface {
	EnrollToProject(application *domain.EnrollmentApplication) (*domain.EnrollmentApplication, *domain.AppError)
}

type EnrollmentRepository interface {
	InsertEnrollment(application *domain.EnrollmentApplication) (int, *domain.AppError)
}
