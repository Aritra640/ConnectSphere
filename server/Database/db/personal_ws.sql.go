// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: personal_ws.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createPersonalWS = `-- name: CreatePersonalWS :one
INSERT INTO personal_ws (id, usera, userb)
VALUES ($1, $2, $3)
RETURNING id, usera, userb
`

type CreatePersonalWSParams struct {
	ID    uuid.UUID
	Usera sql.NullInt32
	Userb sql.NullInt32
}

func (q *Queries) CreatePersonalWS(ctx context.Context, arg CreatePersonalWSParams) (PersonalW, error) {
	row := q.db.QueryRowContext(ctx, createPersonalWS, arg.ID, arg.Usera, arg.Userb)
	var i PersonalW
	err := row.Scan(&i.ID, &i.Usera, &i.Userb)
	return i, err
}

const deletePersonalWSbyID = `-- name: DeletePersonalWSbyID :exec
DELETE FROM personal_ws WHERE id = $1
`

func (q *Queries) DeletePersonalWSbyID(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deletePersonalWSbyID, id)
	return err
}

const getPersonalWSbyUsers = `-- name: GetPersonalWSbyUsers :one
SELECT id, usera, userb FROM personal_ws 
WHERE (usera = $1 AND userb = $2) OR (usera = $2 AND userb = $1)
`

type GetPersonalWSbyUsersParams struct {
	Usera sql.NullInt32
	Userb sql.NullInt32
}

func (q *Queries) GetPersonalWSbyUsers(ctx context.Context, arg GetPersonalWSbyUsersParams) (PersonalW, error) {
	row := q.db.QueryRowContext(ctx, getPersonalWSbyUsers, arg.Usera, arg.Userb)
	var i PersonalW
	err := row.Scan(&i.ID, &i.Usera, &i.Userb)
	return i, err
}
