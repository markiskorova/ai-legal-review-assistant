package main

import (
	"context"
	"encoding/json"
	"net/http"

	dbase "github.com/markiskorova/ai-legal-review-assistant/apps/api/internal/db"
)

type matterReq struct {
	Title, Requester, ContractType, DueDate string
}

func handleCreateMatter(w http.ResponseWriter, r *http.Request) {
	var req matterReq
	json.NewDecoder(r.Body).Decode(&req)
	ctx := context.Background()
	pool, _ := dbase.Pool(ctx)
	defer pool.Close()
	var id int64
	err := pool.QueryRow(ctx, `insert into matter(org_id,title,requester,contract_type,due_date,status)
    values(1,$1,$2,$3,$4,'open') returning id`,
		req.Title, req.Requester, req.ContractType, req.DueDate).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(map[string]any{"id": id})
}

func handleListMatters(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	pool, _ := dbase.Pool(ctx)
	defer pool.Close()
	rows, _ := pool.Query(ctx, `select id,title,requester,contract_type,status from matter order by id desc limit 50`)
	type m struct {
		ID                                     int64
		Title, Requester, ContractType, Status string
	}
	out := []m{}
	for rows.Next() {
		var x m
		rows.Scan(&x.ID, &x.Title, &x.Requester, &x.ContractType, &x.Status)
		out = append(out, x)
	}
	json.NewEncoder(w).Encode(out)
}
