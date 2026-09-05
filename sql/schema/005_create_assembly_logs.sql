-- +goose up
CREATE TABLE assembly_logs(
    id UUID PRIMARY KEY,
    sub_assembly_id UUID NOT NULL REFERENCES sub_assemblies(id),
    file_path TEXT NOT NULL,
    file_hash TEXT NOT NULL,
    captured_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ

);

CREATE INDEX idx_assembly_logs_dedup ON assembly_logs(sub_assembly_id, file_hash);

-- +goose down
DROP TABLE IF EXISTS assembly_logs;
