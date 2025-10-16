package main

import (
	"context"
)

type OpenAILLM struct{}

func (o *OpenAILLM) Validate(ctx context.Context, clause, guidance string) (risk, rationale, suggestion string, err error) {
	// Stub for now; you can integrate real OpenAI later.
	// Return a deterministic response so the UI has something to show.
	return "High", "Clause deviates from guidance keywords.", "Proposed safer language...", nil
}
