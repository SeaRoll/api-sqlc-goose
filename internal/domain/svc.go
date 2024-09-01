package domain

import (
	"api-sqlc-goose/api"
	"api-sqlc-goose/internal/database"
	"api-sqlc-goose/internal/database/db"
	"context"
)

type Service interface {
	CreateBook(ctx context.Context, data api.CreateBookDTO, eq ...db.Querier) (api.BookDTO, error)
	GetBooks(ctx context.Context, eq ...db.Querier) ([]api.BookDTO, error)
}

type service struct {
	bookMapper BookMapper
	db         database.Database
}

func NewService(db database.Database) Service {
	return &service{
		bookMapper: NewBookMapper(),
		db:         db,
	}
}

// GetBooks implements Service.
func (s *service) GetBooks(ctx context.Context, eq ...db.Querier) ([]api.BookDTO, error) {
	var dbBooks []db.Book
	if err := s.db.WithoutTX(func(q db.Querier) error {
		var err error
		dbBooks, err = q.GetBooks(ctx)
		return err
	}, eq...); err != nil {
		return nil, err
	}

	books := []api.BookDTO{}
	for _, dbBook := range dbBooks {
		books = append(books, s.bookMapper.ToDTO(dbBook))
	}

	return books, nil
}

// CreateBook implements Service.
func (s *service) CreateBook(ctx context.Context, data api.CreateBookDTO, eq ...db.Querier) (api.BookDTO, error) {
	var dbBook db.Book
	if err := s.db.WithTX(ctx, func(q db.Querier) error {
		var err error
		dbBook, err = q.CreateBook(ctx, s.bookMapper.ToCreateBookParams(data))
		return err
	}, eq...); err != nil {
		return api.BookDTO{}, err
	}
	return s.bookMapper.ToDTO(dbBook), nil
}
