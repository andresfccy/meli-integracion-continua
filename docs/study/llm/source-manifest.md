# LLM Source Manifest

## Usage

This manifest tells an LLM which sources define truth and which explain or interpret it. Prefer normative sources when answering contract questions. Prefer code/audit sources when answering implementation questions.

| Source | Class | Authority | Topics | Content type |
| --- | --- | --- | --- | --- |
| `specs/SPEC-001-seller-orders.md` | Normative | High for orders | orders, filters, prices | documented requirements |
| `specs/SPEC-002-buyer-questions.md` | Normative | High for questions | lifecycle, ownership, duplicates | documented requirements |
| `specs/SPEC-003-priority-notifications.md` | Normative | High for priority | scoring, scheduler, notifications | documented requirements |
| `specs/data-model.md` | Normative | High for data | entities, invariants, grouping | documented contract |
| `specs/api-contract.md` | Normative | High for transport | REST, GraphQL, errors, auth | documented contract |
| `specs/frontend-contract.md` | Normative | High for UI | routes, roles, states | documented contract |
| `specs/test-plan.md` | Normative target | High for required tests | acceptance, coverage | required evidence, not actual result |
| `docs/decisions/*.md` | Decision | High for accepted direction | architecture, API, auth, lifecycle, notifications | historical documented decisions |
| `docs/architecture/*.md` | Architecture | Medium/high | boundaries and evolution | documented direction |
| `docs/plans/implementation-plan.md` | Status claim | Medium | tasks 001–018 | declared completion |
| `backend/**`, `frontend/src/**`, package metadata | Implementation | High for local reality | all runtime behavior, embedded GraphQL schema and dependency | observed regenerated code |
| `docs/study/01-*` | Derived | Context | provenance | facts and limits |
| `docs/study/02-*` | Derived | Navigation | source authority | synthesis |
| `docs/study/03-*` | Derived | Explanation | architecture | documented + observed + inferred |
| `docs/study/04-*`, `05-*` | Derived | Explanation | backend/frontend | observed + recommendations |
| `docs/study/06-*` | Derived | Explanation | sequences | observed + gaps |
| `docs/study/07-*` | Derived | Defense | technology choices | documented + inferred |
| `docs/study/08-*` | Derived | Traceability | requirement/code/test | synthesis |
| `docs/study/09-*` | Derived | Audit | mismatches | observed comparison |
| `docs/study/10-*` | Derived | Verification | tests/results | verified + required |
| `docs/study/11-*` | Derived | Defense | interview rationale | documented + defense |
| `docs/study/12-*`, `13-*` | Derived | Presentation | speaking/Q&A | defense |
| `docs/study/14-*` | Derived | Reference | terminology | explanation |

## Snapshot metadata

- Analysis date: 2026-07-30
- Repository path used during analysis: `/Users/acardenas/Workspace/MeLi`
- Git metadata: absent
- Original cloud application: inaccessible
- Backend verification: race-enabled tests and vet pass across domain, application, HTTP/GraphQL, scheduler, memory, and notifications
- GraphQL implementation: `graph-gophers/graphql-go` v1.6.0; schema-backed read BFF with no mutation root. The boundary is documented; the exact package/version is reconstructed.
- Frontend verification: installed Vitest passed 5 session/routing assertions, TypeScript/Vite built successfully, and Playwright passed 3 Chromium role-flow scenarios

## Answering policy for an LLM

1. State whether an answer is documented, observed, verified, inferred, or recommended.
2. Do not infer original-cloud parity.
3. Do not interpret task checkboxes as test evidence.
4. Mention audit gaps when they materially qualify an answer.
5. When asked “why,” distinguish recorded rationale from a defensible rationale.
