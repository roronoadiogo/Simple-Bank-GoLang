-- name: CreateTransfers :one
INSERT INTO transfers (
    from_account_id, to_account_id, amount
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: GetTranfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: GetTransfersFromAccount :many
SELECT t.id, t.to_account_id, t.amount, t.created_at FROM transfers AS t
INNER JOIN accounts AS a
ON a.id = t.from_account_id 
WHERE a.id = $1
LIMIT $2
OFFSET $3;

-- name ListAllTransfers
SELECT * FROM transfers
WHERE transfers.from_account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: DeleteTransfers :exec
DELETE FROM transfers
WHERE id = $1;
