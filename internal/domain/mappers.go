package domain

import (
	"api-sqlc-goose/api"
	"api-sqlc-goose/internal/database/db"
	"time"

	"github.com/google/uuid"
)

type BookMapper interface {
	ToDTO(dbBook db.Book) api.BookDTO
	ToCreateBookParams(data api.CreateBookDTO) db.CreateBookParams
}

type bookMapper struct{}

func NewBookMapper() BookMapper {
	return &bookMapper{}
}

// ToCreateBookParams implements BookMapper.
func (m *bookMapper) ToCreateBookParams(data api.CreateBookDTO) db.CreateBookParams {
	return db.CreateBookParams{
		ID:          []byte(uuid.New().String()),
		Title:       data.Title,
		Author:      data.Author,
		Description: data.Description,
		CreatedAt:   time.Now().UTC().Unix(),
		UpdatedAt:   time.Now().UTC().Unix(),
	}
}

func (m *bookMapper) ToDTO(dbBook db.Book) api.BookDTO {
	createdAtStr := time.Unix(dbBook.CreatedAt, 0).UTC().Format(time.RFC3339)
	updatedAtStr := time.Unix(dbBook.UpdatedAt, 0).UTC().Format(time.RFC3339)

	return api.BookDTO{
		Id:          string(dbBook.ID),
		Title:       dbBook.Title,
		Author:      dbBook.Author,
		Description: dbBook.Description,
		CreatedAt:   createdAtStr,
		UpdatedAt:   updatedAtStr,
	}
}
