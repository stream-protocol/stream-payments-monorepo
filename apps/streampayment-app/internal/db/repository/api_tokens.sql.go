// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: api_tokens.sql

package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

const createAPIToken = `-- name: CreateAPIToken :one
INSERT INTO api_tokens (
    entity_type,
    entity_id,
    created_at,
    token,
    uuid,
    name,
    settings
) VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, entity_type, entity_id, created_at, token, name, settings, uuid
`

type CreateAPITokenParams struct {
	EntityType string
	EntityID   int64
	CreatedAt  time.Time
	Token      string
	Uuid       uuid.UUID
	Name       sql.NullString
	Settings   pgtype.JSONB
}

func (q *Queries) CreateAPIToken(ctx context.Context, arg CreateAPITokenParams) (ApiToken, error) {
	row := q.db.QueryRow(ctx, createAPIToken,
		arg.EntityType,
		arg.EntityID,
		arg.CreatedAt,
		arg.Token,
		arg.Uuid,
		arg.Name,
		arg.Settings,
	)
	var i ApiToken
	err := row.Scan(
		&i.ID,
		&i.EntityType,
		&i.EntityID,
		&i.CreatedAt,
		&i.Token,
		&i.Name,
		&i.Settings,
		&i.Uuid,
	)
	return i, err
}

const deleteAPITokenByID = `-- name: DeleteAPITokenByID :exec
DELETE from api_tokens
WHERE id = $1
`

func (q *Queries) DeleteAPITokenByID(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAPITokenByID, id)
	return err
}

const deleteAPITokenByToken = `-- name: DeleteAPITokenByToken :exec
DELETE from api_tokens
WHERE token = $1
`

func (q *Queries) DeleteAPITokenByToken(ctx context.Context, token string) error {
	_, err := q.db.Exec(ctx, deleteAPITokenByToken, token)
	return err
}

const getAPIToken = `-- name: GetAPIToken :one
SELECT id, entity_type, entity_id, created_at, token, name, settings, uuid FROM api_tokens
WHERE entity_type = $1 and token = $2 LIMIT 1
`

type GetAPITokenParams struct {
	EntityType string
	Token      string
}

func (q *Queries) GetAPIToken(ctx context.Context, arg GetAPITokenParams) (ApiToken, error) {
	row := q.db.QueryRow(ctx, getAPIToken, arg.EntityType, arg.Token)
	var i ApiToken
	err := row.Scan(
		&i.ID,
		&i.EntityType,
		&i.EntityID,
		&i.CreatedAt,
		&i.Token,
		&i.Name,
		&i.Settings,
		&i.Uuid,
	)
	return i, err
}

const getAPITokenByUUID = `-- name: GetAPITokenByUUID :one
SELECT id, entity_type, entity_id, created_at, token, name, settings, uuid FROM api_tokens
WHERE uuid = $1 LIMIT 1
`

func (q *Queries) GetAPITokenByUUID(ctx context.Context, uuid uuid.UUID) (ApiToken, error) {
	row := q.db.QueryRow(ctx, getAPITokenByUUID, uuid)
	var i ApiToken
	err := row.Scan(
		&i.ID,
		&i.EntityType,
		&i.EntityID,
		&i.CreatedAt,
		&i.Token,
		&i.Name,
		&i.Settings,
		&i.Uuid,
	)
	return i, err
}

const listAPITokensByEntity = `-- name: ListAPITokensByEntity :many
SELECT id, entity_type, entity_id, created_at, token, name, settings, uuid FROM api_tokens
WHERE entity_id = $1 and entity_type = $2
ORDER BY id DESC
`

type ListAPITokensByEntityParams struct {
	EntityID   int64
	EntityType string
}

func (q *Queries) ListAPITokensByEntity(ctx context.Context, arg ListAPITokensByEntityParams) ([]ApiToken, error) {
	rows, err := q.db.Query(ctx, listAPITokensByEntity, arg.EntityID, arg.EntityType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ApiToken
	for rows.Next() {
		var i ApiToken
		if err := rows.Scan(
			&i.ID,
			&i.EntityType,
			&i.EntityID,
			&i.CreatedAt,
			&i.Token,
			&i.Name,
			&i.Settings,
			&i.Uuid,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}