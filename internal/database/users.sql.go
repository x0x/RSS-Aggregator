// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, created_at ,updated_at, name, api_key, role)
VALUES ($1, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex'), $5)
RETURNING id, created_at, updated_at, name, api_key, role
`

type CreateUserParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Role      string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Role,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.ApiKey,
		&i.Role,
	)
	return i, err
}

const getUserByApiKey = `-- name: GetUserByApiKey :one
SELECT id, created_at, updated_at, name, api_key, role FROM users WHERE api_key = $1
`

func (q *Queries) GetUserByApiKey(ctx context.Context, apiKey string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByApiKey, apiKey)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.ApiKey,
		&i.Role,
	)
	return i, err
}
