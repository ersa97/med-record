-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = ? LIMIT 1;

-- name: GetRoleById :one
SELECT * FROM roles WHERE id = ? LIMIT 1;

-- name: GetPermissionkeysByRoleId :many
SELECT permission_key FROM rolepermissions WHERE roleid = ? ;