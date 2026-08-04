-- +goose Up
CREATE TABLE job_units (
    job_id UUID REFERENCES jobs(id),
    unit_id UUID REFERENCES units(id),
    assigned_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (job_id, unit_id)
);


CREATE INDEX idx_job_units_job_id ON job_units(job_id);
CREATE INDEX idx_job_units_unit_id ON job_units(unit_id);


-- +goose Down
DROP TABLE IF EXISTS job_units;
