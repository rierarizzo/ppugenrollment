package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"ppugenrollment/internal/data/sqlcgen"
	"ppugenrollment/pkg/domain"
)

type DefaultProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) *DefaultProjectRepository {
	return &DefaultProjectRepository{db}
}

func (d *DefaultProjectRepository) SelectAllProjects() ([]domain.Project, *domain.AppError) {
	queries := sqlcgen.New(d.db)

	projectsModel, err := queries.GetProjects(context.Background())

	if err != nil {
		return nil, domain.NewAppError(err.Error(), domain.RepositoryError)
	}

	var projects []domain.Project

	if len(projectsModel) == 0 {
		return projects, nil
	}

	for _, model := range projectsModel {
		schedules, appErr := d.SelectProjectSchedulesByProjectID(int(model.ID))

		if appErr != nil {
			return nil, appErr
		}

		company, appErr := d.SelectCompanyByID(int(model.Company))

		if appErr != nil {
			return nil, appErr
		}

		projects = append(projects, domain.Project{
			ID:          int(model.ID),
			Company:     *company,
			Name:        model.Name,
			Description: model.Description,
			Schedules:   schedules,
			Starts:      model.Starts,
			Ends:        model.Ends,
		})
	}

	return projects, nil
}

func (d *DefaultProjectRepository) SelectProjectByID(projectID int) (*domain.Project, *domain.AppError) {
	queries := sqlcgen.New(d.db)

	projectModel, err := queries.GetProjectById(context.Background(), int32(projectID))

	if err != nil {
		return nil, domain.NewAppError(err.Error(), domain.RepositoryError)
	}

	project := domain.Project{
		ID:          int(projectModel.ID),
		Company:     domain.Company{ID: int(projectModel.Company)},
		Name:        projectModel.Name,
		Description: projectModel.Description,
		Starts:      projectModel.Starts,
		Ends:        projectModel.Ends,
	}

	schedules, appErr := d.SelectProjectSchedulesByProjectID(projectID)

	if appErr != nil {
		return nil, appErr
	}

	project.Schedules = schedules

	company, appErr := d.SelectCompanyByID(int(projectModel.Company))

	if appErr != nil {
		return nil, appErr
	}

	project.Company = *company

	return &project, nil
}

func (d *DefaultProjectRepository) SelectProjectSchedulesByProjectID(projectID int) ([]string, *domain.AppError) {
	queries := sqlcgen.New(d.db)

	schedulesModel, err := queries.GetProjectSchedulesByProjectID(context.Background(), int32(projectID))

	if err != nil {
		return nil, domain.NewAppError(err.Error(), domain.RepositoryError)
	}

	var schedules []string

	for _, model := range schedulesModel {
		schedules = append(schedules, model.Schedule)
	}

	return schedules, nil
}

func (d *DefaultProjectRepository) SelectCompanyByID(companyID int) (*domain.Company, *domain.AppError) {
	queries := sqlcgen.New(d.db)

	companyModel, err := queries.GetCompany(context.Background(), int32(companyID))

	if err != nil {
		return nil, domain.NewAppError(err.Error(), domain.RepositoryError)
	}

	company := domain.Company{
		ID:       int(companyModel.ID),
		Name:     companyModel.Name,
		RUC:      companyModel.Ruc,
		ImageURL: companyModel.ImageUrl.String,
	}

	return &company, nil
}

func (d *DefaultProjectRepository) InsertProject(project *domain.Project) (*domain.Project, *domain.AppError) {
	ctx := context.Background()

	tx, err := d.db.Begin()

	if err != nil {
		return nil, domain.NewAppError(err.Error(), domain.RepositoryError)
	}

	defer func(tx *sql.Tx) {
		err = tx.Rollback()

		if !errors.Is(err, sql.ErrTxDone) {
			slog.Error("an error occurred while rolling back the transaction: " + err.Error())
		}
	}(tx)

	queries := sqlcgen.New(d.db)
	qtx := queries.WithTx(tx)

	result, err := qtx.CreateProject(ctx, sqlcgen.CreateProjectParams{
		Company:     int32(project.Company.ID),
		Name:        project.Name,
		Description: project.Description,
		Starts:      project.Starts,
		Ends:        project.Ends,
	})

	if err != nil {
		return nil, domain.NewAppError(err.Error(), domain.RepositoryError)
	}

	lastInsertedID, _ := result.LastInsertId()
	project.ID = int(lastInsertedID)

	for _, schedule := range project.Schedules {
		_, err := qtx.CreateScheduleForProject(ctx, sqlcgen.CreateScheduleForProjectParams{
			Project:  int32(lastInsertedID),
			Schedule: schedule,
		})

		if err != nil {
			return nil, domain.NewAppError(err.Error(), domain.RepositoryError)
		}

		slog.Debug(fmt.Sprintf("Schedule '%s' created for project with ID %v", schedule, lastInsertedID))
	}

	company, appErr := d.SelectCompanyByID(project.Company.ID)

	if appErr != nil {
		return nil, appErr
	}

	project.Company = *company

	err = tx.Commit()

	if err != nil {
		slog.Error("an error occurred when committing the transaction: " + err.Error())
		return nil, domain.NewAppError(err.Error(), domain.RepositoryError)
	}

	return project, nil
}

func (d *DefaultProjectRepository) UpdateProject(projectID int, project *domain.Project) *domain.AppError {
	ctx := context.Background()

	tx, err := d.db.Begin()

	if err != nil {
		return domain.NewAppError(err.Error(), domain.RepositoryError)
	}

	defer func(tx *sql.Tx) {
		err = tx.Rollback()

		if !errors.Is(err, sql.ErrTxDone) {
			slog.Error("an error occurred while rolling back the transaction: " + err.Error())
		}
	}(tx)

	queries := sqlcgen.New(d.db)
	qtx := queries.WithTx(tx)

	err = qtx.DeleteProjectSchedules(context.Background(), int32(projectID))

	if err != nil {
		return domain.NewAppError(err.Error(), domain.RepositoryError)
	}

	for _, schedule := range project.Schedules {
		_, err := qtx.CreateScheduleForProject(ctx, sqlcgen.CreateScheduleForProjectParams{
			Project:  int32(projectID),
			Schedule: schedule,
		})

		if err != nil {
			return domain.NewAppError(err.Error(), domain.RepositoryError)
		}

		slog.Debug(fmt.Sprintf("Schedule '%s' created for project with ID %v", schedule, projectID))
	}

	err = qtx.UpdateProject(context.Background(), sqlcgen.UpdateProjectParams{
		Company:     int32(project.Company.ID),
		Name:        project.Name,
		Description: project.Description,
		Starts:      project.Starts,
		Ends:        project.Ends,
		ID:          int32(projectID),
	})

	if err != nil {
		return domain.NewAppError(err.Error(), domain.RepositoryError)
	}

	err = tx.Commit()

	if err != nil {
		slog.Error("an error occurred when committing the transaction: " + err.Error())
		return domain.NewAppError(err.Error(), domain.RepositoryError)
	}

	return nil
}

func (d *DefaultProjectRepository) DeleteProject(projectID int) *domain.AppError {
	queries := sqlcgen.New(d.db)

	err := queries.DeleteProjectSchedules(context.Background(), int32(projectID))

	if err != nil {
		return domain.NewAppError(err.Error(), domain.RepositoryError)
	}

	err = queries.DeleteProject(context.Background(), int32(projectID))

	if err != nil {
		return domain.NewAppError(err.Error(), domain.RepositoryError)
	}

	return nil
}

func (d *DefaultProjectRepository) SelectCompanies() ([]domain.Company, *domain.AppError) {
	queries := sqlcgen.New(d.db)

	companiesModel, err := queries.GetCompanies(context.Background())

	if err != nil {
		return nil, domain.NewAppError(err.Error(), domain.RepositoryError)
	}

	var companies []domain.Company

	for _, model := range companiesModel {
		companies = append(companies, domain.Company{
			ID:   int(model.ID),
			Name: model.Name,
			RUC:  model.Ruc,
		})
	}

	return companies, nil
}
