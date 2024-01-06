package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"ppugenrollment/internal/data/sqlcgen"
	"ppugenrollment/pkg/domain"
)

type DefaultProjectRepository struct {
	db *sqlx.DB
}

func NewProjectRepository(db *sqlx.DB) *DefaultProjectRepository {
	return &DefaultProjectRepository{db}
}

func (d *DefaultProjectRepository) SelectAllProjects() ([]domain.Project, *domain.AppError) {
	queries := sqlcgen.New(d.db)

	projectsModel, err := queries.GetProjects(context.Background())

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

func (d *DefaultProjectRepository) InsertProject(project *domain.Project) (*domain.Project, *domain.AppError) {
	queries := sqlcgen.New(d.db)

	result, err := queries.CreateProject(context.Background(), sqlcgen.CreateProjectParams{
		Company:     int32(project.Company.ID),
		Name:        project.Name,
		Description: project.Description,
		Starts:      project.Starts,
		Ends:        project.Ends,
	})

	if err != nil {
		return nil, domain.NewAppError(err, domain.RepositoryError)
	}

	lastInsertedID, _ := result.LastInsertId()
	project.ID = int(lastInsertedID)

	return project, nil
}

func fromModelToProject(model *sqlcgen.Project) domain.Project {
	return domain.Project{
		ID:          int(model.ID),
		Company:     domain.Company{ID: int(model.Company)},
		Name:        model.Name,
		Description: model.Description,
		Starts:      model.Starts,
		Ends:        model.Ends,
	}
}
