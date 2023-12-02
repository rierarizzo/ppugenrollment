package project_manager

import "ppugenrollment/internal/domain"

type Manager interface {
	GetAllProjects() ([]domain.Project, *domain.AppError)
}

type Repository interface {
	SelectAllProjects() ([]domain.Project, *domain.AppError)
}
