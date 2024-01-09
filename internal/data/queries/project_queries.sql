-- name: GetProjects :many
SELECT * FROM project ORDER BY starts;

-- name: GetProjectById :one
SELECT * FROM project WHERE id = ?;

-- name: GetProjectSchedulesByProjectID :many
SELECT * FROM project_schedule WHERE project = ?;

-- name: CreateProject :execresult
INSERT INTO project (company, name, description, starts, ends) VALUES (?, ?, ?, ?, ?);

-- name: CreateScheduleForProject :execresult
INSERT INTO project_schedule (project, schedule) VALUES (?, ?);

-- name: UpdateProject :exec
UPDATE project SET company = ?, name = ?, description = ?, starts = ?, ends = ? WHERE id = ?;

-- name: DeleteProject :exec
DELETE FROM project WHERE id = ?;

-- name: GetCompany :one
SELECT * FROM company WHERE id = ?;