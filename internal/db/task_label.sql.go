// Code generated by sqlc. DO NOT EDIT.
// source: task_label.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createTaskLabelForTask = `-- name: CreateTaskLabelForTask :one
INSERT INTO task_label (task_id, project_label_id, assigned_date)
  VALUES ($1, $2, $3) RETURNING task_label_id, task_id, project_label_id, assigned_date
`

type CreateTaskLabelForTaskParams struct {
	TaskID         uuid.UUID `json:"task_id"`
	ProjectLabelID uuid.UUID `json:"project_label_id"`
	AssignedDate   time.Time `json:"assigned_date"`
}

func (q *Queries) CreateTaskLabelForTask(ctx context.Context, arg CreateTaskLabelForTaskParams) (TaskLabel, error) {
	row := q.db.QueryRowContext(ctx, createTaskLabelForTask, arg.TaskID, arg.ProjectLabelID, arg.AssignedDate)
	var i TaskLabel
	err := row.Scan(
		&i.TaskLabelID,
		&i.TaskID,
		&i.ProjectLabelID,
		&i.AssignedDate,
	)
	return i, err
}

const deleteTaskLabelByID = `-- name: DeleteTaskLabelByID :exec
DELETE FROM task_label WHERE task_label_id = $1
`

func (q *Queries) DeleteTaskLabelByID(ctx context.Context, taskLabelID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTaskLabelByID, taskLabelID)
	return err
}

const deleteTaskLabelForTaskByProjectLabelID = `-- name: DeleteTaskLabelForTaskByProjectLabelID :exec
DELETE FROM task_label WHERE project_label_id = $2 AND task_id = $1
`

type DeleteTaskLabelForTaskByProjectLabelIDParams struct {
	TaskID         uuid.UUID `json:"task_id"`
	ProjectLabelID uuid.UUID `json:"project_label_id"`
}

func (q *Queries) DeleteTaskLabelForTaskByProjectLabelID(ctx context.Context, arg DeleteTaskLabelForTaskByProjectLabelIDParams) error {
	_, err := q.db.ExecContext(ctx, deleteTaskLabelForTaskByProjectLabelID, arg.TaskID, arg.ProjectLabelID)
	return err
}

const getTaskLabelByID = `-- name: GetTaskLabelByID :one
SELECT task_label_id, task_id, project_label_id, assigned_date FROM task_label WHERE task_label_id = $1
`

func (q *Queries) GetTaskLabelByID(ctx context.Context, taskLabelID uuid.UUID) (TaskLabel, error) {
	row := q.db.QueryRowContext(ctx, getTaskLabelByID, taskLabelID)
	var i TaskLabel
	err := row.Scan(
		&i.TaskLabelID,
		&i.TaskID,
		&i.ProjectLabelID,
		&i.AssignedDate,
	)
	return i, err
}

const getTaskLabelForTaskByProjectLabelID = `-- name: GetTaskLabelForTaskByProjectLabelID :one
SELECT task_label_id, task_id, project_label_id, assigned_date FROM task_label WHERE task_id = $1 AND project_label_id = $2
`

type GetTaskLabelForTaskByProjectLabelIDParams struct {
	TaskID         uuid.UUID `json:"task_id"`
	ProjectLabelID uuid.UUID `json:"project_label_id"`
}

func (q *Queries) GetTaskLabelForTaskByProjectLabelID(ctx context.Context, arg GetTaskLabelForTaskByProjectLabelIDParams) (TaskLabel, error) {
	row := q.db.QueryRowContext(ctx, getTaskLabelForTaskByProjectLabelID, arg.TaskID, arg.ProjectLabelID)
	var i TaskLabel
	err := row.Scan(
		&i.TaskLabelID,
		&i.TaskID,
		&i.ProjectLabelID,
		&i.AssignedDate,
	)
	return i, err
}

const getTaskLabelsForTaskID = `-- name: GetTaskLabelsForTaskID :many
SELECT task_label_id, task_id, project_label_id, assigned_date FROM task_label WHERE task_id = $1
`

func (q *Queries) GetTaskLabelsForTaskID(ctx context.Context, taskID uuid.UUID) ([]TaskLabel, error) {
	rows, err := q.db.QueryContext(ctx, getTaskLabelsForTaskID, taskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TaskLabel
	for rows.Next() {
		var i TaskLabel
		if err := rows.Scan(
			&i.TaskLabelID,
			&i.TaskID,
			&i.ProjectLabelID,
			&i.AssignedDate,
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
