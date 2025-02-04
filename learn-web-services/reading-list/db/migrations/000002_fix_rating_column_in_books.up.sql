BEGIN;

ALTER TABLE books rename COLUMN ratings to rating;

COMMIT