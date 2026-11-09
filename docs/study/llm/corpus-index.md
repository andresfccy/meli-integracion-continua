# Semantic Corpus Index

## Topic routing

| Question topic | Read first | Then |
| --- | --- | --- |
| Reconstruction confidence | `01-reconstruction-context.md` | `09-documentation-code-audit.md` |
| Source authority | `02-source-and-documentation-map.md` | normative specs/ADRs |
| Architecture | `03-architecture-explained.md` | `11-senior-architecture-defense.md` |
| Backend | `04-backend-deep-dive.md` | source files |
| Frontend | `05-frontend-deep-dive.md` | source files |
| Request flow | `06-end-to-end-flows.md` | API contract |
| Technology choice | `07-technology-decisions.md` | ADRs/package metadata |
| Requirement mapping | `08-requirements-traceability.md` | specs/test plan |
| Gaps and risks | `09-documentation-code-audit.md` | code |
| Tests | `10-testing-and-quality.md` | test plan/test files |
| Interview defense | `11-senior-architecture-defense.md` | `13-question-bank.md` |
| Presentation | `12-presentation-guide.md` | diagrams/flows |
| Terminology | `14-glossary.md` | specialist document |

## Suggested prompts

### Study

- “Teach me the dependency direction and ask me five questions.”
- “Trace the buyer approval requirement from spec to code and identify missing tests.”
- “Compare coarse question status with lifecycle status.”
- “Explain how a click in React reaches the memory store and returns.”

### Architecture interview

- “Act as a senior architecture interviewer. Ask one question at a time and score answers using the evidence labels.”
- “Challenge the choice of a monolith and provide a model answer with trade-offs.”
- “Ask why REST and GraphQL coexist, then challenge whether GraphQL remains only a read BFF.”
- “Ask what must change before production.”

### Audit

- “List every documented-completed behavior not observable in code.”
- “Separate security flaws from intentional MVP constraints.”
- “Compare required and visible test coverage.”
- “Which claims about the original application are not verifiable?”

### Presentation

- “Create a 15-minute speaking script using only documented or observed claims.”
- “Generate slide notes for architecture, lifecycle, priority, and evolution.”
- “Produce likely follow-up questions and concise answers.”

## Retrieval warnings

- “GraphQL” retrieves both the documented boundary and the regenerated schema-backed implementation; do not claim the exact package/version belonged to the inaccessible original.
- “Scheduler” is implemented and tested locally; do not confuse local verification with proof of original-cloud parity.
- “Completed” refers to plan status, not verification.
- “JWT” is validated for header, signature, claims, and seeded subject, but seed passwords/secrets remain non-production.
- “Tests” include a large required plan plus verified backend, frontend-helper, build, and Playwright evidence; the plan is still broader than exhaustive observed coverage.
