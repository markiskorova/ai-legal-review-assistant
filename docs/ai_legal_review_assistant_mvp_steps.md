# 🚀 Steps to Build the MVP

## 1. Repo & Infrastructure Setup
**Goal:** Have a working scaffold with containers, DB, migrations, and Makefile commands.  
✅ Checklist:
- [ ] Create GitHub repo (`ai-legal-review-assistant`)  
- [ ] Add `apps/api`, `apps/worker`, `apps/web`, `pkg/*` structure  
- [ ] Write `docker-compose.yml` with: api, worker, postgres (+pgvector), redis, web  
- [ ] Add `Makefile` targets (`make up`, `make db`, `make migrate`, `make seed`, etc.)  
- [ ] Add migration tool (`goose` or `sql-migrate`) and first migrations (users, orgs, matters, documents)  
- [ ] Confirm DB + migrations run inside Docker  

---

## 2. Authentication & Orgs/Users
**Goal:** Secure foundation with JWT.  
✅ Checklist:
- [ ] Implement `POST /api/auth/signup` & `POST /api/auth/login`  
- [ ] Add JWT middleware for protected routes  
- [ ] Seed demo users + orgs  

---

## 3. Documents & Clause Extraction
**Goal:** Upload text contracts and split into clauses.  
✅ Checklist:
- [ ] `POST /api/documents` (accept text for now)  
- [ ] Store in Postgres (`document` + `clause` tables)  
- [ ] Implement simple clause splitter (e.g. split on `\n\n` or punctuation)  
- [ ] Add `GET /api/documents/:id` for debugging  

---

## 4. Playbook Engine (Rules First)
**Goal:** Rules engine without LLM.  
✅ Checklist:
- [ ] Define `playbooks` and `rules` tables  
- [ ] Seed 1 playbook with 6–8 rules (liability, indemnity, governing law, etc.)  
- [ ] Implement `RunReview` pipeline (regex patterns → findings)  
- [ ] API endpoints:
  - `POST /api/review/:document_id/start?playbook_id=...`  
  - `GET /api/review/:document_id/findings`  

---

## 5. Async Worker & LLM Integration
**Goal:** Add LLM checks & suggestions.  
✅ Checklist:
- [ ] Add Redis + worker service  
- [ ] Stub LLM client with OpenAI-compatible call  
- [ ] Extend rules engine: regex hit → enqueue job → worker validates w/ LLM → update finding  
- [ ] Add rationale + suggestion text to findings  

---

## 6. Matter Intake & Linking
**Goal:** Manage contract review lifecycle.  
✅ Checklist:
- [ ] `POST /api/matters`, `GET /api/matters`  
- [ ] Link documents to matters (`document.matter_id`)  
- [ ] Matter detail view shows contracts + review sessions  

---

## 7. Knowledge & Search (Lite)
**Goal:** Simple precedent store + hybrid search.  
✅ Checklist:
- [ ] Add `precedent` + embeddings table  
- [ ] Seed 10 sample precedents  
- [ ] Implement hybrid keyword + vector search in Go  
- [ ] API: `GET /api/search?q=...`  

---

## 8. Frontend (React + TS)
**Goal:** Demo-ready UI for upload, findings, playbooks.  
✅ Checklist:
- [ ] Scaffold React app (Vite + TS + shadcn/ui + Tailwind)  
- [ ] Auth pages (login/signup)  
- [ ] Matters list + detail pages  
- [ ] Upload & review screen (file drop → findings board)  
- [ ] `FindingCard` component with Accept/Reject buttons  
- [ ] Export Findings Report (Markdown/CSV)  

---

## 9. Seeds & Demo Data
**Goal:** Have realistic demo contracts and rules.  
✅ Checklist:
- [ ] Add 3 demo contracts (MSA, NDA, DPA excerpt)  
- [ ] Seed 6–8 playbook rules with regex + guidance  
- [ ] Add precedent snippets with embeddings  

---

## 10. Wrap-Up & Demo Readiness
**Goal:** Make it production-looking enough to show.  
✅ Checklist:
- [ ] Add README with build/run steps, screenshots/GIF  
- [ ] ARCHITECTURE.md with diagram  
- [ ] Run through demo script (create matter → upload → run review → accept redline → export report)  
- [ ] Record 3–5 min screencast for LinkedIn/demo  

---

⚡ Recommendation: Work backend-first (Steps 1–5 on Day 1), then frontend + demo flow (Steps 6–10 on Day 2).  
