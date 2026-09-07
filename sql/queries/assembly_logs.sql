-- name: CheckDuplicatePhoto :one
SELECT EXISTS (
    SELECT 1
    FROM assembly_logs
    WHERE sub_assembly_id = $1
      AND file_hash = $2
      AND deleted_at IS NULL
);
