-- +goose Up
CREATE TABLE units (
    id UUID PRIMARY KEY,
    serial_number TEXT UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);


-- +goose Down
DROP TABLE IF EXISTS units;
