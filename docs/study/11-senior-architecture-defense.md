# Senior Architecture Defense

## Opening position

> This is a reconstructed MVP, regenerated from recovered specifications and decision records. I can defend the architecture represented by those sources and the local code, while clearly separating documented decisions from inferred rationale and gaps that cannot be verified against the inaccessible original.

## Why a modular monolith?

**Short answer:** It minimizes operational complexity while preserving internal boundaries.

**Evidence:** ADR-001 explicitly selects a lightweight hexagonal monolith and defers microservices/gRPC.

**Trade-off:** All modules share one process and release. That is acceptable until independent scaling, ownership, reliability, or deployment cadence becomes a measured requirement.

**Reprompt:** “Isn’t it over-engineered for an MVP?”  
**Answer:** The boundaries are lightweight—packages and interfaces, not distributed services. They protect testability without adding network operations.

## Why domain/application/adapters?

Business rules must survive transport and infrastructure changes. Priority calculation belongs in domain; when and for whom it runs belongs in application; HTTP/scheduler translate inbound triggers; memory/providers implement outbound capabilities.

This prevents a new scheduler or database from duplicating ownership and classification rules.

## Why REST commands and GraphQL reads?

REST gives explicit, testable commands: ask, reply, approve, review. GraphQL is intended to compose role-shaped read models. ADR-002 deliberately rejects mutations to keep rule ownership in application.

**Honest caveat:** The reconstruction now uses a real schema-backed GraphQL engine and typed resolvers, but the recovered sources do not prove that the inaccessible original used this exact package or version.

## Why in-memory persistence?

It removes infrastructure from an MVP and makes seed scenarios deterministic. The port boundary preserves an evolution path.

I would not carry it to production because it lacks durability, transactions, multi-instance consistency, and audit retention.

## Why seeded authentication?

ADR-003 deliberately limits identity scope so the project can demonstrate role behavior. It excludes registration, recovery, OAuth, refresh, and external providers.

This is scope control, not a claim of production security. Production would require password hashing, managed secrets, standard token/session libraries, threat modeling, and consistent authorization.

## Why two question states?

The coarse `open/resolved` status preserves existing semantics. `question_active/seller_answered/closed` tells the UI whose action is pending. Order status remains independent because shipping/payment progression is not a customer-support conversation state.

## Why deterministic priority?

A fixed dictionary and scoring formula is explainable, reproducible, cheap, and testable. Storing dictionary version and matched terms provides audit context.

Its limitation is linguistic precision. A future model should still sit behind the same domain policy boundary and expose version/evidence, rather than leaking provider logic through the application.

## Why notification drivers?

Classification decides whether an event should exist; a driver decides how to deliver it. This allows console, email, Slack, or SMS without editing priority rules.

Production evolution needs a durable outbox, idempotency, retries, status tracking, and observability.

## Why feature-oriented frontend?

Each role capability keeps its UI, actions, and local data shape together. This lowers change scattering. Shared transport/session code remains centralized.

Principal feature files are now formatted and typed. Remaining compact legacy files and generic API defaults are localized reconstruction debt.

## What would change first for production?

1. Add durable persistence and migrations.
2. Externalize credentials/secrets and adopt production identity/session tooling.
3. Add durable persistence and an outbox-backed notification delivery path.
4. Expand component-level React coverage beyond session and routing helpers.
5. Externalize production credentials, secrets, and identity configuration.
6. Use exact money representation and complete runtime payload validation.
7. Add observability, configuration, CI/CD, and deployment architecture.

## What debt was acceptable?

Acceptable for demonstration: seeds, memory, console delivery, simple router, no external identity.

Scheduler, window deduplication, authorization, token validation, filters, schema-backed read-only GraphQL, and central backend tests are now present. Remaining disclosed gaps are browser-test coverage and non-production seed/in-memory infrastructure.

## Decision confidence

| Statement | Confidence |
| --- | --- |
| Hexagonal monolith was accepted | High: ADR |
| REST command/GraphQL read split was accepted | High: ADR |
| Memory and seeded auth were MVP choices | High: ADR/specs |
| React was chosen after comparing Vue/Angular | Unknown: no evidence |
| Local code matches original cloud code | Unknown |
| Local lifecycle happy path reflects recovered requirements | High for source inspection |
| Full project works end-to-end | Low until missing verification exists |
