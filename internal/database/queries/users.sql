-- name: CreateUser :one
INSERT INTO users (username, email, password_hash, role)
VALUES ($1, $2, $3, $4)
RETURNING id, username, email, password_hash, role, created_at;

-- name: GetUserByID :one
SELECT id, username, email, password_hash, role, created_at
FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT id, username, email, password_hash, role, created_at
FROM users
WHERE email = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET username = $2, email = $3
WHERE id = $1
RETURNING id, username, email, password_hash, role, created_at;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;