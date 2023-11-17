package project_management

import (
	"log/slog"
	"ppugenrollment/internal/domain"
)

type DefaultManager struct {
	projectRepo Repository
}

func New(projectRepo Repository) *DefaultManager {
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
