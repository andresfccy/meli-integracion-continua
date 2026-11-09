# Documentation versus Code Audit

## Scope

This audit was refreshed after aligning the local implementation on 2026-07-30. It compares active specs, decisions, architecture, declared tasks, source, and executed tests. It does not prove parity with the inaccessible original cloud application.

## Findings

### A1 — Hexagonal monolith

- **[DOCUMENTADO]** ADR-001 requires inward dependencies.
- **[OBSERVADO]** domain, application ports/use cases, inbound HTTP/scheduler adapters, outbound memory/notification adapters, bootstrap, and command exist.
- **Assessment:** aligned. The single application service remains a growth concern, not a contract failure.

### A2 — Scheduler and windows

- **[DOCUMENTADO]** daily 00:00 UTC review, one execution per window, cancellation/recovery, shared use case, and bootstrap wiring.
- **[OBSERVADO]** scheduler calculates next UTC midnight, waits with an injected timer, stops on context cancellation, recovers at the boundary, and calls `ReviewScheduled`.
- **[OBSERVADO]** memory atomically acquires a window; duplicate scheduled calls return `already_executed`; manual review bypasses the lock and adds an override suffix.
- **[VERIFICADO]** scheduler, application, memory-concurrency, and race-detector tests pass.
- **Assessment:** aligned for the in-memory MVP.

### A3 — GraphQL read composition

- **[DOCUMENTADO]** authenticated, role-shaped, read-only GraphQL queries; mutations rejected.
- **[OBSERVADO]** known read operations, roles/ownership, nested filters, mutation/non-read rejection, and error responses exist.
- **[OBSERVADO]** an embedded schema and `graph-gophers/graphql-go` engine provide parsing, field selection, variables, aliases, fragments, introspection, and standard validation.
- **[OBSERVADO]** typed resolvers authenticate and authorize role-shaped reads; absence of a `Mutation` root rejects writes structurally.
- **[VERIFICADO]** HTTP integration tests cover variables, aliases, fragments, unknown-field validation, wrong-role errors, and mutation rejection.
- **[NO VERIFICABLE]** the exact library/version used in the inaccessible original is unknown.
- **Assessment:** aligned with the documented read-only GraphQL composition boundary.

### A4 — Authentication and authorization

- **[OBSERVADO]** token header is valid JSON and algorithm/type, HMAC signature, issuer, audience, expiration, and subject are verified.
- **[OBSERVADO]** Seller, Buyer, and Operations REST/composed routes enforce role and matching identity.
- **[VERIFICADO]** application and HTTP tests cover issued tokens, missing token, wrong role, and protected Operations/Seller paths.
- **Assessment:** aligned for seeded MVP scope.
- **[RECOMENDACIÓN]** production still requires hashed passwords, external secrets, standard identity/session tooling, key rotation, and broader security tests.

### A5 — Question lifecycle and duplicate policy

- **[OBSERVADO]** create → `question_active/open`; reply → `seller_answered/open`; approval → `closed/resolved`.
- **[OBSERVADO]** backend blocks coarse-open duplicates; UI shows a new form after historical questions close.
- **[VERIFICADO]** application and in-process HTTP tests execute the complete Buyer→Seller→Buyer flow.
- **Assessment:** aligned.

### A6 — Filters and read models

- **[OBSERVADO]** application validates supported statuses and date ranges.
- **[OBSERVADO]** Seller, Buyer, and Operations screens forward approved status/text/date/priority filters.
- **[OBSERVADO]** order read models include Seller and Buyer summaries; Operations includes Seller ID.
- **Assessment:** aligned for approved MVP dimensions.

### A7 — Priority and notifications

- **[OBSERVADO]** documented scoring, caps, classifications, dictionary evidence, and deterministic queue ordering exist.
- **[OBSERVADO]** scheduled grouping is duplicate-safe; manual grouping is auditable/resendable; disabled channels skip; unsupported enabled channels error; email console payloads are emitted.
- **[VERIFICADO]** domain, application, memory, and driver tests cover these boundaries.
- **Assessment:** aligned for MVP.
- **[RECOMENDACIÓN]** durable outbox, delivery status, retries, and provider integration remain production evolution.

### A8 — Frontend structure and types

- **[OBSERVADO]** principal role components, session, and GraphQL client are formatted and use typed results/unknown variables.
- **[OBSERVADO]** runtime session validation and robust empty-response parsing were added.
- **[OBSERVADO]** Seller, Buyer, and Operations now render documented question references, timestamps, priority evidence, and mutation success states without forced page reloads.
- **[VERIFICADO]** browser tests exercise role screens and their principal mutations.
- **[BRECHA]** Login/Layout and stylesheet remain compact; isolated component coverage is still narrow; some generic API results still default to `any`.
- **Assessment:** aligned for the visible MVP flow, with maintainability/test-depth debt.

### A9 — Test evidence

- **[VERIFICADO]** `go test -race ./...` passes across domain, application, HTTP, scheduler, memory, and notifications.
- **[VERIFICADO]** `go vet ./...` passes.
- **[VERIFICADO]** Vitest passes 5 session/routing assertions.
- **[VERIFICADO]** TypeScript and Vite production build pass.
- **[VERIFICADO]** Playwright passes three browser scenarios: invalid/valid seeded login, Buyer→Seller→Buyer lifecycle, and Operations seller filtering/manual review.
- **[BRECHA]** isolated React component coverage remains narrow.
- **Assessment:** central backend contracts and principal browser flows now have executable evidence.

### A10 — Remaining technical debt

- `float64` money instead of integer minor units/decimal.
- `latest` dependency ranges despite a lockfile.
- source seed passwords/signing secret.
- in-memory persistence and event delivery.
- no Git metadata in this snapshot.

## Current severity

| Priority | Remaining finding |
| --- | --- |
| Medium | Seed security and in-memory infrastructure are not production-grade |
| Low/Medium | isolated component coverage, money representation, dependency ranges, compact legacy frontend files |

## Recommended next work

1. Externalize/hash credentials and signing configuration.
2. Introduce durable persistence and notification outbox when production scope is approved.
3. Add deeper isolated React component and security-error coverage.
4. Revisit exact money representation if the API contract is approved to change; pin dependency ranges.
5. Finish formatting compact legacy frontend files.
