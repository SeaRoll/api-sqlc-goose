package database

import (
	"api-sqlc-goose/internal/database/db"
	"context"
	"database/sql"
	"embed"
	"log"

	_ "github.com/glebarez/go-sqlite"
	"github.com/pressly/goose/v3"
)

type Database interface {
	WithTX(ctx context.Context, fn func(q db.Querier) error, eq ...db.Querier) error
	WithoutTX(fn func(q db.Querier) error, eq ...db.Querier) error
}

type database struct {
	db *sql.DB
}

// WithTX implements Database.
func (d *database) WithTX(ctx context.Context, fn func(q db.Querier) error, eq ...db.Querier) error {
	if len(eq) > 0 {
		return fn(eq[0])
	}
	queries := db.New(d.db)
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	if err := fn(queries); err != nil {
		return err
	}

	return tx.Commit()
}

// WithoutTX implements Database.
func (d *database) WithoutTX(fn func(q db.Querier) error, eq ...db.Querier) error {
	if len(eq) > 0 {
		return fn(eq[0])
	}
	queries := db.New(d.db)
	return fn(queries)
}

func MustInit() Database {
	db, err := sql.Open("sqlite", "db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	migrate(db)
	return &database{db: db}
}

//go:embed migrations/*.sql
var embedMigrations embed.FS

func migrate(db *sql.DB) {
	// setup database connection
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("sqlite"); err != nil {
		log.Fatalf("failed to set dialect: %v", err)
	}
	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatalf("failed to migrate: %v", err)
		panic(err)
	}
}
