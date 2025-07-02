-- name: CreateSupplier :one
INSERT INTO Suppliers (id, account_id, created_at, is_alive)
VALUES ($1, $2, $3, $4)
    RETURNING *;

-- name: GetSupplier :one
-- Получаем поставщика вместе с его логином
SELECT s.*, a.login
FROM Suppliers s
         JOIN Accounts a ON s.account_id = a.id
WHERE s.id = $1
    LIMIT 1;

-- name: ListSuppliers :many
-- Получаем список поставщиков вместе с их логинами
SELECT s.*, a.login
FROM Suppliers s
         JOIN Accounts a ON s.account_id = a.id
WHERE s.is_alive = true
ORDER BY s.id;

-- name: UpdateSupplier :one
UPDATE Suppliers
SET is_alive = $2
WHERE id = $1
    RETURNING *;

-- name: DeleteSupplier :exec
UPDATE Suppliers
SET is_alive = false
WHERE id = $1;