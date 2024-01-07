package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"ppugenrollment/internal/data/sqlcgen"
	"ppugenrollment/pkg/domain"
)

type DefaultApprovalRepository struct {
	db *sqlx.DB
}

func NewApprovalRepository(db *sqlx.DB) *DefaultApprovalRepository {
	return &DefaultApprovalRepository{db}
}

func (d *DefaultApprovalRepository) InsertEnrollmentApproval(applicationID, approvedBy int) (*domain.EnrollmentGenerated,
	*domain.AppError) {
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

	err = qtx.UpdateEnrollmentApplication(ctx, int32(applicationID))

	if err != nil {
		return nil, domain.NewAppError("application already approved", domain.RepositoryError)
	}

	result, err := qtx.CreateEnrollmentGenerated(ctx, sqlcgen.CreateEnrollmentGeneratedParams{
		EnrollmentApplication: int32(applicationID),
		ApprovedBy:            int32(approvedBy),
	})

	if err != nil {
		return nil, domain.NewAppError(err.Error(), domain.RepositoryError)
	}

	generatedID, _ := result.LastInsertId()

	generatedModel, err := qtx.GetEnrollmentGenerated(ctx, int32(generatedID))

	if err != nil {
		return nil, domain.NewAppError(err.Error(), domain.RepositoryError)
	}

	err = tx.Commit()

	if err != nil {
		slog.Error("an error occurred when committing the transaction: " + err.Error())
		return nil, domain.NewAppError(err.Error(), domain.RepositoryError)
	}

	generated := fromModelToEnrollmentGenerated(&generatedModel)

	return &generated, nil
}

func fromModelToEnrollmentGenerated(model *sqlcgen.GetEnrollmentGeneratedRow) domain.EnrollmentGenerated {
	return domain.EnrollmentGenerated{
		ID:                    int(model.ID),
		EnrollmentApplication: int(model.ApplicationID),
		Project: domain.Project{
			ID: int(model.ProjectID),
			Company: domain.Company{
				ID:   int(model.CompanyID),
				Name: model.CompanyName,
				RUC:  model.CompanyRuc,
			},
			Description: model.ProjectDescription,
			Starts:      model.ProjectStarts,
			Ends:        model.ProjectEnds,
		},
		Schedule: model.ProjectSchedule,
		ApprovedBy: domain.User{
			ID:           int(model.ApproverID),
			IDCardNumber: model.ApproverCardNumber,
			Name:         model.ApproverName,
			Surname:      model.ApproverSurname,
		},
		GeneratedAt: model.GeneratedAt,
	}
}
