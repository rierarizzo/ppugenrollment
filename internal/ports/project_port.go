package ports

import "ppugenrollment/pkg/domain"

type ProjectManager interface {
	GetAllProjects() ([]domain.Project, *domain.AppError)
	AddNewProject(project *domain.Project) (*domain.Project, *domain.AppError)
}

type ProjectRepository interface {
	SelectAllProjects() ([]domain.Project, *domain.AppError)
	InsertProject(project *domain.Project) (*domain.Project, *domain.AppError)
}
