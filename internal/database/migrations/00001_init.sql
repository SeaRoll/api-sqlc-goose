-- +goose Up
CREATE TABLE books (
    id BLOB PRIMARY KEY,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at INTEGER NOT NULL, -- Unix timestamp
    updated_at INTEGER NOT NULL -- Unix timestamp
);

-- +goose Down
DROP TABLE books;
