package repository

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"ppugenrollment/internal/data/mappers"
	"ppugenrollment/internal/data/models"
	"ppugenrollment/pkg/domain"
)

type DefaultUserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *DefaultUserRepository {
	return &DefaultUserRepository{db}
}

func (r *DefaultUserRepository) InsertUser(user *domain.User) *domain.AppError {
	model := mappers.FromUserToModel(user)

	insertInUserSchema := `
		INSERT INTO user (id_card_number, name, surname, email, password, role, date_of_birth, is_a_graduate, level) 
		VALUES (?,?,?,?,?,?,?,?,?)
	`
	_, err := r.db.Exec(insertInUserSchema, model.IDCardNumber, model.Name, model.Surname, model.Email, model.Password,
		model.Role, model.DateOfBirth, model.IsAGraduate, model.Level)

	if err != nil {
		slog.Error(err.Error())
		return domain.NewAppError(err, domain.RepositoryError)
	}

	return nil
}

func (r *DefaultUserRepository) SelectUserByEmail(email string) (*domain.User, *domain.AppError) {
	var model models.UserModel

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

	user := mappers.FromModelToUser(&model)
	return &user, nil
}
