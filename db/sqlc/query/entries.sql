-- name: CreateEntries :one
INSERT INTO entries (
    account_id, amount
) VALUES (
    $1, $2
)
RETURNING *;

-- name: UpdateEntries :one
UPDATE entries SET
amount = $2
WHERE id= $1
RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: GetEntriesFromAccount :many
SELECT e.id, e.amount, e.created_at FROM entries AS e
INNER JOIN accounts AS a
ON a.id = e.account_id
WHERE a.id = $1
LIMIT $2
OFFSET $3;

-- name ListAllEntries
SELECT * FROM entries
WHERE entries.account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = $1;
