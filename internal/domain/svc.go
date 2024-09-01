package domain

import (
	"api-sqlc-goose/api"
	"api-sqlc-goose/internal/database"
	"api-sqlc-goose/internal/database/db"
	"context"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	CreateBook(ctx context.Context, data api.CreateBookDTO, eq ...db.Querier) (api.BookDTO, error)
	GetBooks(ctx context.Context, eq ...db.Querier) ([]api.BookDTO, error)
}

type service struct {
	db database.Database
}

func NewService(db database.Database) Service {
	return &service{db: db}
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
		createdAtStr := time.Unix(dbBook.CreatedAt, 0).Format(time.RFC3339)
		updatedAtStr := time.Unix(dbBook.UpdatedAt, 0).Format(time.RFC3339)

		books = append(books, api.BookDTO{
			Id:          string(dbBook.ID),
			Title:       dbBook.Title,
			Author:      dbBook.Author,
			Description: dbBook.Description,
			CreatedAt:   createdAtStr,
			UpdatedAt:   updatedAtStr,
		})
	}

	return books, nil
}

// CreateBook implements Service.
func (s *service) CreateBook(ctx context.Context, data api.CreateBookDTO, eq ...db.Querier) (api.BookDTO, error) {
	var dbBook db.Book
	if err := s.db.WithTX(ctx, func(q db.Querier) error {
		var err error
		dbBook, err = q.CreateBook(ctx, db.CreateBookParams{
			ID:          []byte(uuid.New().String()),
			Title:       data.Title,
			Author:      data.Author,
			Description: data.Description,
			CreatedAt:   time.Now().Unix(),
			UpdatedAt:   time.Now().Unix(),
		})
		return err
	}, eq...); err != nil {
		return api.BookDTO{}, err
	}

	createdAtStr := time.Unix(dbBook.CreatedAt, 0).Format(time.RFC3339)
	updatedAtStr := time.Unix(dbBook.UpdatedAt, 0).Format(time.RFC3339)

	return api.BookDTO{
		Id:          string(dbBook.ID),
		Title:       dbBook.Title,
		Author:      dbBook.Author,
		Description: dbBook.Description,
		CreatedAt:   createdAtStr,
		UpdatedAt:   updatedAtStr,
	}, nil
}
