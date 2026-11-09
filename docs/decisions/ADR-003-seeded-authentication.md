# ADR-003: Seeded Authentication for MVP

## Status

Accepted

## Context

The MVP needs role-specific Seller, Buyer, and Operations behavior without introducing production identity infrastructure.

## Decision

Use seeded in-memory accounts with deterministic demo credentials. Authentication validates email and password against seed data. The frontend stores authenticated state only in `sessionStorage`. Short-lived JWT access tokens support authenticated GraphQL reads and protected buyer/Operations commands.

Registration, password recovery, social login, OAuth, refresh tokens, external identity providers, and durable backend sessions are out of scope.

## Consequences

- Demo flows remain deterministic.
- Role-specific UI and API behavior can be verified without external services.
- Authentication must not be represented as production-grade identity.

## Alternatives considered

- No authentication: rejected because role-specific behavior needs authorization context.
- Production identity provider: rejected as outside MVP scope.

## References

- [../../specs/frontend-contract.md](../../specs/frontend-contract.md)
- [../../specs/api-contract.md](../../specs/api-contract.md)
- [../architecture/frontend.md](../architecture/frontend.md)
