BEGIN;

ALTER TABLE books rename COLUMN rating to ratings;

COMMIT