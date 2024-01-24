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

func (d *DefaultProjectManager) GetProjectByID(projectID int) (*domain.Project, *domain.AppError) {
	project, appErr := d.projectRepo.SelectProjectByID(projectID)

	if appErr != nil {
		slog.Error(appErr.Error())
		return nil, domain.NewAppErrorWithType(domain.UnexpectedError)
	}

	return project, nil
}

func (d *DefaultProjectManager) AddNewProject(project *domain.Project) (*domain.Project, *domain.AppError) {
	projectWithID, appErr := d.projectRepo.InsertProject(project)

	if appErr != nil {
		slog.Error(appErr.Error())
		return nil, domain.NewAppErrorWithType(domain.UnexpectedError)
	}

	return projectWithID, nil
}

func (d *DefaultProjectManager) UpdateProject(projectID int, project *domain.Project) *domain.AppError {
	appErr := d.projectRepo.UpdateProject(projectID, project)

	if appErr != nil {
		slog.Error(appErr.Error())
		return domain.NewAppErrorWithType(domain.UnexpectedError)
	}

	return nil
}

func (d *DefaultProjectManager) DeleteProject(projectID int) *domain.AppError {
	appErr := d.projectRepo.DeleteProject(projectID)

	if appErr != nil {
		slog.Error(appErr.Error())
		return domain.NewAppErrorWithType(domain.UnexpectedError)
	}

	return nil
}

func (d *DefaultProjectManager) GetCompanies() ([]domain.Company, *domain.AppError) {
	companies, appErr := d.projectRepo.SelectCompanies()

	if appErr != nil {
		slog.Error(appErr.Error())
		return nil, domain.NewAppErrorWithType(domain.UnexpectedError)
	}

	return companies, nil
}

func (d *DefaultProjectManager) GetSchedulesByProjectID(projectID int) ([]domain.Schedule, *domain.AppError) {
	schedules, appErr := d.projectRepo.SelectSchedulesByProjectID(projectID)

	if appErr != nil {
		slog.Error(appErr.Error())
		return nil, domain.NewAppErrorWithType(domain.UnexpectedError)
	}

	return schedules, nil
}
