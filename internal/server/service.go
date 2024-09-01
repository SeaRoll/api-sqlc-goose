package server

import (
	"api-sqlc-goose/api"
	"api-sqlc-goose/internal/domain"
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
)

var _ api.StrictServerInterface = (*server)(nil)

type server struct {
	service domain.Service
}

func NewServer(service domain.Service) api.StrictServerInterface {
	return &server{service: service}
}

// GetApiV1Books implements api.StrictServerInterface.
func (s *server) GetApiV1Books(ctx context.Context, request api.GetApiV1BooksRequestObject) (api.GetApiV1BooksResponseObject, error) {
	books, err := s.service.GetBooks(ctx)
	if err != nil {
		return api.GetApiV1Books500JSONResponse{Message: err.Error()}, nil
	}
	return api.GetApiV1Books200JSONResponse(books), nil
}

// PostApiV1Books implements api.StrictServerInterface.
func (s *server) PostApiV1Books(ctx context.Context, request api.PostApiV1BooksRequestObject) (api.PostApiV1BooksResponseObject, error) {
	book, err := s.service.CreateBook(ctx, *request.Body)
	if err != nil {
		return api.PostApiV1Books500JSONResponse{Message: err.Error()}, nil
	}
	return api.PostApiV1Books201JSONResponse(book), nil
}

// GetPing implements api.StrictServerInterface.
func (s *server) GetPing(ctx context.Context, request api.GetPingRequestObject) (api.GetPingResponseObject, error) {
	// generate a random int from 0 to 10
	n, err := rand.Int(rand.Reader, big.NewInt(10))
	if err != nil {
		return api.GetPing500JSONResponse{Message: "error"}, nil
	}

	// to int
	i := int(n.Int64())
	if i < 1 {
		panic("should not happen")
	}
	if i < 3 {
		return api.GetPing400JSONResponse{ErrorJSONResponse: api.ErrorJSONResponse{Message: fmt.Sprintf("Result %d is less than3", i)}}, nil
	}

	return api.GetPing200JSONResponse{
		Ping: "pong",
	}, nil
}
