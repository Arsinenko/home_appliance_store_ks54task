-- name: CreateGoodsSupplier :one
INSERT INTO Goods_Suppliers (id, supplier_id, good_id, created_at, is_alive)
VALUES ($1, $2, $3, $4, $5)
    RETURNING *;

-- name: CreateGoodsSuppliers :one
INSERT INTO Goods_Suppliers (id, supplier_id, good_id, created_at, is_alive)
VALUES ($1, $2, $3, $4, $5)
    RETURNING *;

-- name: GetGoodsSupplier :one
SELECT *
FROM Goods_Suppliers
WHERE id = $1
    LIMIT 1;

-- name: ListGoodsBySupplier :many
-- Получаем все товары для конкретного поставщика
SELECT g.*
FROM Goods g
         JOIN Goods_Suppliers gs ON g.id = gs.good_id
WHERE gs.supplier_id = $1
  AND g.is_alive = true
  AND gs.is_alive = true;

-- name: ListSuppliersByGood :many
-- Получаем всех поставщиков для конкретного товара, включая их логины
SELECT s.*, a.login
FROM Suppliers s
         JOIN Goods_Suppliers gs ON s.id = gs.supplier_id
         JOIN Accounts a ON s.account_id = a.id
WHERE gs.good_id = $1
  AND s.is_alive = true
  AND gs.is_alive = true;

-- name: DeleteGoodsSupplier :exec
-- "Мягкое" удаление связи товара с поставщиком
UPDATE Goods_Suppliers
SET is_alive = false
WHERE id = $1;

-- name: UpdateGoodsSupplier :one
UPDATE Goods_Suppliers
SET is_alive = $2
WHERE id = $1
    RETURNING *;