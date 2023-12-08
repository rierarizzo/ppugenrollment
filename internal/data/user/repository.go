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
	return &DefaultRepository{db}
}

func insertUserAndGetLastID(tx *sqlx.Tx, user *Model) (
	int, *domain.AppError) {
	insertInUserSchema := `
		INSERT INTO sys_user (id_card_number, name, surname, email, password, role) 
		VALUES (?,?,?,?,?,?)
	`
	result, err := tx.Exec(
		insertInUserSchema, user.IDCardNumber, user.Name, user.Surname, user.Email, user.Password, user.Role)
	if err != nil {
		_ = tx.Rollback()
		return 0, domain.NewAppError(err, domain.RepositoryError)
	}

	lastInsertID, _ := result.LastInsertId()

	return int(lastInsertID), nil
}

func (r *DefaultRepository) InsertStudent(student *domain.Student) *domain.AppError {
	model := fromStudentToStudentModel(student)

	user := Model{
		IDCardNumber: model.IDCardNumber,
		Name:         model.Name,
		Surname:      model.Surname,
		Email:        model.Email,
		Password:     model.Password,
		Role:         model.Role,
	}

	tx, _ := r.db.Beginx()
	lastInsertID, appErr := insertUserAndGetLastID(tx, &user)
	if appErr != nil {
		return appErr
	}

	insertInStudentSchema := `
		INSERT INTO student (sys_user, date_of_birth, is_a_graduate, level)
		VALUES (?,?,?,?)
	`
	_, err := tx.Exec(insertInStudentSchema, lastInsertID, model.DateOfBirth, model.IsAGraduate, model.Level)
	if err != nil {
		_ = tx.Rollback()
		return domain.NewAppError(err, domain.RepositoryError)
	}

	_ = tx.Commit()

	return nil
}

func (r *DefaultRepository) InsertAdmin(admin *domain.Admin) *domain.AppError {
	model := fromAdminToAdminModel(admin)

	user := Model{
		IDCardNumber: model.IDCardNumber,
		Name:         model.Name,
		Surname:      model.Surname,
		Email:        model.Email,
		Password:     model.Password,
		Role:         model.Role,
	}

	tx, _ := r.db.Beginx()
	lastInsertID, appErr := insertUserAndGetLastID(tx, &user)
	if appErr != nil {
		return appErr
	}

	insertInAdminSchema := `
		INSERT INTO admin (sys_user) VALUES (?)
	`
	_, err := r.db.Exec(insertInAdminSchema, lastInsertID)
	if err != nil {
		_ = tx.Rollback()
		return domain.NewAppError(err, domain.RepositoryError)
	}

	_ = tx.Commit()

	return nil
}

func (r *DefaultRepository) InsertApprover(approver *domain.Approver) *domain.AppError {
	model := fromApproverToApproverModel(approver)

	user := Model{
		IDCardNumber: model.IDCardNumber,
		Name:         model.Name,
		Surname:      model.Surname,
		Email:        model.Email,
		Password:     model.Password,
		Role:         model.Role,
	}

	tx, _ := r.db.Beginx()
	lastInsertID, appErr := insertUserAndGetLastID(tx, &user)
	if appErr != nil {
		return appErr
	}

	insertInApproverSchema := `
		INSERT INTO approver (sys_user) VALUES (?)
	`
	_, err := r.db.Exec(insertInApproverSchema, lastInsertID)
	if err != nil {
		_ = tx.Rollback()
		return domain.NewAppError(err, domain.RepositoryError)
	}

	_ = tx.Commit()

	return nil
}

func (r *DefaultRepository) SelectUserByEmail(email string) (*domain.User, *domain.AppError) {
	var model Model

	selectInUserSchema := `
		SELECT * FROM sys_user WHERE email=?
	`
	err := r.db.Get(&model, selectInUserSchema, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.NewAppError(err, domain.NotFoundError)
		}

		return nil, domain.NewAppError(err, domain.RepositoryError)
	}

	user := fromCommonFieldsModelToUser(&model)
	return &user, nil
}
