package main

import (
	"context"
	"encoding/json"
	"net/http"

	dbase "github.com/markiskorova/ai-legal-review-assistant/apps/api/internal/db"
	"github.com/markiskorova/ai-legal-review-assistant/pkg/search"
)

func handleSearch(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	ctx := context.Background()
	pool, _ := dbase.Pool(ctx)
	defer pool.Close()
	res, err := search.SearchPrecedents(ctx, pool, q)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(res)
}
