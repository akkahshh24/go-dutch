// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: member.sql

package db

import (
	"context"
)

const addGroupMember = `-- name: AddGroupMember :one
INSERT INTO members (
  user_id, group_id
) VALUES (
  $1, $2
)
RETURNING id, user_id, group_id, joined_at
`

type AddGroupMemberParams struct {
	UserID  int32 `json:"user_id"`
	GroupID int32 `json:"group_id"`
}

func (q *Queries) AddGroupMember(ctx context.Context, arg AddGroupMemberParams) (Member, error) {
	row := q.db.QueryRow(ctx, addGroupMember, arg.UserID, arg.GroupID)
	var i Member
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.GroupID,
		&i.JoinedAt,
	)
	return i, err
}
