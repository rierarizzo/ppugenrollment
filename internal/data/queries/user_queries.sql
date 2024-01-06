-- name: GetUserByEmail :one
SELECT *
FROM user
WHERE email = ?;

-- name: CreateUser :execresult
INSERT INTO user (id_card_number, name, surname, email, password, role, date_of_birth, is_a_graduate, level)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);