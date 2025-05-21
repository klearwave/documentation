-- +goose Up
CREATE TABLE container_images (
    id SERIAL PRIMARY KEY,
    image VARCHAR(512) NOT NULL UNIQUE,
    image_registry VARCHAR(255) NOT NULL,
    image_name VARCHAR(255) NOT NULL,
    image_tag VARCHAR(128) NOT NULL,
    sha256_sum CHAR(64) NOT NULL,
    commit_hash CHAR(40) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE version_container_images (
    id SERIAL PRIMARY KEY,
    version_id VARCHAR(32),
    container_image_id INTEGER,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (version_id, container_image_id),
    FOREIGN KEY (version_id) REFERENCES versions (id) ON DELETE CASCADE,
    FOREIGN KEY (container_image_id) REFERENCES container_images (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE version_container_images;
DROP TABLE container_images;
