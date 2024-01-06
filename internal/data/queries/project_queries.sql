-- name: GetProjects :many
SELECT *
FROM project
ORDER BY starts;

-- name: CreateProject :execresult
INSERT INTO project (company, name, description, starts, ends)
VALUES (?, ?, ?, ?, ?);