package approval

import (
	"ppugenrollment/internal/domain"
)

func fromModelToEnrollmentGenerated(model *EnrollmentGeneratedModel) domain.EnrollmentGenerated {
	return domain.EnrollmentGenerated{
		ID:                    model.ID,
		EnrollmentApplication: model.EnrollmentApplication,
		Project: domain.Project{
			ID: model.Project.ID,
			Company: domain.Company{
				ID:   model.Project.Company.ID,
				Name: model.Project.Company.Name,
				RUC:  model.Project.Company.RUC,
			},
			Description: model.Project.Description,
			Schedule:    model.Project.Schedule,
			Starts:      model.Project.Starts,
			Ends:        model.Project.Ends,
		},
		ApprovedBy: domain.Approver{
			User: domain.User{
				ID:           model.ApprovedBy.ID,
				IDCardNumber: model.ApprovedBy.IDCardNumber,
				Name:         model.ApprovedBy.Name,
				Surname:      model.ApprovedBy.Surname,
			}},
		GeneratedAt: model.GeneratedAt,
	}
}
