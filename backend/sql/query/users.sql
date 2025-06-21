-- name: CreateUser :one
INSERT INTO users (
    username,
    password_hash
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY username ASC;

-- name: UpdateUserPassword :one
UPDATE users
SET 
    password_hash = $2
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1; 