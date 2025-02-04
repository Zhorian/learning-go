BEGIN;

ALTER TABLE books rename COLUMN create_at to created_at;

COMMIT