# AI Legal Review Assistant (Go + React) — 48‑Hour MVP Plan

## 0) Purpose & Positioning

Build a small, production‑looking MVP that mirrors LegalOn’s core experience and the Senior Full Stack Engineer role’s stack: **Go backend + React frontend**, LLM‑powered **contract review using playbooks**, a **matter intake** mini‑workflow, and a **Legal Document Graph (lite)**. Optimized for a fast, hybrid human–AI workflow to ship in \~48 hours and demo confidently.

---

## 1) Scope (What we’ll ship)

**Core**

1. **Contract Review (Playbook‑Driven)**

   - Upload DOCX/PDF (start with text + .txt to move fast; add DOCX soon).
   - Extract clauses → run **Playbook Rules** (regex+keywords) + **LLM checks**.
   - Show findings: risk level, rationale, suggested redlines.
   - Accept/Reject each suggestion; export a **Findings Report** (Markdown/CSV).

2. **Matter Intake (Mini)**

   - Simple form: title, requester, contract type, due date, status.
   - Link matters ↔ uploaded contracts & review sessions.

3. **Clause/Knowledge Search (Lite)**

   - Full‑text + embedding search across stored clauses.
   - Saved snippets as “precedents.”

4. **Legal Document Graph (Lite)**

   - Postgres schema for **documents, clauses, clause\_types, playbooks, rules, findings, matters, precedents, users, orgs**.

**Non‑Goals (for now)**

- MS Word plugin, deep redline diff in DOCX, SSO, multi‑tenant billing.

---

## 2) High‑Level Architecture

```
apps/
  api/            # Go (REST), Fiber/Gin, Auth (JWT), Postgres (sqlc/GORM)
  worker/         # Go worker for async LLM tasks + embeddings
  web/            # React + TS + Vite, shadcn/ui, Tailwind
pkg/
  review/         # playbook engine, LLM adapters
  graph/          # graph-lite data access (queries, sqlc)
  search/         # embedding + keyword search
  matter/
  auth/
infra/
  docker-compose.yml  # api, worker, db, q, vector (pgvector), minio(optional)
  migrations/         # goose/sqldef/migrate files
  Makefile            # make up, make db, make seed, make test
```

**Datastores**: Postgres (+ pgvector), optional MinIO for file blobs; Redis/RabbitMQ for job queue.\
**LLM**: OpenAI-compatible client (abstracted).\
**Embeddings**: text‑embedding‑3‑large (or local fallback later).\
**Observability**: HTTP logs, pprof, simple Prometheus metrics later.

---

## 3) Data Model (initial)

```sql
-- orgs & users
org(id, name)
user(id, org_id, email, password_hash, role)

-- matters
matter(id, org_id, title, requester, contract_type, due_date, status)

-- documents & clauses
document(id, org_id, matter_id, title, mime, stored_path, text, created_at)
clause(id, document_id, start_idx, end_idx, text, clause_type_id)
clause_type(id, name, description)

-- playbooks & rules
playbook(id, org_id, name, description, active)
rule(id, playbook_id, name, severity, pattern, guidance, llm_check)

-- findings & actions
finding(id, document_id, clause_id, rule_id, severity, rationale, suggestion, status)
action_log(id, finding_id, user_id, action, note, ts)

-- knowledge & search
precedent(id, org_id, title, text, tags)
precedent_embedding(precedent_id, vector)
clause_embedding(clause_id, vector)
```

---

## 4) API Design (REST)

**Auth**

- `POST /api/auth/signup`
- `POST /api/auth/login` → {access\_jwt}

**Matters**

- `POST /api/matters`
- `GET /api/matters?status=…`
- `GET /api/matters/:id`

**Documents & Review**

- `POST /api/documents` (multipart or JSON text)
- `POST /api/review/:document_id/start?playbook_id=…` → job\_id
- `GET /api/review/:document_id/findings`
- `POST /api/findings/:id/accept` | `/reject`
- `GET /api/reports/:document_id` → Markdown/CSV

**Playbooks**

- `GET /api/playbooks`
- `POST /api/playbooks` | `POST /api/playbooks/:id/rules`

**Search**

- `GET /api/search?q=...` (keyword + vector hybrid)

---

## 5) Playbook Engine (Hybrid Rules + LLM)

**Execution order**: (1) Regex/keyword rule screening → (2) LLM validator → (3) Suggestion generator.

- **Rule.pattern**: e.g., detect “limitation of liability” w/ caps like “fees paid in the 12 months prior…”.
- **LLM check**: Prompt to confirm deviation from guidance; return rationale + confidence.
- **Suggestion**: Draft a redline or safer clause with traceable justification.

**LLM Prompts (sketch)**

- *Validator*: “Given the clause and our policy guidance, is this compliant? Explain deviations briefly and rate risk Low/Med/High.”
- *Redline*: “Rewrite to meet the guidance. Keep style consistent. Return JSON: {rationale, suggestion\_text}.”

---

## 6) Frontend UX (React + TS)

Pages: Login • Matters List • Matter Detail • Upload/Review • Findings Board • Playbooks.\
Components: FileDrop, ClauseViewer (side‑by‑side original vs suggestion), FindingCard, RiskBadge, PlaybookRuleEditor.

**Review Screen Flow**

1. Upload text → split into clauses → run checks (job status banner).
2. Findings list (filter by severity/rule).
3. Click finding → show clause, rationale, proposed redline → Accept/Reject.
4. Export report.

Visual polish: shadcn/ui Cards, Tabs for “Overview / Findings / Precedents,” minimal Tailwind tokens.

---

## 7) Reuse From Existing Repos (time‑savers)

- **JWT/Auth + Go project scaffolding** → borrow from *Trend Summary Engine* (graphql-go auth patterns → refit to REST).
- **Worker pattern & OpenAI client** → reuse from *AI Ops Assistant*.
- **LLM Contract pieces** → prompts and UX ideas from *AI Contract Simplifier*.
- **Infra** → Docker Compose, Make targets, GitHub Actions from AI projects to cut setup time.

---

## 8) Infra & DevEx

**Docker Compose services**: api, worker, postgres (w/ pgvector), redis (jobs), web, (minio optional).\
**Makefile**: `make up`, `make db`, `make migrate`, `make seed`, `make test`, `make web`.\
**Migration tool**: goose / sql-migrate.\
**CI**: lint, unit tests, build containers, basic API smoke.

---

## 9) Seeds & Demo Data

- 3 sample contracts (SaaS MSA, DPA excerpt, NDA).
- Playbook: Liability cap, indemnity, governing law, confidentiality, data security (SOC2), assignment.
- 6–8 rules with regex patterns + guidance text (json seeds).
- 10 precedent snippets with embeddings.

---

## 10) Demo Script (5–7 minutes)

1. Create Matter → upload contract.
2. Click “Run Review” → findings populate.
3. Open a High‑risk item → accept the redline.
4. Run search for a precedent → insert as suggestion alternative.
5. Export Findings Report; show Playbook editor briefly.

---

## 11) 48‑Hour Build Plan

**Day 1 (Backend‑first)**

- H1: Repo scaffold, Docker, Postgres + pgvector, migrations.
- H2: Auth + orgs/users; seed script.
- H3: Documents API (text upload), simple clause splitter.
- H4: Playbook + rules models; run **rules engine** (no LLM yet).
- H5: Findings API + acceptance flow; wire Redis jobs; stub LLM client.

**Day 2 (LLM + UI)**

- H1: LLM validator + suggestion prompts; embeddings + hybrid search.
- H2: React app scaffold; auth + Matters List + Upload.
- H3: Findings Board + ClauseViewer; accept/reject actions.
- H4: Report export (Markdown/CSV).
- H5: Seed data polish; run end‑to‑end; record a quick screencast.

**Stretch (if time allows)**: DOCX ingest, role‑based permissions, Prometheus, rate limiters, simple cost logging.

---

## 12) Code Sketches (selected)

**Go: Review pipeline (pseudo)**

```go
func RunReview(ctx context.Context, docID, playbookID int64) error {
  clauses := SplitIntoClauses(loadDocText(docID))
  rules := loadRules(playbookID)
  for _, c := range clauses {
    for _, r := range rules {
      if r.Pattern.MatchString(c.Text) {
        f := Finding{DocID: docID, ClauseID: c.ID, RuleID: r.ID, Severity: r.Severity}
        if r.LLMCheck {
          out := llm.Validate(c.Text, r.Guidance)
          f.Rationale = out.Rationale
          f.Suggestion = out.Suggestion
          f.Severity = out.Risk
        } else {
          f.Rationale = "Pattern matched: " + r.Name
        }
        saveFinding(f)
      }
    }
  }
  return nil
}
```

**React: FindingCard (sketch)**

```tsx
export function FindingCard({ f }: { f: Finding }) {
  return (
    <Card>
      <CardHeader>
        <Badge variant={f.severity.toLowerCase()}>{f.severity}</Badge>
        <h3 className="font-semibold">{f.ruleName}</h3>
      </CardHeader>
      <CardContent>
        <p className="text-sm opacity-80">{f.rationale}</p>
        <Diff original={f.clauseText} suggestion={f.suggestion} />
      </CardContent>
      <CardFooter className="gap-2">
        <Button onClick={() => accept(f.id)}>Accept</Button>
        <Button variant="outline" onClick={() => reject(f.id)}>Reject</Button>
      </CardFooter>
    </Card>
  )
}
```

---

## 13) Trust, Safety, and Auditability (MVP)

- Show **why** each finding was created (rule match + LLM rationale).
- Store prompt/response hashes for audit; log which version of playbook used.
- “Never‑auto‑apply” default; human‑in‑the‑loop.

---

## 14) How This Mirrors LegalOn & The Role

- **Go + React** full‑stack delivery with **agentic LLM checks**.
- **Playbooks** & **Legal Document Graph (lite)** echo core product concepts.
- **Matter intake** reflects adjacent workflow coverage.
- **Explainability UI** and **audit trails** address trust concerns.
- **Global‑team hygiene**: clean code, tests, docs, CI, containerization.

---

## 15) “Show‑Off” Readme Checklist

- Badges (build passing, docker run).
- Screenshots/GIF of review flow.
- One‑page ARCHITECTURE.md diagram.
- Demo dataset + seeds.
- Mapping table: Job bullet → Feature/commit.

---

## 16) Next Steps (You can assign me)

- Create GitHub repo (private for now) with this scaffold.
- I’ll port auth + worker patterns from prior repos; add migrations + seeds.
- You focus Day‑2 on React screens; I’ll wire LLM + search.

---

## 17) Appendix — Prompt Templates (editable)

**Validator**

```
SYSTEM: You are a senior contracts counsel. Apply our policy strictly.
USER: Clause:\n{{clause}}\n\nPolicy Guidance:\n{{guidance}}\n\nReturn JSON: {"risk":"Low|Medium|High","rationale":"…","suggestion":"…"}
```

**Precedent Finder**

```
SYSTEM: Match clauses to our precedent snippets.
USER: Clause:\n{{clause}}\nCandidates: {{k top snippets}}
Return best 3 with short rankings.
```

