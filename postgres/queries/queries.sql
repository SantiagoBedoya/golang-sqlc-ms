-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (email, username, password) VALUES ($1, $2, $3) RETURNING *;

-- name: DeleteUser :exec
UPDATE users SET is_deleted = TRUE where id = $1;