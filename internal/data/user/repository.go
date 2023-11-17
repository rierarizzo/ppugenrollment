package user

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"ppugenrollment/internal/domain"
)

type DefaultRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *DefaultRepository {
	return &DefaultRepository{db: db}
}

func (r *DefaultRepository) InsertStudent(student *domain.Student) *domain.AppError {
	model := fromStudentToStudentModel(student)

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

func (r *DefaultRepository) InsertAdmin(admin *domain.Admin) *domain.AppError {
	model := fromAdminToAdminModel(admin)

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
		INSERT INTO admin (sys_user) VALUES (?)
	`
	_, err = r.db.Exec(insertInStudentSchema, lastInsertID)
	if err != nil {
		return domain.NewAppError(err, domain.RepositoryError)
	}

	return nil
}

func (r *DefaultRepository) InsertApprover(approver *domain.Approver) *domain.AppError {
	model := fromApproverToApproverModel(approver)

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
		INSERT INTO approver (sys_user) VALUES (?)
	`
	_, err = r.db.Exec(insertInStudentSchema, lastInsertID)
	if err != nil {
		return domain.NewAppError(err, domain.RepositoryError)
	}

	return nil
}

func (r *DefaultRepository) SelectUserByEmail(email string) (*domain.CommonUserFields, *domain.AppError) {
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
