-- name: CreateGood :one
INSERT INTO Goods (article, price, name, quantity, is_alive)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: CreateManyGoods :copyfrom
INSERT INTO Goods (article, price, name, quantity, is_alive)
VALUES ($1, $2, $3, $4, $5);

-- name: GetGood :one
SELECT *
FROM Goods
WHERE id = $1
LIMIT 1;

-- name: ListGoods :many
SELECT *
FROM Goods
WHERE is_alive = true
ORDER BY name;

-- name: UpdateGood :one
UPDATE Goods
SET article  = $2,
    price    = $3,
    name     = $4,
    quantity = $5,
    is_alive = $6
WHERE id = $1
RETURNING *;

-- name: DeleteGood :exec
UPDATE Goods
SET is_alive = false
WHERE id = $1;