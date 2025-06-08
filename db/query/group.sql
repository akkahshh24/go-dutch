-- name: CreateGroup :one
INSERT INTO groups (
  name, description, created_by
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetGroup :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;