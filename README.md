# 🧠 AI Legal Review Assistant

An open-source, **Go + React** prototype that simulates a LegalOn-style AI contract-review workflow — combining **playbook-driven rule checks**, **LLM validation**, and **precedent search** inside a lightweight, containerized architecture.

---

## 🧾 Summary

The AI Legal Review Assistant mirrors the core ideas behind LegalOn Technologies’ contract-review platform:
- Automated clause detection and risk analysis.
- Playbook-driven compliance checks (regex + LLM).
- Matter and document management.
- Search across precedents and stored clauses using embeddings (`pgvector`).

It’s designed as a **production-looking MVP** demonstrating hybrid AI + deterministic logic, human-in-the-loop review, and modular backend-worker separation.

---

## ⚙️ Functionality Overview

### 1. Authentication
- JWT-based signup/login endpoints.  
- Multi-organization schema (`org`, `user`).

### 2. Matter Management
- Create and list legal “matters” that group uploaded documents.

### 3. Document Review
- Upload plain text or (later) DOCX/PDF files.
- Split text into clauses → apply regex rules.
- For flagged clauses, the worker calls an **LLM validator** to explain risk and suggest a redline.
- Findings are stored and can be accepted/rejected via API.

### 4. Playbooks & Rules
- Each playbook defines legal policy rules (pattern + guidance + optional LLM check).
- Example: _Cap liability at 12 months fees_, _Mutual indemnity_, _Governing law_.

### 5. Precedent Search
- Hybrid keyword + embedding search (via `pgvector`) across stored precedent snippets.

### 6. Reports
- Export review findings as Markdown or CSV for sharing or audit.

---

## 🏗️ Architecture

```
apps/
  api/            → REST API (Go + Fiber/Gin)
  worker/         → Async LLM & embedding jobs
pkg/              → Shared packages: review, search
infra/            → Docker Compose, Makefile, .env templates
```

### Key Components
| Layer | Responsibility |
|-------|----------------|
| **API** | Auth, matter/document endpoints, review orchestration |
| **Worker** | Consumes Redis queue jobs, runs LLM validations, stores findings |
| **Database** | PostgreSQL + pgvector for structured + semantic data |
| **Queue** | Redis for async job dispatch |
| **LLM Client** | OpenAI-compatible adapter for validation & redlines |
| **Frontend (planned)** | React + TypeScript + shadcn/ui interface for review UX |

---

## 🧩 Tech Stack

| Category | Technology |
|-----------|-------------|
| **Language** | Go 1.22 |
| **Frameworks** | Fiber / Gin (REST), Vite + React + TypeScript (web planned) |
| **Database** | PostgreSQL 15 + pgvector |
| **Async Queue** | Redis 7 |
| **Containers** | Docker + docker-compose |
| **Auth** | JWT |
| **Infra / DevOps** | Makefile, `.env`, GitHub Actions (ready), PowerShell scripts |
| **AI / LLM** | OpenAI API (Validator + Redline prompts) |
| **Search** | Hybrid keyword + vector similarity |
| **Optional** | MinIO (for file blobs), Prometheus metrics |

---

## 📂 Database Schema (simplified)

```
org, user
matter
document, clause, clause_type
playbook, rule
finding, action_log
precedent, precedent_embedding, clause_embedding
```

---

## 🚀 Running Locally

```bash
# 1. Start services
make up

# 2. Run migrations and seed data
make db
make seed

# 3. Access API
http://localhost:8080

# 4. Logs
make logs
```

Docker-Compose services include:
- **db** (Postgres + pgvector)
- **redis**
- **api**
- **worker**

Environment example (`infra/.env.example`):
```env
DATABASE_URL=postgres://postgres:postgres@db:5432/legal_assistant?sslmode=disable
REDIS_URL=redis://redis:6379
JWT_SECRET=supersecret
```

---

## 🧠 Playbook Engine Flow

1. Regex pattern → detect clause match  
2. Optional LLM validator → risk + rationale + suggested redline  
3. Store Finding(record) in DB  
4. Findings accessible via `/review/:document_id/findings`  
5. User accepts/rejects → audit logged  

---

## 📦 Current MVP Status

✅ **Completed / Implemented**
- Repo scaffold (api + worker + pkg + infra)  
- Docker Compose setup  
- Postgres schema and seed SQL (playbook, precedents, org)  
- Basic Auth, Matters, Documents, Review handlers  
- Worker queue + LLM pipeline stub  
- Hybrid search (pgvector + keyword)

🧩 **In Progress / Next**
- Integrate real LLM responses (OpenAI API client)  
- Add React web frontend (login, upload, review UI)  
- Report export (Markdown/CSV)  
- Basic Prometheus metrics & observability  

🚧 **Post-MVP Ideas**
- DOCX/PDF ingestion  
- Role-based permissions  
- Multi-tenant org support  
- Cost logging & rate limits  
- Search UI and Playbook editor  

---

## 🧭 Design Philosophy

- **Explainable AI:** every finding shows its rationale and matched rule.  
- **Human-in-the-loop:** no automatic edits — reviewers decide.  
- **Traceability:** log prompts + responses for audit.  
- **Composable architecture:** modular Go packages (`pkg/review`, `pkg/search`) mirror real SaaS patterns.  

---

## 🧮 Related Projects

This MVP builds on components proven in prior work:

| Project | Tech | Role |
|----------|------|------|
| [AI Ops Assistant](https://github.com/markiskorova/ai-ops-assistant) | Go + GraphQL + OpenAI | Worker pattern & LLM integration |
| [Trend Summary Engine](https://github.com/markiskorova/trend-summary-engine) | Go + GraphQL + Postgres | Auth & schema foundation |
| [AI Contract Simplifier](https://github.com/markiskorova/ai-contract-simplifier) | Node + React + OpenAI | Prompt design and UI ideas |

---

## 📚 References

- **LegalOn Technologies** — for conceptual inspiration (playbooks, legal graph, AI validation).  
- **OpenAI API Docs** — for LLM integration.  
- **pgvector** — vector similarity search extension for Postgres.  

---

## 🧾 License

MIT License © 2025 Marc McAllister  
_This project was built collaboratively with AI assistants as part of the Human–AI Software Engineering workflow described in ["How I Build Software with AI as My Teammate"](https://github.com/markiskorova/articles)._  
