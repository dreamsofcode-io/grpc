package database

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context) (*pgxpool.Pool, error) {
	dbURL, exists := os.LookupEnv("DATABASE_URL")
	if !exists {
		return nil, errors.New("DATABASE_URL not set")
	}

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	migrationsPath, exists := os.LookupEnv("MIGRATIONS_PATH")
	if !exists {
		return nil, errors.New("MIGRATIONS_PATH not set")
	}

	source := fmt.Sprintf("file://%s", migrationsPath)

	m, err := migrate.New(source, dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create migrations: %w", err)
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	return pool, nil
}
