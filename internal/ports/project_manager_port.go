package ports

import "ppugenrollment/internal/domain"

type ProjectManager interface {
	GetAllProjects() ([]domain.Project, *domain.AppError)
}

type ProjectRepository interface {
	SelectAllProjects() ([]domain.Project, *domain.AppError)
}
