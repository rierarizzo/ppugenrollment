package mappers

import (
	"ppugenrollment/internal/data/models"
	"ppugenrollment/internal/domain"
)

func FromEnrollmentApplicationToModel(application *domain.EnrollmentApplication) models.EnrollmentApplicationModel {
	return models.EnrollmentApplicationModel{
		ID:       application.ID,
		Student:  application.Student.ID,
		Project:  application.Project.ID,
		Schedule: application.Schedule,
	}
}
