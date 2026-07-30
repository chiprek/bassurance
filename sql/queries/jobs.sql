-- name: GetJOBS :many
SELECT * FROM jobs
ORDER BY created_at ASC;

-- name: CreateJob :one
INSERT INTO jobs (id, created_at, updated_at, name, status)
VALUES(
    gen_random_uuid(), NOW(), NOW(), $1, $2
)
RETURNING id, name, status, created_at;
