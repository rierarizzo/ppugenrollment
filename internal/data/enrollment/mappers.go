package enrollment

import (
	"ppugenrollment/internal/domain"
)

func fromModelToEnrollmentApplication(model *ApplicationModel) domain.EnrollmentApplication {
	return domain.EnrollmentApplication{
		ID: model.ID,
		Student: domain.User{
			ID: model.Student,
		},
		Project: domain.Project{
			ID: model.Project,
		},
		Schedule:  model.Schedule,
		AppliedOn: model.AppliedOn,
	}
}

func fromEnrollmentApplicationToModel(application *domain.EnrollmentApplication) ApplicationModel {
	return ApplicationModel{
		ID:       application.ID,
		Student:  application.Student.ID,
		Project:  application.Project.ID,
		Schedule: application.Schedule,
	}
}
