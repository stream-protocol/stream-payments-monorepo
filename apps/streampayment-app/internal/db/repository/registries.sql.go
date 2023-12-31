// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: registries.sql

package repository

import (
	"context"
	"time"
)

const createRegistryItem = `-- name: CreateRegistryItem :one
insert into registries (
created_at, updated_at,
merchant_id, key, value,
description
) values ($1, $2, $3, $4, $5, $6)
returning id, created_at, updated_at, merchant_id, key, value, description
`

type CreateRegistryItemParams struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	MerchantID  int64
	Key         string
	Value       string
	Description string
}

func (q *Queries) CreateRegistryItem(ctx context.Context, arg CreateRegistryItemParams) (Registry, error) {
	row := q.db.QueryRow(ctx, createRegistryItem,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.MerchantID,
		arg.Key,
		arg.Value,
		arg.Description,
	)
	var i Registry
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.MerchantID,
		&i.Key,
		&i.Value,
		&i.Description,
	)
	return i, err
}

const getRegistryItemByKey = `-- name: GetRegistryItemByKey :one
select id, created_at, updated_at, merchant_id, key, value, description from registries where merchant_id = $1 and key = $2
`

type GetRegistryItemByKeyParams struct {
	MerchantID int64
	Key        string
}

func (q *Queries) GetRegistryItemByKey(ctx context.Context, arg GetRegistryItemByKeyParams) (Registry, error) {
	row := q.db.QueryRow(ctx, getRegistryItemByKey, arg.MerchantID, arg.Key)
	var i Registry
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.MerchantID,
		&i.Key,
		&i.Value,
		&i.Description,
	)
	return i, err
}

const updateRegistryItem = `-- name: UpdateRegistryItem :one
update registries
set updated_at = $3, value = $4, description = $5
where key = $1 and merchant_id = $2
returning id, created_at, updated_at, merchant_id, key, value, description
`

type UpdateRegistryItemParams struct {
	Key         string
	MerchantID  int64
	UpdatedAt   time.Time
	Value       string
	Description string
}

func (q *Queries) UpdateRegistryItem(ctx context.Context, arg UpdateRegistryItemParams) (Registry, error) {
	row := q.db.QueryRow(ctx, updateRegistryItem,
		arg.Key,
		arg.MerchantID,
		arg.UpdatedAt,
		arg.Value,
		arg.Description,
	)
	var i Registry
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.MerchantID,
		&i.Key,
		&i.Value,
		&i.Description,
	)
	return i, err
}
