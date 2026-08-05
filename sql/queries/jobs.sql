-- name: GetJobs :many
SELECT * FROM jobs
WHERE deleted_at IS NULL
ORDER BY created_at ASC;

-- name: GetJob :one
SELECT * FROM jobs
WHERE name = $1 and deleted_at IS NULL
LIMIT 1;

-- name: CreateJob :one
INSERT INTO jobs (id, name, status, created_at, updated_at)
VALUES(
    $1, $2, $3, NOW(), NOW()
)
RETURNING *;

-- name: SoftDeleteJob :exec
UPDATE jobs
SET deleted_at = NOW()
WHERE id = $1;
