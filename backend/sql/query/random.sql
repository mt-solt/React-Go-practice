-- name: CreateRandom :exec
INSERT INTO random (id, random_val, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5);

-- name: GetRandom :one
SELECT id, random_val, user_id, created_at, updated_at
FROM random
WHERE id = $1;

-- name: UpdateRandom :exec
UPDATE random
SET random_val = $2, updated_at = $3
WHERE id = $1;

-- name: DeleteRandom :exec
DELETE FROM random
WHERE id = $1;

-- name: GetAllRandoms :many
SELECT id, random_val, user_id, created_at, updated_at
FROM random
ORDER BY id ASC;

-- name: GetRandomsByUser :many
SELECT id, random_val, user_id, created_at, updated_at
FROM random
WHERE user_id = $1
ORDER BY id ASC; 