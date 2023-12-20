package repository

import (
	"database/sql"
	"errors"
	"log/slog"
	"ppugenrollment/internal/data/mappers"
	"ppugenrollment/internal/data/models"
	"ppugenrollment/internal/domain"

	"github.com/jmoiron/sqlx"
)

type DefaultUserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *DefaultUserRepository {
	return &DefaultUserRepository{db}
}

// InsertUser inserts a new user into the repository.
// It maps the user data from the domain.User object to a models.UserModel object.
// The user data is then inserted into the database using the INSERT statement.
// If there is an error during the insertion process, it returns a domain.RepositoryError AppError.
// If the user is successfully inserted, it returns nil.
func (r *DefaultUserRepository) InsertUser(user *domain.User) *domain.AppError {
	model := mappers.FromUserToModel(user)

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

// SelectUserByEmail selects a user from the repository by email.
// It retrieves the user's data from the database and maps it to a domain.User object.
// If the user is not found, it returns a domain.NotFoundError AppError.
// If there is an error during the retrieval process, it returns a domain.RepositoryError AppError.
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
