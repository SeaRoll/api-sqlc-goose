-- name: CreateBook :one
INSERT INTO books
(id, title, author, description, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetBooks :many
SELECT * FROM books;
