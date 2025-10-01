package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	dbase "github.com/markiskorova/ai-legal-review-assistant/apps/api/internal/db"
)

type docReq struct {
	Title    string `json:"title"`
	Text     string `json:"text"`
	MatterID *int64 `json:"matter_id"`
}

func handleCreateDocument(w http.ResponseWriter, r *http.Request) {
	var req docReq
	json.NewDecoder(r.Body).Decode(&req)
	ctx := context.Background()
	pool, _ := dbase.Pool(ctx)
	defer pool.Close()
	var docID int64
	err := pool.QueryRow(ctx, `insert into document(org_id,matter_id,title,mime,text) values(1,$1,$2,'text/plain',$3) returning id`,
		req.MatterID, req.Title, req.Text).Scan(&docID)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// naive clause split
	parts := splitClauses(req.Text)
	idx := 0
	for _, p := range parts {
		start := idx
		end := idx + len(p)
		_, _ = pool.Exec(ctx, `insert into clause(document_id,start_idx,end_idx,text) values($1,$2,$3,$4)`,
			docID, start, end, p)
		idx = end + 1
	}
	json.NewEncoder(w).Encode(map[string]any{"id": docID, "clauses": len(parts)})
}

func splitClauses(s string) []string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	raw := strings.Split(s, "\n\n")
	out := make([]string, 0, len(raw))
	for _, r := range raw {
		t := strings.TrimSpace(r)
		if t != "" {
			out = append(out, t)
		}
	}
	return out
}
