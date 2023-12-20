package mappers

import (
	"ppugenrollment/internal/data/models"
	"ppugenrollment/internal/domain"
)

func FromModelToProject(model *models.ProjectModel) domain.Project {
	return domain.Project{
		ID:          model.ID,
		Company:     domain.Company{ID: model.Company},
		Name:        model.Name,
		Description: model.Description,
		Starts:      model.Starts,
		Ends:        model.Ends,
	}
}
