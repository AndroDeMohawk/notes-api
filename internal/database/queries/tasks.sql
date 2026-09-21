-- name: CreateTask :one
INSERT INTO tasks (user_id, title, description, status)
VALUES ($1, $2, $3, $4)
RETURNING id, user_id, title, description, status, created_at;

-- name: GetTaskByID :one
SELECT id, user_id, title, description, status, created_at
FROM tasks
WHERE id = $1 LIMIT 1;

-- name: ListTasksByUserID :many
SELECT id, user_id, title, description, status, created_at
FROM tasks
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: UpdateTask :one
UPDATE tasks
SET title = $3, description = $4, status = $5
WHERE id = $1 AND user_id = $2
RETURNING id, user_id, title, description, status, created_at;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1 AND user_id = $2;