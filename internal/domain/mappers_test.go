package domain_test

import (
	"api-sqlc-goose/api"
	"api-sqlc-goose/internal/database/db"
	"api-sqlc-goose/internal/domain"
	"testing"
)

type testArgs[A any, E any] struct {
	args     A
	expected E
}

func TestBookMapperToDTO(t *testing.T) {
	tests := []testArgs[db.Book, api.BookDTO]{
		{
			args: db.Book{
				ID:          []byte("123"),
				Title:       "title",
				Author:      "author",
				Description: "description",
				CreatedAt:   0,
				UpdatedAt:   0,
			},
			expected: api.BookDTO{
				Id:          "123",
				Title:       "title",
				Author:      "author",
				Description: "description",
				CreatedAt:   "1970-01-01T00:00:00Z",
				UpdatedAt:   "1970-01-01T00:00:00Z",
			},
		},
	}

	for _, tt := range tests {
		mapper := domain.NewBookMapper()
		got := mapper.ToDTO(tt.args)
		if got != tt.expected {
			t.Errorf("ToDTO(%v) = %v; want %v", tt.args, got, tt.expected)
		}
	}
}
