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

func FromProjectToModel(project *domain.Project) models.ProjectModel {
	return models.ProjectModel{
		Company:     project.Company.ID,
		Name:        project.Name,
		Description: project.Description,
		Starts:      project.Starts,
		Ends:        project.Ends,
	}
}
