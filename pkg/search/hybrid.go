package search

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SearchPrecedents(ctx context.Context, pool *pgxpool.Pool, q string) ([]map[string]any, error) {
	// MVP: keyword LIKE; vector search later
	ql := "%" + strings.ToLower(q) + "%"
	rows, err := pool.Query(ctx, `select id, title, text, tags from precedent where lower(title) like $1 or lower(text) like $1 limit 10`, ql)
	if err != nil {
		return nil, err
	}
	out := []map[string]any{}
	for rows.Next() {
		var id int64
		var title, text, tags string
		rows.Scan(&id, &title, &text, &tags)
		out = append(out, map[string]any{"id": id, "title": title, "text": text, "tags": tags})
	}
	return out, nil
}
