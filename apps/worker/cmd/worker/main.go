package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, os.Getenv("POSTGRES_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	log.Println("worker started")
	for {
		// naive poll: select findings without suggestion where rule.llm_check=true
		rows, _ := pool.Query(ctx, `
      select f.id, c.text, r.guidance
      from finding f
      join rule r on r.id=f.rule_id
      join clause c on c.id=f.clause_id
      where (f.suggestion is null or f.suggestion='')
        and r.llm_check=true
      limit 5`)
		type item struct {
			ID               int64
			Clause, Guidance string
		}
		items := []item{}
		for rows.Next() {
			var it item
			rows.Scan(&it.ID, &it.Clause, &it.Guidance)
			items = append(items, it)
		}
		for _, it := range items {
			// call LLM
			rationale, suggestion, risk := callLLM(it.Clause, it.Guidance)
			_, _ = pool.Exec(ctx, `update finding set rationale=$1, suggestion=$2, severity=$3 where id=$4`,
				rationale, suggestion, risk, it.ID)
		}
		time.Sleep(2 * time.Second)
	}
}
