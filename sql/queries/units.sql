-- name: GetUnits :many
SELECT * FROM units
ORDER BY created_at ASC;

-- name: GetUnit :one
SELECT * FROM units
WHERE $1 = serial_number;

-- name: CreateUnit :one
INSERT INTO units (id, job_id, serial_number, status,  created_at, updated_at)
VALUES(
    gen_random_uuid(), $1, $2, $3, NOW(), NOW()
)
RETURNING id, job_id, serial_number, status, created_at;
