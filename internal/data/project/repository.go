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

func (d *DefaultRepository) SelectAllProjects() ([]domain.Project, *domain.AppError) {
	var projectsModel []Model

	err := d.db.Select(&projectsModel, "SELECT * FROM project")
	if err != nil {
		return nil, domain.NewAppError(err, domain.RepositoryError)
	}

	var projects []domain.Project

	if len(projectsModel) == 0 {
		return projects, nil
	}

	for _, v := range projectsModel {
		projects = append(projects, fromModelToProject(&v))
	}

	return projects, nil
}
