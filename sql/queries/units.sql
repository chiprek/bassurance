-- name: GetUnits :many
SELECT * FROM units
WHERE deleted_at IS NULL
ORDER BY
    CASE WHEN @sort_direction::text = 'desc' THEN created_at END DESC,
    CASE WHEN @sort_direction::text = 'asc' THEN created_at END ASC,
    created_at ASC
LIMIT $1 OFFSET $2;

-- name: GetUnit :one
SELECT * FROM units
WHERE serial_number = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: GetUnitsByJobName :many
SELECT units.* FROM units
JOIN job_units ON units.id = job_units.unit_id
JOIN jobs ON job_units.job_id = jobs.id
WHERE jobs.name = $1 AND units.deleted_at IS NULL
LIMIT $2 OFFSET $3;

-- name: GetJobsByUnitSerialNumber :many
SELECT jobs.* FROM jobs
JOIN job_units ON jobs.id = job_units.job_id
JOIN units ON job_units.unit_id = units.id
WHERE units.serial_number = $1 AND jobs.deleted_at IS NULL
LIMIT $2 OFFSET $3;

-- name: GetUnitBySerialNumber :one
SELECT id FROM units
WHERE serial_number = $1 AND deleted_at IS NULL;

-- name: CreateUnit :one
INSERT INTO units (id, serial_number, created_at)
VALUES(
    $1, $2, NOW()
)
RETURNING *;

-- name: CreateJobUnit :exec
INSERT INTO job_units (job_id, unit_id, assigned_at)
VALUES (
    $1, $2, NOW()
);
