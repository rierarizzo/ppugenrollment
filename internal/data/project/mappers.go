package project

import (
	"ppugenrollment/internal/domain"
)

func fromModelToProject(model *Model) domain.Project {
	return domain.Project{
		ID:          model.ID,
		Company:     domain.Company{ID: model.Company},
		Description: model.Description,
		Schedule:    model.Schedule,
		Starts:      model.Starts,
		Ends:        model.Ends,
	}
}

func fromProjectToModel(project *domain.Project) Model {
	return Model{
		ID:          project.ID,
		Company:     project.Company.ID,
		Description: project.Description,
		Schedule:    project.Schedule,
		Starts:      project.Starts,
		Ends:        project.Ends,
	}
}
