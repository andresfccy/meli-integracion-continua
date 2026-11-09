# ADR-002: REST Commands and GraphQL Reads

## Status

Accepted

## Context

The frontend benefits from role-shaped dashboard read models, while state-changing operations require simple, explicit command boundaries.

## Decision

REST remains the MVP command API for login, backend question creation, seller replies, buyer question creation, buyer approval, and priority review triggers. GraphQL is approved only as an authenticated frontend-facing read composition layer.

GraphQL must not own business rules, expose public contracts, replace REST commands, or perform mutations in the MVP.

## Consequences

- Commands stay explicit and easy to test through HTTP handlers.
- GraphQL reduces frontend read composition without becoming a domain layer.
- Future GraphQL mutations require a new accepted decision and spec updates.

## Alternatives considered

- REST-only reads: acceptable but less convenient for role-shaped dashboard composition.
- Full GraphQL API with mutations: rejected for MVP because it would blur command ownership.

## References

- [../../specs/api-contract.md](../../specs/api-contract.md)
- [../../specs/frontend-contract.md](../../specs/frontend-contract.md)
- [../architecture/backend.md](../architecture/backend.md)
- [../architecture/frontend.md](../architecture/frontend.md)
