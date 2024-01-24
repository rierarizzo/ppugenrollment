package mappers

import (
	"ppugenrollment/internal/api/types"
	"ppugenrollment/pkg/domain"
)

func FromApplicationToResponse(application *domain.EnrollmentApplication) types.EnrollmentApplicationResponse {
	return types.EnrollmentApplicationResponse{
		ID:             application.ID,
		Student:        application.Student.ID,
		StudentName:    application.Student.Name,
		StudentSurname: application.Student.Surname,
		Project:        application.Project.ID,
		ProjectName:    application.Project.Name,
		CompanyName:    application.Project.Company.Name,
		Schedule:       application.Schedule,
		Status:         application.Status,
	}
}

func FromApplicationsToResponse(applications []domain.EnrollmentApplication) []types.EnrollmentApplicationResponse {
	var response []types.EnrollmentApplicationResponse

	for _, application := range applications {
		response = append(response, FromApplicationToResponse(&application))
	}

	return response
}

func FromRequestToApplication(request *types.EnrollmentApplicationRequest) domain.EnrollmentApplication {
	return domain.EnrollmentApplication{
		Project:  domain.Project{ID: request.Project},
		Schedule: request.Schedule,
	}
}
