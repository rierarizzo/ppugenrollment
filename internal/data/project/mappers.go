package project

import (
	"ppugenrollment/internal/domain"
)

func fromModelToProject(model *Model) domain.Project {
	return domain.Project{
		ID:          model.ID,
		Company:     domain.Company{ID: model.Company},
		Name:        model.Name,
		Description: model.Description,
		Starts:      model.Starts,
		Ends:        model.Ends,
	}
}

func fromProjectToModel(project *domain.Project) Model {
	return Model{
		ID:          project.ID,
		Company:     project.Company.ID,
		Name:        project.Name,
		Description: project.Description,
		Starts:      project.Starts,
		Ends:        project.Ends,
	}
}
