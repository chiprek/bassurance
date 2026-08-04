-- +goose Up

CREATE EXTENSION IF NOT EXISTS moddatetime;

CREATE TABLE jobs (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    status TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);


CREATE TRIGGER jobs_updated_at_trigger
BEFORE UPDATE ON jobs
FOR EACH ROW
EXECUTE PROCEDURE moddatetime(updated_at);


-- +goose Down
DROP TRIGGER IF EXISTS jobs_updated_at_trigger ON jobs;
DROP TABLE IF EXISTS jobs;
