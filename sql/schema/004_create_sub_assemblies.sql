-- +goose Up
CREATE TABLE sub_assemblies (
    id UUID PRIMARY KEY,
    unit_id UUID REFERENCES units(id) NOT NULL,
    name TEXT NOT NULL,
    serial_number TEXT,
    status TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);


CREATE INDEX idx_sub_assemblies_unit_id ON sub_assemblies(unit_id);


-- +goose Down
DROP TABLE IF EXISTS sub_assemblies;
