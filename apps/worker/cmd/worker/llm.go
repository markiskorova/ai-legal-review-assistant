package main

import (
	"os"
	"strings"
)

func callLLM(clause, guidance string) (rationale, suggestion, risk string) {
	_ = os.Getenv("OPENAI_API_KEY")
	// MVP stub; replace with actual OpenAI client later
	if strings.Contains(strings.ToLower(clause), "liabil") {
		return "Clause deviates from 12-month fee cap.", "Limit total liability to fees paid in the prior 12 months.", "High"
	}
	return "Looks acceptable per guidance.", "No change required.", "Low"
}
