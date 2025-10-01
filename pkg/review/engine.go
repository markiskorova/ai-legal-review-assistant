package review

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// RunRules is a placeholder implementation that matches the handler signature.
// It should eventually scan the document, apply rules, and insert findings into DB.
func RunRules(ctx context.Context, db *pgxpool.Pool, documentID int64, playbookID int64) error {
	// TODO: implement review pipeline
	return nil
}
