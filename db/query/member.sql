-- name: AddGroupMember :one
INSERT INTO members (
  user_id, group_id
) VALUES (
  $1, $2
)
RETURNING *;