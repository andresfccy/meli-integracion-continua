# Technology Decisions and Defensible Rationale

## Decision record format

The “historical basis” column separates explicit records from reconstruction reasoning.

| Choice | Historical basis | Defensible rationale | Cost / change signal |
| --- | --- | --- | --- |
| Go 1.22 | Decision log and package metadata | Static types, standard HTTP/JSON/testing, single binary, small interfaces | Split packages/add dependencies only for measured complexity |
| Hexagonal monolith | ADR-001 | One deployment with internal testable boundaries | Split services only for independent scaling/ownership/reliability |
| Standard HTTP library | Decision log | Low dependency count and transparent behavior | Adopt router when manual dispatch becomes error-prone |
| In-memory store | ADR-001/003, specs | Deterministic MVP without infrastructure | Replace when durability, transactions, multi-instance operation are required |
| REST commands | ADR-002 | Explicit intent, status codes, clear mutation ownership | Revisit only with new contract/ADR |
| GraphQL reads | ADR-002 and TASK-003B | Schema-backed, role-shaped UI composition without moving rules or commands into GraphQL | Add caching/complexity policy only if measured use requires it |
| React | Architecture/package metadata; no comparative ADR | State-driven UI and feature components | Reconsider only with organizational/product evidence |
| TypeScript | Architecture/package metadata | Compile-time contracts and safer refactoring | Add runtime validation; types do not validate JSON |
| Vite | Architecture/package metadata | Small, fast React/TS development/build setup | Change for SSR or platform requirements |
| pnpm/Corepack | Decision log/package metadata | Reproducible package-manager selection and efficient store | Pin dependencies; preserve platform scripts |
| sessionStorage | ADR-003 | Session-scoped demo auth with no durable session | Replace according to production threat model |
| short JWT | ADR-003, API contract | Authenticate requests without durable session store | Use standards library, external secret, refresh/session design if required |
| notification driver | ADR-005 | Provider independence and testable boundary | Add durable outbox/retries for production |
| Vitest/jsdom/Playwright/Go testing | test plan/tooling | Multiple distances: unit, component, browser, static analysis | Increase actual coverage before relying on claims |

## Go and the monolith

**[DOCUMENTADO]** Go 1.22 and a lightweight hexagonal monolith are explicit. ADR-001 rejects layerless handlers because rules become difficult to trace/test and defers microservices/gRPC as unnecessary complexity.

**[DEFENSA]** This balances simplicity and separation: deployment stays monolithic while code boundaries preserve future adapters. Network boundaries would add failure modes, versioned contracts, distributed tracing, and operations without a demonstrated scaling need.

## REST and GraphQL

**[DOCUMENTADO]** ADR-002 reserves commands for REST and authenticated role-shaped reads for GraphQL. Full GraphQL mutations were rejected because they blur command ownership.

**[DEFENSA]** “Approve answer” is an intentional command with invariants, not a generic field update. GraphQL is valuable when a screen needs a composed read model. Both transports should invoke the same application layer.

**[OBSERVADO]** The local endpoint now uses `graph-gophers/graphql-go` v1.6.0 and an embedded schema. It supports field selection, variables, aliases, fragments, validation, introspection, and typed resolvers. With no `Mutation` root, it enforces the documented read-only boundary structurally.

**[INFERIDO]** The dependency family and exact version are reconstruction choices because the recovered task records the need for a Go GraphQL package, but not its identity. The criterion was a schema-first API, typed Go resolvers, compatibility with Go 1.22, and a small adapter footprint.

**[DEFENSA]** GraphQL acts as a role-oriented read BFF: Seller, Buyer, and Operations screens request only their composed fields. REST owns commands so validation, lifecycle transitions, status codes, and idempotency do not acquire a second mutation route.

## Memory

**[DEFENSA]** Memory is appropriate when the objective is to demonstrate behavior deterministically without database setup. Placing it behind a port prevents the MVP shortcut from becoming a domain dependency.

It is not suitable when restart survival, multi-instance consistency, audit retention, transactions, or production security matter.

## Frontend stack

React, TypeScript, and Vite are documented constraints, but no ADR compares UI frameworks. The defensible rationale is therefore **[INFERIDO]**, not historical:

- React models asynchronous role-shaped screens as stateful components.
- TypeScript supports safe refactoring and internal contracts.
- Vite supplies a small development/build pipeline.
- Feature organization aligns code with user capabilities.

## Authentication

ADR-003 explicitly chooses seeded accounts, sessionStorage, and short JWTs for MVP scope. It excludes registration, recovery, OAuth, refresh tokens, and durable sessions.

**[DEFENSA]** This isolates the demonstration from identity-provider infrastructure. **[RECOMENDACIÓN]** Never present it as production security: passwords and secrets must be externalized and protected, standards-compliant token handling is required, and the browser storage choice needs threat modeling.

## Notification driver

ADR-005 separates event creation from provider delivery. `NotifiableFactory` normalizes the supported MVP recipients—Seller and OpsTeamMember—into a target that owns channel IDs. The review use case follows those IDs rather than deriving a destination from the entity email. This avoids a hidden coupling that would fail as soon as one recipient had multiple channels or a non-email channel.

The domain type is `email`; `Console` is the current driver implementation that simulates sending that email by writing the payload. The dispatcher maps channel types to drivers, handles disabled/unsupported channels, and therefore becomes the only composition point that changes when real email or Slack is introduced.

The evolution path is:

```text
notifiable + logical channel + console simulation
→ durable event/outbox
→ provider driver
→ retries, idempotency, delivery status, observability
```

## Testing tools

Unit tests cheaply prove deterministic rules. HTTP tests prove adaptation/auth/error shape. jsdom proves component behavior without a browser. Playwright proves integrated browser flows. Builds prove type and bundling compatibility. None replaces the others.

**[VERIFICADO]** The current suite combines Go unit/integration/race tests, Vitest helper tests, production build checks, and Playwright role-flow tests. Required scenarios not represented by executable tests must still be treated as targets rather than evidence.
