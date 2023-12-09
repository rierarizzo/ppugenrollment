package user

import (
	"database/sql"
	"errors"
	"log/slog"
	"ppugenrollment/internal/domain"

	"github.com/jmoiron/sqlx"
)

type DefaultRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *DefaultRepository {
	return &DefaultRepository{db}
}

func (r *DefaultRepository) InsertUser(user *domain.User) *domain.AppError {
	model := fromUserToModel(user)

	insertInUserSchema := `
		INSERT INTO user (id_card_number, name, surname, email, password, role, date_of_birth, is_a_graduate, level) 
		VALUES (?,?,?,?,?,?,?,?,?)
	`
	_, err := r.db.Exec(
		insertInUserSchema,
		model.IDCardNumber,
		model.Name,
		model.Surname,
		model.Email,
		model.Password,
		model.Role,
		model.DateOfBirth,
		model.IsAGraduate,
		model.Level)
	if err != nil {
		slog.Error(err.Error())
		return domain.NewAppError(err, domain.RepositoryError)
	}

	return nil
}

func (r *DefaultRepository) SelectUserByEmail(email string) (*domain.User, *domain.AppError) {
	var model Model

	selectInUserSchema := `
		SELECT * FROM user WHERE email=?
	`
	err := r.db.Get(&model, selectInUserSchema, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.NewAppError(err, domain.NotFoundError)
		}

		return nil, domain.NewAppError(err, domain.RepositoryError)
	}

	user := fromModelToUser(&model)
	return &user, nil
}
