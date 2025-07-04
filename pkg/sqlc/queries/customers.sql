-- name: CreateCustomer :one
INSERT INTO Customers (account_id, balance, created_at, is_alive)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetCustomer :one
-- Получаем покупателя вместе с его логином
SELECT
    c.*,
    a.login as account_login,
    a.created_at as account_created_at,
    a.is_alive as account_is_alive

FROM Customers c
         JOIN Accounts a ON c.account_id = a.id
WHERE c.id = $1
LIMIT 1;

-- name: ListCustomers :many
-- Получаем список покупателей вместе с их логинами
SELECT c.*,
       a.login as account_login,
       a.created_at as account_created_at,
       a.is_alive as account_is_alive
FROM Customers c
         JOIN Accounts a ON c.account_id = a.id
WHERE c.is_alive = true
ORDER BY c.id;

-- name: UpdateCustomer :one
UPDATE Customers
SET balance  = $2,
    is_alive = $3
WHERE id = $1
RETURNING *;

-- name: DeleteCustomer :exec
UPDATE Customers
SET is_alive = false
WHERE id = $1;