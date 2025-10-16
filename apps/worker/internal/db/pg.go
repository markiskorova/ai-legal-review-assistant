package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context) *pgxpool.Pool {
	dsn := getenv("DATABASE_URL", "postgres://postgres:postgres@db:5432/legal_assistant?sslmode=disable")
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("db parse error: %v", err)
	}
	cfg.MaxConns = 10
	cfg.MinConns = 1
	cfg.MaxConnLifetime = 30 * time.Minute
	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		log.Fatalf("db connect error: %v", err)
	}
	return pool
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
