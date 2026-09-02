-- +goose up
CREATE TABLE assembly_logs(
    id UUID PRIMARY KEY,
    sub_assembly_id UUID REFERENCES sub_assemblies(id),
    file_path TEXT NOT NULL,
    file_hash TEXT NOT NULL,
    captured_at TIMESTAMPTZ NOT NULL,

)
