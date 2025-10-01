package search

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Result struct {
	ID      int64   `json:"id"`
	Title   string  `json:"title"`
	Snippet string  `json:"snippet"`
	Score   float64 `json:"score"`
	Tags    string  `json:"tags"`
}

// SearchPrecedents is a placeholder implementation that matches the handler signature.
// It should eventually run hybrid search (keywords + embeddings) in DB.
func SearchPrecedents(ctx context.Context, db *pgxpool.Pool, q string) ([]Result, error) {
	// TODO: implement hybrid search
	return []Result{}, nil
}
