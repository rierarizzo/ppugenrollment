package project

import "ppugenrollment/internal/domain"

func fromModelToProject(model *Model) domain.Project {
	return domain.Project{
		ID:          model.ID,
		Company:     model.Company,
		Description: model.Description,
		Schedule:    model.Schedule,
	}
}

func fromProjectToModel(project *domain.Project) Model {
	return Model{
		ID:          project.ID,
		Company:     project.Company,
		Description: project.Description,
		Schedule:    project.Schedule,
	}
}
