package student

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"ppugenrollment/internal/domain"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) InsertStudent(student *domain.Student) *domain.AppError {
	model := fromStudentToModel(student)

	insertInUserSchema := `
		INSERT INTO sys_user (id_card_number, name, surname, email, password, role) 
		VALUES (?,?,?,?,?,?)
	`
	result, err := r.db.Exec(insertInUserSchema, model.IDCardNumber, model.Name, model.Surname, model.Email,
		model.Password, model.Role)
	if err != nil {
		return domain.NewAppError(err, domain.RepositoryError)
	}

	lastInsertID, _ := result.LastInsertId()

	insertInStudentSchema := `
		INSERT INTO student (sys_user, date_of_birth, is_a_graduate, level)
		VALUES (?,?,?,?)
	`
	_, err = r.db.Exec(insertInStudentSchema, lastInsertID, model.DateOfBirth, model.IsAGraduate, model.Level)
	if err != nil {
		return domain.NewAppError(err, domain.RepositoryError)
	}

	return nil
}

func (r Repository) GetStudentByEmail(email string) (*domain.Student, *domain.AppError) {
	model := Model{}

	selectInUserSchema := `
		SELECT * FROM sys_user u INNER JOIN student s ON u.id = s.sys_user WHERE email=?
	`
	err := r.db.Select(&model, selectInUserSchema, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.NewAppError(err, domain.NotFoundError)
		}

		return nil, domain.NewAppError(err, domain.RepositoryError)
	}

	student := fromModelToStudent(&model)
	return &student, nil
}
