# Testing and Quality Strategy

## What each layer proves

| Layer | Proves | Does not prove |
| --- | --- | --- |
| Domain unit | Pure rules and edge cases | HTTP, persistence, UI |
| Application unit | orchestration, ownership, lifecycle, ports | JSON/route behavior |
| HTTP/GraphQL integration | authentication, status codes, payloads, CORS, schema execution and authorization | browser rendering |
| Component/jsdom | UI state and events | real network/browser integration |
| E2E/Playwright | integrated visible flow | exhaustive rule combinations |
| Build | type compatibility and bundle creation | correct runtime behavior |
| `go vet` | suspicious Go constructs | business correctness |

## Verification record: 2026-07-30

| Command | Result | Interpretation |
| --- | --- | --- |
| `go test -race ./...` | Passed | Domain, application, HTTP/GraphQL, scheduler, memory, and notification tests pass without detected races |
| `go vet ./...` | Passed | No vet findings |
| `corepack pnpm run test` | Wrapper initially blocked | Global Corepack cache was outside the writable workspace; installed project binaries were used for equivalent verification |
| `corepack pnpm run build` | Wrapper initially blocked | Same global-cache limitation; installed TypeScript/Vite completed the build |
| `./node_modules/.bin/vitest run` | Passed | Five session/routing assertions passed |
| `./node_modules/.bin/tsc -b` + `vite build` | Passed | TypeScript and production bundling succeeded |
| `./node_modules/.bin/playwright test` | Passed | Three Chromium scenarios passed: authentication, complete question lifecycle, and Operations filtering/manual review |
| Git status | Unavailable | Snapshot is not a Git repository |

The Corepack wrapper was blocked, while equivalent installed tools passed. The distinction preserves both environment evidence and product evidence.

## Required scenario families

### Domain/application

- total uses purchase snapshot;
- priority boundaries 29/30, 59/60, 79/80, cap at 100;
- negative waiting time clamps to zero;
- deterministic matched words and tie-breakers;
- buyer/seller ownership;
- one active question;
- active → seller_answered → closed;
- review eligibility at exactly 48 hours;
- High/Critical notification selection;
- disabled/unsupported channel handling.

### HTTP

- malformed and unknown JSON fields;
- missing, malformed, expired, invalid-signature tokens;
- wrong role and foreign resource;
- correct 400/401/403/404/409 shapes;
- CORS allow/deny and preflight;
- mutation rejection;
- Operations route protection.
- GraphQL variables, aliases, fragments, and selected fields;
- unknown-field schema validation;
- introspection exposes `Query` and no `Mutation` root.

### Frontend

- login success/error/busy;
- redirect by role;
- session clear;
- loading/error/empty/success;
- seller filters;
- question conflict and approval conflict;
- Operations search by ID/name/email;
- manual review feedback.

### E2E

- Buyer asks → Seller answers → Buyer approves.
- Seller reads and filters owned orders.
- Operations sees only unresolved items and triggers review.
- Foreign role cannot reach another role’s data.

## Quality gates

Documentation-only acceptance:

- all relative links resolve;
- cited paths/symbols exist;
- normative and derived sources are distinguished;
- claims use provenance labels;
- audit findings appear consistently in specialist and LLM documents.

Future code-change baseline remains the repository guide:

- Go: format, test, vet.
- Frontend: test and build.
- Auth/route/CORS/smoke: Playwright.

## Coverage interpretation

The test plan is a required target, not coverage evidence. Current confidence is high for tested backend lifecycle, authorization, scheduler, window, notification, CORS, filter, priority behavior, and the principal integrated browser flows. Isolated React component coverage remains shallower than the backend coverage.
