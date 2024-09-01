package main

import (
	"api-sqlc-goose/api"
	"api-sqlc-goose/internal/database"
	"api-sqlc-goose/internal/domain"
	"api-sqlc-goose/internal/server"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

func main() {
	dbo := database.MustInit()
	service := domain.NewService(dbo)
	srv := server.NewServer(service)
	srvWrapper := api.NewStrictHandler(srv, []api.StrictMiddlewareFunc{})
	r := http.NewServeMux()
	h := api.HandlerFromMux(srvWrapper, r)

	// cors middleware
	mw := cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	})
	h = mw(h)

	// recover middleware
	mw = middleware.Recoverer
	h = mw(h)

	// http rate limiting
	mw = httprate.LimitByIP(100, 1*time.Minute)
	h = mw(h)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(s.ListenAndServe())
}
