package ports

import "ppugenrollment/pkg/domain"

type ProjectManager interface {
	GetAllProjects() ([]domain.Project, *domain.AppError)
	GetProjectByID(projectID int) (*domain.Project, *domain.AppError)
	AddNewProject(project *domain.Project) (*domain.Project, *domain.AppError)
	UpdateProject(projectID int, project *domain.Project) *domain.AppError
	DeleteProject(projectID int) *domain.AppError
	GetCompanies() ([]domain.Company, *domain.AppError)
	GetSchedulesByProjectID(projectID int) ([]domain.Schedule, *domain.AppError)
}

type ProjectRepository interface {
	SelectAllProjects() ([]domain.Project, *domain.AppError)
	SelectProjectByID(projectID int) (*domain.Project, *domain.AppError)
	InsertProject(project *domain.Project) (*domain.Project, *domain.AppError)
	UpdateProject(projectID int, project *domain.Project) *domain.AppError
	DeleteProject(projectID int) *domain.AppError
	SelectCompanies() ([]domain.Company, *domain.AppError)
	SelectSchedulesByProjectID(projectID int) ([]domain.Schedule, *domain.AppError)
}
