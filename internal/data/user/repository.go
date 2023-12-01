package user

import (
	"database/sql"
	"errors"
	"ppugenrollment/internal/domain"

	"github.com/jmoiron/sqlx"
)

type DefaultRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *DefaultRepository {
	return &DefaultRepository{db: db}
}

func (r *DefaultRepository) InsertStudent(student *domain.Student) *domain.AppError {
	model := fromStudentToStudentModel(student)

	lastInsertID, _ := r.insertUserAndGetLastID(
		model.IDCardNumber, model.Name, model.Surname, model.Email, model.Password, model.Role)

	insertInStudentSchema := `
		INSERT INTO student (sys_user, date_of_birth, is_a_graduate, level)
		VALUES (?,?,?,?)
	`
	_, err := r.db.Exec(insertInStudentSchema, lastInsertID, model.DateOfBirth, model.IsAGraduate, model.Level)
	if err != nil {
		return domain.NewAppError(err, domain.RepositoryError)
	}

	return nil
}

func (r *DefaultRepository) InsertAdmin(admin *domain.Admin) *domain.AppError {
	model := fromAdminToAdminModel(admin)

	lastInsertID, _ := r.insertUserAndGetLastID(
		model.IDCardNumber, model.Name, model.Surname, model.Email, model.Password, model.Role)

	insertInStudentSchema := `
		INSERT INTO admin (sys_user) VALUES (?)
	`
	_, err := r.db.Exec(insertInStudentSchema, lastInsertID)
	if err != nil {
		return domain.NewAppError(err, domain.RepositoryError)
	}

	return nil
}

func (r *DefaultRepository) InsertApprover(approver *domain.Approver) *domain.AppError {
	model := fromApproverToApproverModel(approver)

	lastInsertID, _ := r.insertUserAndGetLastID(
		model.IDCardNumber, model.Name, model.Surname, model.Email, model.Password, model.Role)

	insertInStudentSchema := `
		INSERT INTO approver (sys_user) VALUES (?)
	`
	_, err := r.db.Exec(insertInStudentSchema, lastInsertID)
	if err != nil {
		return domain.NewAppError(err, domain.RepositoryError)
	}

	return nil
}

func (r *DefaultRepository) insertUserAndGetLastID(idCardNumber, name, surname, email, password, role string) (
	int,
	*domain.AppError) {
	insertInUserSchema := `
		INSERT INTO sys_user (id_card_number, name, surname, email, password, role) 
		VALUES (?,?,?,?,?,?)
	`
	result, err := r.db.Exec(insertInUserSchema, idCardNumber, name, surname, email, password, role)
	if err != nil {
		return 0, domain.NewAppError(err, domain.RepositoryError)
	}

	lastInsertID, _ := result.LastInsertId()

	return int(lastInsertID), nil
}

func (r *DefaultRepository) SelectUserByEmail(email string) (*domain.User, *domain.AppError) {
	model := CommonFieldsModel{}

	selectInUserSchema := `
		SELECT * FROM sys_user WHERE email=?
	`
	err := r.db.Select(&model, selectInUserSchema, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.NewAppError(err, domain.NotFoundError)
		}

		return nil, domain.NewAppError(err, domain.RepositoryError)
	}

	user := fromCommonFieldsModelToUser(&model)
	return &user, nil
}
