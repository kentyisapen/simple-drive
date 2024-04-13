CREATE TYPE item_type AS ENUM('FILE', 'DIRECTORY');

CREATE TABLE items (
    id UUID PRIMARY KEY, parent_id UUID NULL, thumbnail_id UUID NULL, name VARCHAR(255) NOT NULL, type item_type NOT NULL, size BIGINT NULL, created_at TIMESTAMP NULL, last_modified_at TIMESTAMP NULL
);
