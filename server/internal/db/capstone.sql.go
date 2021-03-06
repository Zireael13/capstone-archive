// Code generated by sqlc. DO NOT EDIT.
// source: capstone.sql

package db

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
)

const createCapstone = `-- name: CreateCapstone :one
INSERT INTO capstones (id, created_at, updated_at, title, description, author, semester, slug)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING
    id, created_at, updated_at, deleted_at, title, description, author, semester, slug
`

type CreateCapstoneParams struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Description string
	Author      string
	Semester    string
	Slug        string
}

func (q *Queries) CreateCapstone(ctx context.Context, arg CreateCapstoneParams) (Capstone, error) {
	row := q.db.QueryRowContext(ctx, createCapstone,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Title,
		arg.Description,
		arg.Author,
		arg.Semester,
		arg.Slug,
	)
	var i Capstone
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Title,
		&i.Description,
		&i.Author,
		&i.Semester,
		&i.Slug,
	)
	return i, err
}

const getCapstoneById = `-- name: GetCapstoneById :one
SELECT
    id, created_at, updated_at, deleted_at, title, description, author, semester, slug
FROM
    capstones
WHERE
    id = $1
LIMIT 1
`

func (q *Queries) GetCapstoneById(ctx context.Context, id uuid.UUID) (Capstone, error) {
	row := q.db.QueryRowContext(ctx, getCapstoneById, id)
	var i Capstone
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Title,
		&i.Description,
		&i.Author,
		&i.Semester,
		&i.Slug,
	)
	return i, err
}

const getCapstoneBySlug = `-- name: GetCapstoneBySlug :one
SELECT
    id, created_at, updated_at, deleted_at, title, description, author, semester, slug
FROM
    capstones
WHERE
    slug = $1
LIMIT 1
`

func (q *Queries) GetCapstoneBySlug(ctx context.Context, slug string) (Capstone, error) {
	row := q.db.QueryRowContext(ctx, getCapstoneBySlug, slug)
	var i Capstone
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Title,
		&i.Description,
		&i.Author,
		&i.Semester,
		&i.Slug,
	)
	return i, err
}

const getCapstoneByTitle = `-- name: GetCapstoneByTitle :one
SELECT
    id, created_at, updated_at, deleted_at, title, description, author, semester, slug
FROM
    capstones
WHERE
    title = $1
LIMIT 1
`

func (q *Queries) GetCapstoneByTitle(ctx context.Context, title string) (Capstone, error) {
	row := q.db.QueryRowContext(ctx, getCapstoneByTitle, title)
	var i Capstone
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Title,
		&i.Description,
		&i.Author,
		&i.Semester,
		&i.Slug,
	)
	return i, err
}

const getCapstones = `-- name: GetCapstones :many
SELECT
    id, created_at, updated_at, deleted_at, title, description, author, semester, slug
FROM
    capstones
ORDER BY
    created_at DESC
LIMIT $1
`

func (q *Queries) GetCapstones(ctx context.Context, limit int32) ([]Capstone, error) {
	rows, err := q.db.QueryContext(ctx, getCapstones, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Capstone
	for rows.Next() {
		var i Capstone
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Title,
			&i.Description,
			&i.Author,
			&i.Semester,
			&i.Slug,
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

const getCapstonesWithCursor = `-- name: GetCapstonesWithCursor :many
SELECT
    id, created_at, updated_at, deleted_at, title, description, author, semester, slug
FROM
    capstones
WHERE
    created_at < $1
ORDER BY
    created_at DESC
LIMIT $2
`

type GetCapstonesWithCursorParams struct {
	CreatedAt time.Time
	Limit     int32
}

func (q *Queries) GetCapstonesWithCursor(ctx context.Context, arg GetCapstonesWithCursorParams) ([]Capstone, error) {
	rows, err := q.db.QueryContext(ctx, getCapstonesWithCursor, arg.CreatedAt, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Capstone
	for rows.Next() {
		var i Capstone
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Title,
			&i.Description,
			&i.Author,
			&i.Semester,
			&i.Slug,
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

const searchCapstones = `-- name: SearchCapstones :many
SELECT
    id, created_at, updated_at, deleted_at, title, description, author, semester, slug
FROM
    capstones c
WHERE
    to_tsvector(c.Title) || to_tsvector(c.Description) || to_tsvector(c.Author) || to_tsvector(c.Semester) @@ to_tsquery('english', $1)
LIMIT $2 OFFSET $3
`

type SearchCapstonesParams struct {
	ToTsquery string
	Limit     int32
	Offset    int32
}

func (q *Queries) SearchCapstones(ctx context.Context, arg SearchCapstonesParams) ([]Capstone, error) {
	rows, err := q.db.QueryContext(ctx, searchCapstones, arg.ToTsquery, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Capstone
	for rows.Next() {
		var i Capstone
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Title,
			&i.Description,
			&i.Author,
			&i.Semester,
			&i.Slug,
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
