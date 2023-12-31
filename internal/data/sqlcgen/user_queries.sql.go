// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user_queries.sql

package sqlcgen

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :execresult
INSERT INTO user (id_card_number, name, surname, email, password, role, date_of_birth, is_a_graduate, level)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateUserParams struct {
	IDCardNumber string
	Name         string
	Surname      string
	Email        string
	Password     string
	Role         string
	DateOfBirth  sql.NullTime
	IsAGraduate  sql.NullBool
	Level        sql.NullInt32
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser,
		arg.IDCardNumber,
		arg.Name,
		arg.Surname,
		arg.Email,
		arg.Password,
		arg.Role,
		arg.DateOfBirth,
		arg.IsAGraduate,
		arg.Level,
	)
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, id_card_number, name, surname, email, password, role, date_of_birth, is_a_graduate, level FROM user WHERE email = ?
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.IDCardNumber,
		&i.Name,
		&i.Surname,
		&i.Email,
		&i.Password,
		&i.Role,
		&i.DateOfBirth,
		&i.IsAGraduate,
		&i.Level,
	)
	return i, err
}
