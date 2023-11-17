package project

import (
	"github.com/jmoiron/sqlx"
	"ppugenrollment/internal/domain"
)

type DefaultRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *DefaultRepository {
	return &DefaultRepository{db}
}

func (d DefaultRepository) SelectAllProjects() ([]domain.Project, *domain.AppError) {
	var projectsModel []Model

	selectAllInProjectSchema := `
		SELECT * FROM project
	`
	err := d.db.Select(projectsModel, selectAllInProjectSchema)
	if err != nil {
		return nil, domain.NewAppError(err, domain.RepositoryError)
	}

	if len(projectsModel) <= 0 {
		return nil, domain.NewAppErrorWithType(domain.NotFoundError)
	}

	var projects []domain.Project
	for _, v := range projectsModel {
		projects = append(projects, fromModelToProject(&v))
	}

	return projects, nil
}
