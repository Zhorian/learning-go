BEGIN;

ALTER TABLE books rename COLUMN created_at to create_at;

COMMIT