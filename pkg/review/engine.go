package review

import (
	"context"
	"regexp"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Rule struct {
	ID                                int64
	Name, Severity, Pattern, Guidance string
	LLMCheck                          bool
}

func RunRules(ctx context.Context, pool *pgxpool.Pool, documentID, playbookID int64) error {
	rows, err := pool.Query(ctx, `select id,text from clause where document_id=$1`, documentID)
	if err != nil {
		return err
	}
	clauses := map[int64]string{}
	for rows.Next() {
		var id int64
		var text string
		rows.Scan(&id, &text)
		clauses[id] = text
	}
	rws, err := pool.Query(ctx, `select id,name,severity,pattern,guidance,llm_check from rule where playbook_id=$1`, playbookID)
	if err != nil {
		return err
	}
	for rws.Next() {
		var rl Rule
		rws.Scan(&rl.ID, &rl.Name, &rl.Severity, &rl.Pattern, &rl.Guidance, &rl.LLMCheck)
		rx, _ := regexp.Compile(rl.Pattern)
		for cid, ctext := range clauses {
			if rx.MatchString(ctext) {
				// initial finding (LLM to fill later)
				_, _ = pool.Exec(ctx, `insert into finding(document_id,clause_id,rule_id,severity,rationale) values($1,$2,$3,$4,$5)`,
					documentID, cid, rl.ID, rl.Severity, "Pattern matched: "+rl.Name)
			}
		}
	}
	return nil
}
