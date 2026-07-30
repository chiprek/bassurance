-- +goose Up
CREATE TABLE units (
    id UUID PRIMARY KEY,
    job_id UUID NOT NULL REFERENCES jobs(id) ON DELETE CASCADE,
    serial_number TEXT UNIQUE,
    status TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_units_job_id ON units(job_id);

CREATE TRIGGER units_updated_at_trigger
BEFORE UPDATE ON units
FOR EACH ROW
EXECUTE PROCEDURE moddatetime(updated_at);

-- +goose Down
DROP TRIGGER IF EXISTS units_updated_at_trigger ON units;
DROP TABLE IF EXISTS units;
