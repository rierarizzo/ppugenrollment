package mappers

import (
	"ppugenrollment/internal/data/models"
	"ppugenrollment/pkg/domain"
)

func FromModelToEnrollmentGenerated(model *models.EnrollmentGeneratedModel) domain.EnrollmentGenerated {
	return domain.EnrollmentGenerated{
		ID:                    model.ID,
		EnrollmentApplication: model.EnrollmentApplication,
		Project: domain.Project{
			ID: model.ProjectID,
			Company: domain.Company{
				ID:   model.CompanyID,
				Name: model.CompanyName,
				RUC:  model.CompanyRUC,
			},
			Description: model.ProjectDescription,
			Starts:      model.ProjectStarts,
			Ends:        model.ProjectEnds,
		},
		Schedule: model.ProjectSchedule,
		ApprovedBy: domain.User{
			ID:           model.ApproverID,
			IDCardNumber: model.ApproverIDCardNumber,
			Name:         model.ApproverName,
			Surname:      model.ApproverSurname,
		},
		GeneratedAt: model.GeneratedAt,
	}
}
