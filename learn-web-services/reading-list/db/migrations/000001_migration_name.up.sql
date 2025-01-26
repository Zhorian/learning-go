BEGIN;

CREATE TABLE books (
    id bigserial PRIMARY KEY,
    create_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    published integer NOT NULL,
    pages integer NOT NULL,
    genres text[] NOT NULL,
    ratings real NOT NULL,
    version integer NOT NULL DEFAULT 1    
);

COMMIT;