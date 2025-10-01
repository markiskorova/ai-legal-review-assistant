package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Pool(ctx context.Context) (*pgxpool.Pool, error) {
	dsn := os.Getenv("POSTGRES_DSN")
	return pgxpool.New(ctx, dsn)
}
