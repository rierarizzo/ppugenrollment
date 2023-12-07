package approval

import (
	"ppugenrollment/internal/domain"
)

func fromGeneratedToResponse(generated *domain.EnrollmentGenerated) EnrollmentGeneratedResponse {
	return EnrollmentGeneratedResponse{
		ID:                    generated.ID,
		EnrollmentApplication: generated.EnrollmentApplication,
		Project: projectResponse{
			ID: generated.Project.ID,
			Company: companyResponse{
				ID:   generated.Project.Company.ID,
				Name: generated.Project.Company.Name,
				RUC:  generated.Project.Company.RUC,
			},
			Description: generated.Project.Description,
			Schedule:    generated.Project.Schedule,
			Starts:      generated.Project.Starts,
			Ends:        generated.Project.Ends,
		},
		ApprovedBy:  approvedByResponse{},
		GeneratedAt: generated.GeneratedAt,
	}
}
