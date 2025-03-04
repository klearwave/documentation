-- +goose Up
CREATE TABLE container_images (
    id SERIAL PRIMARY KEY,
    image VARCHAR(512) NOT NULL UNIQUE,
    image_path VARCHAR(255) NOT NULL,
    image_tag VARCHAR(128) NOT NULL,
    sha256_sum VARCHAR(64) NOT NULL,
    commit_hash VARCHAR(40) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE version_container_images (
    id SERIAL PRIMARY KEY,
    version_id INTEGER,
    container_image_id INTEGER,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (version_id, container_image_id),
    FOREIGN KEY (version_id) REFERENCES versions (id) ON DELETE CASCADE,
    FOREIGN KEY (container_image_id) REFERENCES container_images (id) ON DELETE CASCADE
);

-- Function to automatically split image into path and tag
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION trigger_split_image()
RETURNS TRIGGER AS $$
BEGIN
    NEW.image_path = 
        CASE 
            WHEN POSITION(':' IN NEW.image) > 0 THEN SPLIT_PART(NEW.image, ':', 1)
            ELSE NEW.image
        END;
    NEW.image_tag = 
        CASE 
            WHEN POSITION(':' IN NEW.image) > 0 THEN SPLIT_PART(NEW.image, ':', 2)
            ELSE 'latest'
        END;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- Create a trigger to call the function before inserting data
CREATE TRIGGER trigger_split_image
BEFORE INSERT ON container_images
FOR EACH ROW
EXECUTE FUNCTION trigger_split_image();

-- +goose Down
DROP TABLE version_container_images;
DROP TABLE container_images;
DROP FUNCTION IF EXISTS trigger_split_image;
