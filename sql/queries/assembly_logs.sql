-- name: CheckDuplicatePhoto :one
SELECT EXISTS (
    SELECT 1
    FROM assembly_logs
    WHERE sub_assembly_id = $1
      AND file_hash = $2
      AND deleted_at IS NULL
);

-- name: CreateAssemblyLog :one
INSERT INTO assembly_logs (
    id,
    sub_assembly_id,
    file_path,
    file_hash,
    captured_at
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetPhotosBySubAssembly :many
SELECT *
FROM assembly_logs
WHERE sub_assembly_id = $1
    AND deleted_at IS NULL
ORDER BY captured_at DESC
LIMIT $2 OFFSET $3;

-- name: DeleteAssemblyLog :exec
UPDATE assembly_logs
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted_at IS NULL;
