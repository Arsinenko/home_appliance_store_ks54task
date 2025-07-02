-- name: CreateStore :one
INSERT INTO Stores (address, created_at, is_alive)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetStore :one
select *
from stores
where id = $1
limit 1;

-- name: GetStores :many
select *
from stores
where is_alive = true
order by id;

-- name: UpdateStore :one
update stores
set address = $2,
    updated_at = now(),
    is_alive = $3
where id = $1
returning *;

-- name: DeleteStore :exec
update stores
set is_alive = false
where id = $1;