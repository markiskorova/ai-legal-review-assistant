package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	wdb "ai-legal-review-assistant/apps/worker/internal/db"
	"ai-legal-review-assistant/apps/worker/internal/jobs"
	"ai-legal-review-assistant/apps/worker/internal/queue"
)

// import your llm adapter from llm.go
// ensure llm.go provides something like:
// type OpenAILLM struct {}
// func (o *OpenAILLM) Validate(ctx context.Context, clause, guidance string) (risk, rationale, suggestion string, err error) { ... }

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	log.Println("[worker] starting…")

	// DB
	db := wdb.Connect(ctx)
	defer db.Close()

	// Queue
	q := queue.NewClient()

	// LLM (can be nil initially to run rules-only)
	var llm jobs.LLM = new(OpenAILLM) // or set to nil to skip LLM

	// Main loop
	for {
		select {
		case <-ctx.Done():
			log.Println("[worker] shutting down")
			return
		default:
		}

		job, err := q.Dequeue(ctx)
		if err != nil {
			log.Printf("[worker] dequeue error: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}
		if job == nil {
			// no job this tick
			continue
		}

		log.Printf("[worker] processing document_id=%d playbook_id=%v", job.DocumentID, job.PlaybookID)
		if err := jobs.ReviewDocument(ctx, db, llm, job.DocumentID, job.PlaybookID); err != nil {
			log.Printf("[worker] review error: %v", err)
		}
	}
}
