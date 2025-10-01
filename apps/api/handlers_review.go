package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	dbase "github.com/markiskorova/ai-legal-review-assistant/apps/api/internal/db"
	"github.com/markiskorova/ai-legal-review-assistant/pkg/review"
)

func handleStartReview(w http.ResponseWriter, r *http.Request) {
	docID, _ := strconv.ParseInt(chi.URLParam(r, "document_id"), 10, 64)
	pbID, _ := strconv.ParseInt(r.URL.Query().Get("playbook_id"), 10, 64)
	ctx := context.Background()
	pool, _ := dbase.Pool(ctx)
	defer pool.Close()
	if err := review.RunRules(ctx, pool, docID, pbID); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(map[string]any{"started": true})
}

func handleGetFindings(w http.ResponseWriter, r *http.Request) {
	docID, _ := strconv.ParseInt(chi.URLParam(r, "document_id"), 10, 64)
	ctx := context.Background()
	pool, _ := dbase.Pool(ctx)
	defer pool.Close()
	rows, _ := pool.Query(ctx, `
    select f.id, f.severity, f.rationale, f.suggestion, c.text as clause_text, r.name as rule_name
    from finding f
    join clause c on c.id=f.clause_id
    join rule r on r.id=f.rule_id
    where f.document_id=$1
    order by f.id desc`, docID)
	type resp struct {
		ID                                                    int64
		Severity, Rationale, Suggestion, ClauseText, RuleName string
	}
	out := []resp{}
	for rows.Next() {
		var x resp
		rows.Scan(&x.ID, &x.Severity, &x.Rationale, &x.Suggestion, &x.ClauseText, &x.RuleName)
		out = append(out, x)
	}
	json.NewEncoder(w).Encode(out)
}
