package usecases

import (
	"log/slog"
	"ppugenrollment/internal/domain"
	"ppugenrollment/internal/ports"
)

type DefaultManager struct {
	projectRepo ports.ProjectRepository
}

func NewProjectManager(projectRepo ports.ProjectRepository) *DefaultManager {
	return &DefaultManager{projectRepo}
}

func (d DefaultManager) GetAllProjects() ([]domain.Project, *domain.AppError) {
	projects, appErr := d.projectRepo.SelectAllProjects()
	if appErr != nil {
		slog.Error(appErr.Error())
		return nil, appErr
	}

	return projects, nil
}
