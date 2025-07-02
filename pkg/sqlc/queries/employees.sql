-- name: CreateEmployee :one
INSERT INTO Employees (account_id, role_id, created_at, is_alive)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetEmployee :one
SELECT
    e.id,
    e.account_id,
    e.role_id,
    e.created_at,
    e.is_alive,

    a.id as id_account,
    a.login as account_login,
    a.created_at as account_created_at,
    a.is_alive as account_is_alive,

    r.id as id_role,
    r.name as role_name,
    r.created_at as role_created_at

FROM Employees e
         JOIN Accounts a ON e.account_id = a.id
         JOIN Roles r ON e.role_id = r.id
WHERE e.id = $1
LIMIT 1;

-- name: ListEmployees :many
SELECT
    e.id,
    e.account_id,
    e.role_id,
    e.created_at,
    e.is_alive,

    a.id as id_account,
    a.login as account_login,
    a.created_at as account_created_at,
    a.is_alive as account_is_alive,

    r.id as id_role,
    r.name as role_name,
    r.created_at as role_created_at

FROM Employees e
         JOIN Accounts a ON e.account_id = a.id
         JOIN Roles r ON e.role_id = r.id
WHERE e.is_alive = true
ORDER BY e.id;

-- name: UpdateEmployee :one
UPDATE Employees
SET role_id  = $2,
    is_alive = $3
WHERE id = $1
RETURNING *;

-- name: DeleteEmployee :exec
UPDATE Employees
SET is_alive = false
WHERE id = $1;