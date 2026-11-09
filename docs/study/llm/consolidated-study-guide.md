# Consolidated Study Guide: Regenerated Seller Dashboard

## 1. Provenance

The original cloud application is inaccessible. This repository is a secondary regeneration from recovered documentation. Requirements and decisions can be cited; exact parity with the original cannot.

Use these labels:

- **[DOCUMENTADO]** recovered source says it.
- **[OBSERVADO]** local code shows it.
- **[VERIFICADO]** a command confirms it.
- **[INFERIDO]** technically reasonable conclusion.
- **[DEFENSA]** presentation rationale.
- **[BRECHA]** definition/evidence mismatch.
- **[NO VERIFICABLE]** original evidence unavailable.

## 2. Requirement-to-code method

```text
specification → contract → ADR → task → code → test → observed result
```

Specs own required behavior. ADRs own accepted direction. Plans declare completion. Code shows local reality. Tests prove only exercised cases.

## 3. Architecture

The documented architecture is a lightweight hexagonal Go monolith:

```text
domain <- application <- adapters <- bootstrap/cmd
```

React calls REST commands and schema-backed GraphQL reads. HTTP adapts requests to application. Application coordinates domain policy with memory and notification ports. Bootstrap chooses concrete implementations.

**Why:** one deployment keeps MVP operations simple; package boundaries keep rules independent and testable. Microservices are deferred until independent scaling, ownership, reliability, or release needs justify distributed cost.

## 4. Domain

Domain types represent identity, commerce, questions, and notifications. `Order.Total` uses purchase-time item prices.

Question has:

- `status`: open/resolved;
- `lifecycleStatus`: question_active/seller_answered/closed.

Seller reply moves active/open to seller_answered/open. Owning buyer approval moves it to closed/resolved. Order status remains separate.

Priority:

```text
critical phrase 80 OR high phrase 60
+ five points/full waiting day capped 30
+ floor(order value/100) capped 20
= score capped 100

Critical >=80
High Priority >=60
Medium >=30
Low otherwise
```

The pure function is deterministic and records matched words. The application stores dictionary version `mvp-2026-07-13`.

## 5. Application

Application defines `Store` and `Delivery` interfaces. It owns:

- seed login and token validation;
- seller/buyer ownership;
- order filters and deterministic sorting;
- question validation/duplicate detection;
- reply/approval lifecycle;
- Operations queue;
- review eligibility;
- notification event creation and delivery delegation.

This is orchestration, not transport. The current single service is acceptable for a small reconstruction but should split when complexity/change ownership demands it.

## 6. Adapters

HTTP owns paths, JSON, bearer extraction, CORS, status codes, and response mapping. Memory owns seed state guarded by a mutex and returns defensive copies of nested mutable data.

Notification targeting is a separate chain. `NotifiableFactory` converts the supported MVP recipients—Seller and OpsTeamMember—into generic targets whose channel identifiers are explicit. Review follows the seller target and those IDs rather than matching email fields. The dispatcher maps the logical channel type to a driver; in the MVP, `email` maps to `Console`, which simulates sending by writing JSON to an injected writer. Thus Console is not the business channel and future real email or Slack registration does not change domain classification or application orchestration.

Bootstrap creates memory, console, application, HTTP handler, and scheduler. Main starts the scheduler and listens on 8080. The scheduler calculates 00:00 UTC, supports cancellation, and uses atomic one-window acquisition.

## 7. Backend technology

Go 1.22 is documented. A defensible rationale is static types, built-in HTTP/JSON/testing, small interfaces, and a single compiled process with few dependencies.

Risks:

- `float64` money;
- one large service;
- raw passwords/secret;
- source seed passwords/secret;
- in-memory durability.

## 8. Frontend

React components are organized by feature: auth, dashboard, Seller, Buyer, Operations, GraphQL. Shared session/API logic adds bearer tokens and stores the session in `sessionStorage`.

TypeScript checks source and the session parser validates its minimal runtime shape. Principal GraphQL results are typed. Vite builds and serves port 3000. pnpm/Corepack define package tooling.

Observed gaps:

- compact legacy files and stylesheet;
- no component/E2E coverage;
- full page reload after commands;
- minimal visible tests.

## 9. REST and GraphQL

ADR-002:

```text
REST = explicit state-changing commands
GraphQL = authenticated role-shaped reads
```

This keeps business rules in application and mutations explicit.

**[OBSERVADO]** Local `/graphql` uses an embedded schema and `graph-gophers/graphql-go` v1.6.0. Typed resolvers compose `viewer`, Seller, Buyer, and Operations reads from application services. Standard variables, aliases, fragments, field selection, validation, and introspection are available. No `Mutation` root exists, so commands remain REST-only.

**[DOCUMENTADO]** ADR-002 establishes the REST-command/GraphQL-read split, and TASK-003B explicitly asks for a Go GraphQL dependency and read-only schema. **[INFERIDO]** The precise library/version is a reconstruction choice because it was not preserved in the recovered records.

## 10. Main flows

Login: React → REST → application accounts → one-hour token → sessionStorage.

Question: Buyer REST command → ownership/validation/duplicate → priority → memory.

Reply: Seller REST command → seller ownership/open/no reply → seller_answered.

Approval: Buyer REST command → buyer ownership/reply/seller_answered → closed/resolved.

Operations: composed queue read → open questions sorted by priority/age/value/ID.

Manual review: Operations-authenticated REST → reclassify after 48 hours → group High/Critical by seller/channel → save event → console delivery.

Scheduled review: daily 00:00 UTC, cancellation-aware, with atomic one-execution-per-window acquisition; composed locally from bootstrap.

## 11. Technology defense

| Choice | Defense | Limitation |
| --- | --- | --- |
| Modular monolith | low operations + internal boundaries | shared process/release |
| Memory | deterministic infrastructure-free MVP | no durability/multi-instance |
| REST commands | explicit intent/invariants/errors | more endpoints |
| GraphQL reads | schema-backed role-shaped UI composition | intentionally no mutations/subscriptions/federation |
| React/TS/Vite | state UI, typed source, small build setup | no runtime validation; framework rationale inferred |
| sessionStorage/JWT | session-scoped demo without durable sessions | not production security |
| Driver boundary | provider independence | no durable delivery/retry |

## 12. Audit summary

Aligned:

- package boundary shape;
- purchase-price total;
- question lifecycle;
- priority policy;
- memory and console adapters;
- manual and scheduled review;
- window deduplication and manual override;
- backend authorization and filters;
- JWT header/signature/claims;
- focused tests across backend boundaries.

Partial/absent:

- isolated React component coverage remains narrow, although principal Playwright flows pass;
- seed credentials, memory, and console delivery remain non-production.

## 13. Verification

On 2026-07-30:

- `go test -race ./...`: passed across domain, application, HTTP/GraphQL, scheduler, memory, and notifications, including schema variables/aliases/fragments/validation/introspection, role enforcement, and mutation rejection.
- `go vet ./...`: passed.
- prescribed Corepack test/build wrapper: blocked because it required a registry download;
- installed Vitest: passed five session/routing assertions;
- installed TypeScript and Vite: production build passed.
- Playwright Chromium: three integrated role-flow scenarios passed.
- Git audit: unavailable because `.git` is absent.

The successful runs verify central backend contracts, frontend build/helper behavior, and the principal browser E2E flows. They do not provide exhaustive isolated component coverage.

## 14. Production evolution

1. Durable database with migrations and exact money representation.
2. Production secrets/password/identity/session handling.
3. Deeper isolated React component and security/error-path tests.
4. Durable notification outbox, retry, status, observability.
5. Complete runtime validation and remaining frontend formatting.
6. CI/CD, configuration, metrics, logging, and deployment design.

## 15. Presentation statement

> The recovered design intentionally uses a modular monolith to keep the MVP operationally simple while protecting business rules behind domain, application, and adapter boundaries. REST owns commands; a schema-backed GraphQL read BFF composes data for each role; persistence and delivery are replaceable. The regenerated code now includes authenticated role flows, filters, lifecycle, priority, manual/scheduled review, window idempotency, notification drivers, standard GraphQL semantics, and focused race-enabled tests. Remaining gaps are browser E2E evidence and production-grade data/security infrastructure; none can establish exact parity with the inaccessible original.

## 16. Core interview questions

1. Why is priority domain logic but review application logic?
2. Why do coarse status and lifecycle coexist?
3. Why is memory behind a port?
4. Why not microservices?
5. Why REST writes and GraphQL reads?
6. What evidence supports each historical claim?
7. What is missing before production?
8. Which “completed” tasks lack executable evidence?

For expanded answers, consult `../11-senior-architecture-defense.md` and `../13-question-bank.md`.
