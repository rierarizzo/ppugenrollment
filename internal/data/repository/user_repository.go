package repository

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"ppugenrollment/internal/data/sqlcgen"
	"ppugenrollment/pkg/domain"
)

type DefaultUserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *DefaultUserRepository {
	return &DefaultUserRepository{db}
}

func (r *DefaultUserRepository) InsertUser(user *domain.User) *domain.AppError {
	queries := sqlcgen.New(r.db)

	_, err := queries.CreateUser(context.Background(), sqlcgen.CreateUserParams{
		IDCardNumber: user.IDCardNumber,
		Name:         user.Name,
		Surname:      user.Surname,
		Email:        user.Email,
		Password:     user.Password,
		Role:         user.Role,
		DateOfBirth:  sql.NullTime{Time: user.DateOfBirth},
		IsAGraduate:  sql.NullBool{Bool: user.IsAGraduate},
		Level:        sql.NullInt32{Int32: int32(user.Level)},
	})

	if err != nil {
		slog.Error(err.Error())
		return domain.NewAppError(err, domain.RepositoryError)
	}

	return nil
}

func (r *DefaultUserRepository) SelectUserByEmail(email string) (*domain.User, *domain.AppError) {
	queries := sqlcgen.New(r.db)

	userModel, err := queries.GetUserByEmail(context.Background(), email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.NewAppError(err, domain.NotFoundError)
		}

		return nil, domain.NewAppError(err, domain.RepositoryError)
	}

	user := domain.User{
		ID:           int(userModel.ID),
		IDCardNumber: userModel.IDCardNumber,
		Name:         userModel.Name,
		Surname:      userModel.Surname,
		Email:        userModel.Email,
		Password:     userModel.Password,
		Role:         userModel.Role,
		DateOfBirth:  userModel.DateOfBirth.Time,
		IsAGraduate:  userModel.IsAGraduate.Bool,
		Level:        int(userModel.Level.Int32),
	}

	return &user, nil
}
