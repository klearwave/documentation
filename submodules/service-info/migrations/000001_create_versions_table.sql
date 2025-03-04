-- +goose Up
CREATE TABLE versions (
    id SERIAL PRIMARY KEY,
    version_id VARCHAR(32) NOT NULL UNIQUE,
    latest BOOLEAN DEFAULT FALSE,
    x_version INT GENERATED ALWAYS AS (CAST(SPLIT_PART(SUBSTRING(version_id FROM 2), '.', 1) AS INT)) STORED,
    y_version INT GENERATED ALWAYS AS (CAST(SPLIT_PART(SUBSTRING(version_id FROM 2), '.', 2) AS INT)) STORED,
    z_version INT GENERATED ALWAYS AS (CAST(SPLIT_PART(SUBSTRING(version_id FROM 2), '.', 3) AS INT)) STORED,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Function to validate semantic versioning format
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION trigger_validate_version_id()
RETURNS TRIGGER AS $$
BEGIN
    IF NOT NEW.version_id ~ '^v\d+\.\d+\.\d+$' THEN
        RAISE EXCEPTION 'Invalid version_id format. Must be in the form vX.Y.Z';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- Function to set all versions to latest=false if a new version is submitted with latest=true
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION trigger_ensure_single_latest()
RETURNS TRIGGER AS $$
BEGIN
    -- If the new row is set to latest=true
    IF NEW.latest = TRUE THEN
        -- Update all other rows to set latest=false
        UPDATE versions
        SET latest = FALSE
        WHERE latest = TRUE;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- Ensure version_id starts with "v" and follows the semantic versioning format
CREATE TRIGGER trigger_validate_version_id
BEFORE INSERT OR UPDATE ON versions
FOR EACH ROW
EXECUTE FUNCTION trigger_validate_version_id();

-- Ensure only one latest version exists
CREATE TRIGGER trigger_ensure_single_latest
BEFORE INSERT OR UPDATE ON versions
FOR EACH ROW
EXECUTE FUNCTION trigger_ensure_single_latest();

-- +goose Down
DROP TABLE versions;
DROP FUNCTION IF EXISTS trigger_ensure_single_latest;
DROP FUNCTION IF EXISTS trigger_validate_version_id;
