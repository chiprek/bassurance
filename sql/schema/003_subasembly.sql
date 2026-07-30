-- +goose Up
CREATE TABLE sub_assemblies (
    id UUID PRIMARY KEY,
    unit_id UUID NOT NULL REFERENCES units(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    part_serial_number TEXT,
    status TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_sub_assemblies_unit_id ON sub_assemblies(unit_id);

CREATE TRIGGER sub_assemblies_updated_at_trigger
BEFORE UPDATE ON sub_assemblies
FOR EACH ROW
EXECUTE PROCEDURE moddatetime(updated_at);

-- +goose Down
DROP TRIGGER IF EXISTS sub_assemblies_updated_at_trigger ON sub_assemblies;
DROP TABLE IF EXISTS sub_assemblies;
