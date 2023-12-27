package usecases

import (
	"log/slog"
	"ppugenrollment/internal/ports"
	"ppugenrollment/pkg/domain"
)

type DefaultProjectManager struct {
	projectRepo ports.ProjectRepository
}

func NewProjectManager(projectRepo ports.ProjectRepository) *DefaultProjectManager {
	return &DefaultProjectManager{projectRepo}
}

func (d *DefaultProjectManager) GetAllProjects() ([]domain.Project, *domain.AppError) {
	projects, appErr := d.projectRepo.SelectAllProjects()

	if appErr != nil {
		slog.Error(appErr.Error())
		return nil, domain.NewAppErrorWithType(domain.UnexpectedError)
	}

	return projects, nil
}

func (d *DefaultProjectManager) AddNewProject(project *domain.Project) (*domain.Project, *domain.AppError) {
	projectWithID, appErr := d.projectRepo.InsertProject(project)

	if appErr != nil {
		slog.Error(appErr.Error())
		return nil, domain.NewAppErrorWithType(domain.UnexpectedError)
	}

	return projectWithID, nil
}
