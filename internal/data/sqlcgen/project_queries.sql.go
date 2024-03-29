// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: project_queries.sql

package sqlcgen

import (
	"context"
	"database/sql"
	"time"
)

const createProject = `-- name: CreateProject :execresult
INSERT INTO project (company, name, description, starts, ends) VALUES (?, ?, ?, ?, ?)
`

type CreateProjectParams struct {
	Company     int32
	Name        string
	Description string
	Starts      time.Time
	Ends        time.Time
}

func (q *Queries) CreateProject(ctx context.Context, arg CreateProjectParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createProject,
		arg.Company,
		arg.Name,
		arg.Description,
		arg.Starts,
		arg.Ends,
	)
}

const createScheduleForProject = `-- name: CreateScheduleForProject :execresult
INSERT INTO project_schedule (project, schedule) VALUES (?, ?)
`

type CreateScheduleForProjectParams struct {
	Project  int32
	Schedule string
}

func (q *Queries) CreateScheduleForProject(ctx context.Context, arg CreateScheduleForProjectParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createScheduleForProject, arg.Project, arg.Schedule)
}

const deleteProject = `-- name: DeleteProject :exec
DELETE FROM project WHERE id = ?
`

func (q *Queries) DeleteProject(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteProject, id)
	return err
}

const deleteProjectSchedules = `-- name: DeleteProjectSchedules :exec
DELETE FROM project_schedule WHERE project = ?
`

func (q *Queries) DeleteProjectSchedules(ctx context.Context, project int32) error {
	_, err := q.db.ExecContext(ctx, deleteProjectSchedules, project)
	return err
}

const getCompanies = `-- name: GetCompanies :many
SELECT id, name, ruc, image_url FROM company ORDER BY name
`

func (q *Queries) GetCompanies(ctx context.Context) ([]Company, error) {
	rows, err := q.db.QueryContext(ctx, getCompanies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Company
	for rows.Next() {
		var i Company
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Ruc,
			&i.ImageUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCompany = `-- name: GetCompany :one
SELECT id, name, ruc, image_url FROM company WHERE id = ?
`

func (q *Queries) GetCompany(ctx context.Context, id int32) (Company, error) {
	row := q.db.QueryRowContext(ctx, getCompany, id)
	var i Company
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Ruc,
		&i.ImageUrl,
	)
	return i, err
}

const getProjectById = `-- name: GetProjectById :one
SELECT id, company, name, description, starts, ends FROM project WHERE id = ?
`

func (q *Queries) GetProjectById(ctx context.Context, id int32) (Project, error) {
	row := q.db.QueryRowContext(ctx, getProjectById, id)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Company,
		&i.Name,
		&i.Description,
		&i.Starts,
		&i.Ends,
	)
	return i, err
}

const getProjectSchedulesByProjectID = `-- name: GetProjectSchedulesByProjectID :many
SELECT id, project, schedule FROM project_schedule WHERE project = ?
`

func (q *Queries) GetProjectSchedulesByProjectID(ctx context.Context, project int32) ([]ProjectSchedule, error) {
	rows, err := q.db.QueryContext(ctx, getProjectSchedulesByProjectID, project)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProjectSchedule
	for rows.Next() {
		var i ProjectSchedule
		if err := rows.Scan(&i.ID, &i.Project, &i.Schedule); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProjects = `-- name: GetProjects :many
SELECT id, company, name, description, starts, ends FROM project ORDER BY starts
`

func (q *Queries) GetProjects(ctx context.Context) ([]Project, error) {
	rows, err := q.db.QueryContext(ctx, getProjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Project
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ID,
			&i.Company,
			&i.Name,
			&i.Description,
			&i.Starts,
			&i.Ends,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProject = `-- name: UpdateProject :exec
UPDATE project SET company = ?, name = ?, description = ?, starts = ?, ends = ? WHERE id = ?
`

type UpdateProjectParams struct {
	Company     int32
	Name        string
	Description string
	Starts      time.Time
	Ends        time.Time
	ID          int32
}

func (q *Queries) UpdateProject(ctx context.Context, arg UpdateProjectParams) error {
	_, err := q.db.ExecContext(ctx, updateProject,
		arg.Company,
		arg.Name,
		arg.Description,
		arg.Starts,
		arg.Ends,
		arg.ID,
	)
	return err
}
