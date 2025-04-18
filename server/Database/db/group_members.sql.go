// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: group_members.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const addUserToGroup = `-- name: AddUserToGroup :exec
INSERT INTO chat_group_member (group_id, user_id , is_admin)
VALUES ($1 , $2 , $3)
ON CONFLICT (group_id , user_id) DO NOTHING
`

type AddUserToGroupParams struct {
	GroupID uuid.UUID
	UserID  int32
	IsAdmin sql.NullBool
}

func (q *Queries) AddUserToGroup(ctx context.Context, arg AddUserToGroupParams) error {
	_, err := q.db.ExecContext(ctx, addUserToGroup, arg.GroupID, arg.UserID, arg.IsAdmin)
	return err
}

const getAppAdmins = `-- name: GetAppAdmins :many
SELECT u.id, u.username, u.email, u.password_hashed, u.created_at
FROM chat_group_member cm
JOIN users u ON cm.user_id = u.id
WHERE cm.group_id = $1 AND cm.is_admin = TRUE
`

func (q *Queries) GetAppAdmins(ctx context.Context, groupID uuid.UUID) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getAppAdmins, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.PasswordHashed,
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

const getGroupMembers = `-- name: GetGroupMembers :many
SELECT u.id, u.username, u.email, u.password_hashed, u.created_at
FROM chat_group_member cm
JOIN users u ON cm.user_id = u.id
WHERE cm.group_id = $1
`

func (q *Queries) GetGroupMembers(ctx context.Context, groupID uuid.UUID) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getGroupMembers, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.PasswordHashed,
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

const isUserGroupAdmin = `-- name: IsUserGroupAdmin :one
SELECT EXISTS (
  SELECT 1 FROM chat_group_member
  WHERE user_id = $1 AND group_id = $2 AND is_admin = TRUE
) AS is_admin
`

type IsUserGroupAdminParams struct {
	UserID  int32
	GroupID uuid.UUID
}

func (q *Queries) IsUserGroupAdmin(ctx context.Context, arg IsUserGroupAdminParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, isUserGroupAdmin, arg.UserID, arg.GroupID)
	var is_admin bool
	err := row.Scan(&is_admin)
	return is_admin, err
}

const isUserInGroup = `-- name: IsUserInGroup :one
SELECT EXISTS (
  SELECT 1 FROM chat_group_member
  WHERE user_id = $1 AND group_id = $2
) AS is_member
`

type IsUserInGroupParams struct {
	UserID  int32
	GroupID uuid.UUID
}

func (q *Queries) IsUserInGroup(ctx context.Context, arg IsUserInGroupParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, isUserInGroup, arg.UserID, arg.GroupID)
	var is_member bool
	err := row.Scan(&is_member)
	return is_member, err
}

const removeUserFromGroup = `-- name: RemoveUserFromGroup :exec
DELETE FROM chat_group_member WHERE user_id = $1 AND group_id = $2
`

type RemoveUserFromGroupParams struct {
	UserID  int32
	GroupID uuid.UUID
}

func (q *Queries) RemoveUserFromGroup(ctx context.Context, arg RemoveUserFromGroupParams) error {
	_, err := q.db.ExecContext(ctx, removeUserFromGroup, arg.UserID, arg.GroupID)
	return err
}

const setUserToAdminStatus = `-- name: SetUserToAdminStatus :exec
UPDATE chat_group_member
SET is_admin = $3
WHERE user_id = $1 AND group_id = $2
`

type SetUserToAdminStatusParams struct {
	UserID  int32
	GroupID uuid.UUID
	IsAdmin sql.NullBool
}

func (q *Queries) SetUserToAdminStatus(ctx context.Context, arg SetUserToAdminStatusParams) error {
	_, err := q.db.ExecContext(ctx, setUserToAdminStatus, arg.UserID, arg.GroupID, arg.IsAdmin)
	return err
}
