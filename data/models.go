// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package data

import (
	"database/sql"
	"time"
)

type Role struct {
	ID        int64
	Rolename  string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}

type Rolepermission struct {
	ID            int64
	Roleid        int64
	PermissionKey sql.NullString
	CreatedAt     time.Time
	UpdatedAt     sql.NullTime
	DeletedAt     sql.NullTime
}

type User struct {
	ID           int64
	Roleid       int64
	Username     string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    sql.NullTime
	DeletedAt    sql.NullTime
}