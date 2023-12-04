package enrollment

import "ppugenrollment/internal/domain"

func fromApplicationToResponse(application *domain.EnrollmentApplication) ApplicationResponse {
	return ApplicationResponse{
		ID:       application.ID,
		Student:  application.Student,
		Project:  application.Project,
		Schedule: application.Schedule,
	}
}

func fromRequestToApplication(request *ApplicationRequest) domain.EnrollmentApplication {
	return domain.EnrollmentApplication{
		Student:  request.Student,
		Project:  request.Project,
		Schedule: request.Schedule,
	}
}
