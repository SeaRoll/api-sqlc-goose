// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	CreateBook(ctx context.Context, arg CreateBookParams) (Book, error)
	GetBooks(ctx context.Context) ([]Book, error)
}

var _ Querier = (*Queries)(nil)
