package jobs

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Rule struct {
	ID       int64
	Pattern  string
	LLMCheck bool
	Severity string
	Guidance string // optional: add a guidance column in schema if you want
}

type Clause struct {
	ID   int64
	Text string
}

type LLM interface {
	Validate(ctx context.Context, clause, guidance string) (risk, rationale, suggestion string, err error)
}

func ReviewDocument(ctx context.Context, db *pgxpool.Pool, llm LLM, documentID int64, playbookID *int64) error {
	clauses, err := loadClauses(ctx, db, documentID)
	if err != nil {
		return fmt.Errorf("loadClauses: %w", err)
	}
	rules, err := loadRules(ctx, db, playbookID)
	if err != nil {
		return fmt.Errorf("loadRules: %w", err)
	}

	for _, c := range clauses {
		for _, r := range rules {
			ok, matchErr := regexMatch(r.Pattern, c.Text)
			if matchErr != nil {
				log.Printf("bad regex (rule %d): %v", r.ID, matchErr)
				continue
			}
			if !ok {
				continue
			}

			risk := r.Severity
			rationale := "Pattern matched"
			suggestion := ""

			if r.LLMCheck && llm != nil {
				rrisk, rrationale, rsugg, lerr := llm.Validate(ctx, c.Text, r.Guidance)
				if lerr == nil {
					if rrisk != "" {
						risk = rrisk
					}
					if rrationale != "" {
						rationale = rrationale
					}
					suggestion = rsugg
				} else {
					log.Printf("LLM validate error: %v", lerr)
				}
			}

			if err := insertFinding(ctx, db, r.ID, c.ID, risk, rationale, suggestion); err != nil {
				log.Printf("insertFinding error: %v", err)
			}
		}
	}

	return nil
}

func loadClauses(ctx context.Context, db *pgxpool.Pool, docID int64) ([]Clause, error) {
	rows, err := db.Query(ctx, `select id, text from clause where document_id=$1 order by id asc`, docID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Clause
	for rows.Next() {
		var c Clause
		if err := rows.Scan(&c.ID, &c.Text); err != nil {
			return nil, err
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

func loadRules(ctx context.Context, db *pgxpool.Pool, playbookID *int64) ([]Rule, error) {
	q := `select id, pattern, coalesce(llm_check,false), coalesce(severity,'Medium'), coalesce(description,'') from rule`
	args := []any{}
	if playbookID != nil {
		q += ` where playbook_id=$1`
		args = append(args, *playbookID)
	}
	rows, err := db.Query(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Rule
	for rows.Next() {
		var r Rule
		if err := rows.Scan(&r.ID, &r.Pattern, &r.LLMCheck, &r.Severity, &r.Guidance); err != nil {
			return nil, err
		}
		out = append(out, r)
	}
	return out, rows.Err()
}

func insertFinding(ctx context.Context, db *pgxpool.Pool, ruleID, clauseID int64, risk, rationale, suggestion string) error {
	_, err := db.Exec(ctx, `
	  insert into finding (rule_id, clause_id, risk, rationale, suggestion)
	  values ($1,$2,$3,$4,$5)
	`, ruleID, clauseID, strings.Title(strings.ToLower(risk)), rationale, suggestion)
	return err
}

func regexMatch(pat, text string) (bool, error) {
	re, err := regexp.Compile(pat)
	if err != nil {
		return false, err
	}
	return re.MatchString(text), nil
}
