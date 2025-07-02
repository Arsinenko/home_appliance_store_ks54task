-- name: CreateAccount :one
INSERT INTO Accounts (login, password, created_at, is_alive)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetAccount :one
SELECT *
FROM Accounts
WHERE id = $1
LIMIT 1;

-- name: ListAccounts :many
SELECT *
FROM Accounts
WHERE is_alive = true
ORDER BY login;

-- name: UpdateAccount :one
UPDATE Accounts
SET login    = $2,
    password = $3,
    is_alive = $4
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
UPDATE Accounts
SET is_alive = false
WHERE id = $1;