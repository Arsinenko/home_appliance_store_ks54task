-- name: CreateRole :one
insert into roles (name, created_at) VALUES ($1, now()) returning *;

-- name: GetRole :one
select * from roles
where id = $1
limit 1;

-- name: GetRoles :many
select *
from roles
order by id;

-- name: UpdateRole :one
update roles
set name = $2
where id = $1
returning *;
