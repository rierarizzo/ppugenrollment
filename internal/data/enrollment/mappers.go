package enrollment

import "ppugenrollment/internal/domain"

func fromModelToEnrollmentApplication(model *ApplicationModel) domain.EnrollmentApplication {
	return domain.EnrollmentApplication{
		ID:       model.ID,
		Student:  model.Student,
		Project:  model.Project,
		Schedule: model.Schedule,
	}
}

func fromEnrollmentApplicationToModel(application *domain.EnrollmentApplication) ApplicationModel {
	return ApplicationModel{
		ID:       application.ID,
		Student:  application.Student,
		Project:  application.Project,
		Schedule: application.Schedule,
	}
}
