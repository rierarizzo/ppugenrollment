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

// SelectAllProjects retrieves all projects from the database and maps them to domain.Project objects.
// If there are no projects, an empty slice will be returned.
// Returns:
//   - projects: a slice of domain.Project objects
//   - appErr: an error of type *domain.AppError if there was an error retrieving the projects from the database
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
