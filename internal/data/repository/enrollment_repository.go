package repository

import (
	"context"
	"database/sql"
	"ppugenrollment/internal/data/sqlcgen"
	"ppugenrollment/pkg/domain"
)

type DefaultEnrollmentRepository struct {
	db *sql.DB
}

func NewEnrollmentRepository(db *sql.DB) *DefaultEnrollmentRepository {
	return &DefaultEnrollmentRepository{db}
}

func (d *DefaultEnrollmentRepository) InsertEnrollment(application *domain.EnrollmentApplication) (int,
	*domain.AppError) {
	queries := sqlcgen.New(d.db)

	scheduleResult, err := queries.CreateScheduleForProject(context.Background(), sqlcgen.CreateScheduleForProjectParams{
		Project:  int32(application.Project.ID),
		Schedule: "M",
	})

	if err != nil {
		return 0, domain.NewAppError(err.Error(), domain.RepositoryError)
	}

	scheduleID, _ := scheduleResult.LastInsertId()

	result, err := queries.CreateEnrollmentApplication(context.Background(), sqlcgen.CreateEnrollmentApplicationParams{
		Student:  int32(application.Student.ID),
		Project:  int32(application.Project.ID),
		Schedule: int32(scheduleID),
	})

	if err != nil {
		return 0, domain.NewAppError(err.Error(), domain.RepositoryError)
	}

	lastInsertedID, _ := result.LastInsertId()

	return int(lastInsertedID), nil
}

func (d *DefaultEnrollmentRepository) SelectEnrollmentApplications() ([]domain.EnrollmentApplication, *domain.AppError) {
	queries := sqlcgen.New(d.db)

	applicationRows, err := queries.GetEnrollmentApplications(context.Background())

	if err != nil {
		return nil, domain.NewAppError(err.Error(), domain.RepositoryError)
	}

	applications := make([]domain.EnrollmentApplication, 0)
	for _, app := range applicationRows {
		applications = append(applications, domain.EnrollmentApplication{
			ID: int(app.ApplicationID),
			Student: domain.User{
				ID:      int(app.StudentID),
				Name:    app.StudentName,
				Surname: app.StudentSurname,
			},
			Project: domain.Project{
				ID:   int(app.ProjectID),
				Name: app.ProjectName,
				Company: domain.Company{
					ID:   int(app.CompanyID),
					Name: app.CompanyName,
					RUC:  app.CompanyRuc,
				},
				Description: app.ProjectDescription,
				Starts:      app.ProjectStarts,
				Ends:        app.ProjectEnds,
			},
			ScheduleCode: app.ProjectSchedule,
			Status:       app.ApplicationStatus,
		})
	}

	return applications, nil
}
