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
			Schedule:    generated.Schedule,
			Description: generated.Project.Description,
			Starts:      generated.Project.Starts,
			Ends:        generated.Project.Ends,
		},
		ApprovedBy: approvedByResponse{
			ID:           generated.ApprovedBy.ID,
			IDCardNumber: generated.ApprovedBy.IDCardNumber,
			Name:         generated.ApprovedBy.Name,
			Surname:      generated.ApprovedBy.Surname,
		},
		GeneratedAt: generated.GeneratedAt,
	}
}
