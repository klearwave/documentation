-- +goose Up
CREATE TABLE versions (
    id SERIAL PRIMARY KEY,
    version_id VARCHAR(32) NOT NULL UNIQUE,
    stable BOOLEAN DEFAULT FALSE,
    x_version INT,
    y_version INT,
    z_version INT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
);

-- +goose Down
DROP TABLE versions;
