// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: groups.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createGroup = `-- name: CreateGroup :one
INSERT INTO chat_group (id, name, about, ppic)
VALUES ($1, $2, $3, $4)
RETURNING id, name, about, ppic, created_at
`

type CreateGroupParams struct {
	ID    uuid.UUID
	Name  string
	About string
	Ppic  sql.NullString
}

func (q *Queries) CreateGroup(ctx context.Context, arg CreateGroupParams) (ChatGroup, error) {
	row := q.db.QueryRowContext(ctx, createGroup,
		arg.ID,
		arg.Name,
		arg.About,
		arg.Ppic,
	)
	var i ChatGroup
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.About,
		&i.Ppic,
		&i.CreatedAt,
	)
	return i, err
}

const deleteGroup = `-- name: DeleteGroup :exec
DELETE FROM chat_group
WHERE id = $1
`

func (q *Queries) DeleteGroup(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteGroup, id)
	return err
}

const getGroupByID = `-- name: GetGroupByID :one
SELECT id, name, about, ppic, created_at FROM chat_group
WHERE id = $1
`

func (q *Queries) GetGroupByID(ctx context.Context, id uuid.UUID) (ChatGroup, error) {
	row := q.db.QueryRowContext(ctx, getGroupByID, id)
	var i ChatGroup
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.About,
		&i.Ppic,
		&i.CreatedAt,
	)
	return i, err
}

const listGroups = `-- name: ListGroups :many
SELECT id, name, about, ppic, created_at FROM chat_group
ORDER BY created_at DESC
LIMIT $1 OFFSET $2
`

type ListGroupsParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListGroups(ctx context.Context, arg ListGroupsParams) ([]ChatGroup, error) {
	rows, err := q.db.QueryContext(ctx, listGroups, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ChatGroup
	for rows.Next() {
		var i ChatGroup
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.About,
			&i.Ppic,
			&i.CreatedAt,
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

const updateGroupInfo = `-- name: UpdateGroupInfo :one
UPDATE chat_group
SET name = $2,
    about = $3,
    ppic = $4
WHERE id = $1
RETURNING id, name, about, ppic, created_at
`

type UpdateGroupInfoParams struct {
	ID    uuid.UUID
	Name  string
	About string
	Ppic  sql.NullString
}

func (q *Queries) UpdateGroupInfo(ctx context.Context, arg UpdateGroupInfoParams) (ChatGroup, error) {
	row := q.db.QueryRowContext(ctx, updateGroupInfo,
		arg.ID,
		arg.Name,
		arg.About,
		arg.Ppic,
	)
	var i ChatGroup
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.About,
		&i.Ppic,
		&i.CreatedAt,
	)
	return i, err
}
