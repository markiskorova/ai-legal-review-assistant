# AI Legal Review Assistant — pgvector & NLP Overview
**Date:** October 5, 2025

## 🧭 Overview
This discussion clarified how the AI Legal Review Assistant MVP handles language understanding, what pgvector does, and how it connects to your past NLP experience.

---

## 🧩 Key Concepts

### 1. pgvector
- **Purpose:** Stores semantic embeddings (numerical representations of meaning) directly in Postgres.
- **Tables:**
  - `precedent_embedding(precedent_id, vector(1536))`
  - `clause_embedding(clause_id, vector(1536))`
- **Usage:** Enables similarity search between clauses or precedents by meaning, not exact wording.

### 2. Relation to NLP Work
Your past NLP projects (Predict Author, Trend Summary Engine, AI Contract Simplifier, etc.) already used the same conceptual building blocks:
- Tokenization and TF-IDF → early numeric meaning representations.
- Embeddings → modern dense semantic representation.
- Classification, summarization, and search → same architecture now applied to legal text.

**Now:** Instead of classifying or summarizing generic text, you’re embedding legal clauses and comparing their meaning with pgvector.

### 3. Tokenization vs Embedding
| Stage | Purpose | Output | Analogy |
|--------|----------|---------|----------|
| **Tokenization** | Split text into pieces a model can process | Tokens/IDs | Lego bricks |
| **Embedding** | Encode the semantic meaning of those tokens | Numeric vector | The 3D blueprint of meaning |

Tokenization happens inside the LLM; embeddings are what get stored in pgvector.

### 4. External vs Local LLM
- The project **does not host an LLM**.  
- It **connects to external APIs** (e.g. OpenAI) for:
  - Generating embeddings (for pgvector storage)
  - Running playbook validation and suggestions
- Your app orchestrates jobs and stores results; inference happens remotely.

### 5. Core Project Functions
1. **Matter Intake** – track contract reviews.  
2. **Document Upload** – split contracts into clauses.  
3. **Playbook Review** – detect risks using regex + LLM.  
4. **Precedent Search** – semantic retrieval via pgvector.  
5. **Findings Review** – human accept/reject workflow.  
6. **Report Export** – generate Markdown/CSV reports.

### 6. Use Case Example — “Liability Clause Review”
1. Upload a contract.  
2. Worker runs rule-based and LLM checks.  
3. Generates embeddings for each clause.  
4. Stores vectors in pgvector tables.  
5. User searches for similar safe precedents.  
6. pgvector returns the closest matches semantically.  
7. Lawyer accepts a recommended replacement.  
8. System logs and exports a report.

### 7. What pgvector *Actually* Stores
Only embeddings for **your organization’s documents and precedents**.  
It is *not* a global database of legal text — it’s a **private, company-specific knowledge base** that grows as you process contracts.

| Table | Content | Example |
|--------|----------|----------|
| `clause_embedding` | Vectors representing each uploaded clause | `[0.12, -0.48, 0.07, …]` |
| `precedent_embedding` | Vectors for approved precedents | `[0.22, -0.15, 0.05, …]` |

### 8. Summary of Insights
- pgvector = your **semantic memory** layer.  
- LLM API = your **reasoning** and **embedding** generator.  
- Database = your **private knowledge base**.  
- Project = a **legal AI workflow** combining rules, LLMs, and vector search.  

---

**In short:**  
> The AI Legal Review Assistant is a legal AI workflow engine that reviews, explains, and cross-references contract clauses using external LLMs and locally stored semantic embeddings in pgvector — building a private, explainable knowledge base over time.
