-- name: CreateEnrollmentApplication :execresult
INSERT INTO enrollment_application (student, project, schedule)
VALUES (?, ?, ?);