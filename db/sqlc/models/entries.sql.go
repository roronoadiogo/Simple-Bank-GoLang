// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: entries.sql

package models

import (
	"context"
	"database/sql"
)

const createEntries = `-- name: CreateEntries :one
INSERT INTO entries (
    account_id, amount
) VALUES (
    $1, $2
)
RETURNING id, account_id, amount, created_at
`

type CreateEntriesParams struct {
	AccountID sql.NullInt64 `db:"account_id" json:"account_id"`
	Amount    int64         `db:"amount" json:"amount"`
}

func (q *Queries) CreateEntries(ctx context.Context, arg CreateEntriesParams) (Entry, error) {
	row := q.db.QueryRowContext(ctx, createEntries, arg.AccountID, arg.Amount)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteEntry = `-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = $1
`

func (q *Queries) DeleteEntry(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteEntry, id)
	return err
}

const getEntriesFromAccount = `-- name: GetEntriesFromAccount :many
SELECT e.id, e.amount, e.created_at FROM entries AS e
INNER JOIN accounts AS a
ON a.id = e.account_id
WHERE a.id = $1
LIMIT $2
OFFSET $3
`

type GetEntriesFromAccountParams struct {
	ID     int64 `db:"id" json:"id"`
	Limit  int32 `db:"limit" json:"limit"`
	Offset int32 `db:"offset" json:"offset"`
}

type GetEntriesFromAccountRow struct {
	ID        int64        `db:"id" json:"id"`
	Amount    int64        `db:"amount" json:"amount"`
	CreatedAt sql.NullTime `db:"created_at" json:"created_at"`
}

func (q *Queries) GetEntriesFromAccount(ctx context.Context, arg GetEntriesFromAccountParams) ([]GetEntriesFromAccountRow, error) {
	rows, err := q.db.QueryContext(ctx, getEntriesFromAccount, arg.ID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetEntriesFromAccountRow
	for rows.Next() {
		var i GetEntriesFromAccountRow
		if err := rows.Scan(&i.ID, &i.Amount, &i.CreatedAt); err != nil {
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

const getEntry = `-- name: GetEntry :one
SELECT id, account_id, amount, created_at FROM entries
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetEntry(ctx context.Context, id int64) (Entry, error) {
	row := q.db.QueryRowContext(ctx, getEntry, id)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const updateEntries = `-- name: UpdateEntries :one
UPDATE entries SET
amount = $2
WHERE id= $1
RETURNING id, account_id, amount, created_at
`

type UpdateEntriesParams struct {
	ID     int64 `db:"id" json:"id"`
	Amount int64 `db:"amount" json:"amount"`
}

func (q *Queries) UpdateEntries(ctx context.Context, arg UpdateEntriesParams) (Entry, error) {
	row := q.db.QueryRowContext(ctx, updateEntries, arg.ID, arg.Amount)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}
