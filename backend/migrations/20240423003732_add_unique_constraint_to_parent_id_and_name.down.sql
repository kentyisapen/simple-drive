ALTER TABLE items ADD CONSTRAINT unique_parent_id_name UNIQUE (parent_id, name);
