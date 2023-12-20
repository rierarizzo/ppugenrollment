package repository

import (
	"github.com/jmoiron/sqlx"
	"ppugenrollment/internal/data/mappers"
	"ppugenrollment/internal/data/models"
	"ppugenrollment/internal/domain"
)

type DefaultProjectRepository struct {
	db *sqlx.DB
}

func NewProjectRepository(db *sqlx.DB) *DefaultProjectRepository {
	return &DefaultProjectRepository{db}
}

func (d *DefaultProjectRepository) SelectAllProjects() ([]domain.Project, *domain.AppError) {
	var projectsModel []models.ProjectModel

	err := d.db.Select(&projectsModel, "SELECT * FROM project")
	if err != nil {
		return nil, domain.NewAppError(err, domain.RepositoryError)
	}

	var projects []domain.Project

	if len(projectsModel) == 0 {
		return projects, nil
	}

	for _, v := range projectsModel {
		projects = append(projects, mappers.FromModelToProject(&v))
	}

	return projects, nil
}
