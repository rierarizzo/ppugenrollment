package mappers

import (
	"ppugenrollment/internal/api/types"
	"ppugenrollment/pkg/domain"
)

func FromGeneratedToResponse(generated *domain.EnrollmentGenerated) types.EnrollmentGeneratedResponse {
	return types.EnrollmentGeneratedResponse{
		ID:                    generated.ID,
		EnrollmentApplication: generated.EnrollmentApplication,
		Project: types.ProjectResponse{
			ID: generated.Project.ID,
			Company: types.CompanyResponse{
				ID:   generated.Project.Company.ID,
				Name: generated.Project.Company.Name,
				RUC:  generated.Project.Company.RUC,
			},
			Schedules:   generated.Project.Schedules,
			Description: generated.Project.Description,
			Starts:      generated.Project.Starts,
			Ends:        generated.Project.Ends,
		},
		ApprovedBy: types.ApprovedByResponse{
			ID:           generated.ApprovedBy.ID,
			IDCardNumber: generated.ApprovedBy.IDCardNumber,
			Name:         generated.ApprovedBy.Name,
			Surname:      generated.ApprovedBy.Surname,
		},
		GeneratedAt: generated.GeneratedAt,
	}
}
