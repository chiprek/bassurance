-- +goose up
CREATE TABLE assembly_logs(
    id UUID PRIMARY KEY,
    sub_assembly_id UUID REFERENCES sub_assemblies(id),
    file_path TEXT NOT NULL,
    file_hash TEXT NOT NULL,
    captured_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ,

)

CREATE INDEX idx_assembly_logs_if ON assembly_logs(sub_assembly_id, file_hash);
