package enrollment

import "ppugenrollment/internal/domain"

func fromApplicationToResponse(application *domain.EnrollmentApplication) ApplicationResponse {
	return ApplicationResponse{
		ID:       application.ID,
		Student:  application.Student.ID,
		Project:  application.Project.ID,
		Schedule: application.Schedule,
	}
}

func fromRequestToApplication(request *ApplicationRequest) domain.EnrollmentApplication {
	return domain.EnrollmentApplication{
		Student:  domain.User{ID: request.Student},
		Project:  domain.Project{ID: request.Project},
		Schedule: request.Schedule,
	}
}
