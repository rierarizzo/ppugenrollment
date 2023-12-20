package mappers

import (
	"ppugenrollment/internal/api/types"
	"ppugenrollment/internal/domain"
)

func FromApplicationToResponse(application *domain.EnrollmentApplication) types.EnrollmentApplicationResponse {
	return types.EnrollmentApplicationResponse{
		ID:       application.ID,
		Student:  application.Student.ID,
		Project:  application.Project.ID,
		Schedule: application.Schedule,
	}
}

func FromRequestToApplication(request *types.EnrollmentApplicationRequest) domain.EnrollmentApplication {
	return domain.EnrollmentApplication{
		Student:  domain.User{ID: request.Student},
		Project:  domain.Project{ID: request.Project},
		Schedule: request.Schedule,
	}
}
