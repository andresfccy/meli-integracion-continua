# ADR-001: Lightweight Hexagonal Monolith

## Status

Accepted

## Context

The MVP is a full-stack challenge with a Go backend, in-memory persistence, seeded data, REST commands, read-only GraphQL composition, and a Vite React frontend. The backend needs clear boundaries without introducing distributed-system complexity.

## Decision

Use a Go monolith with lightweight hexagonal boundaries:

```text
domain <- application <- adapters <- bootstrap/cmd
```

Domain owns entities and policies. Application owns use cases and ports. Adapters own HTTP, GraphQL, scheduler, persistence, and notification-driver details. Bootstrap composes concrete implementations.

## Consequences

- Business behavior remains testable without transport or provider dependencies.
- Future storage or notification adapters can be introduced behind ports.
- The monolith stays simple enough for the MVP.
- Package boundaries must be maintained during future changes.

## Alternatives considered

- Layerless handlers and storage for speed: rejected because business rules would be harder to trace and test.
- Microservices or gRPC service split: deferred as unnecessary MVP complexity.

## References

- [../architecture/backend.md](../architecture/backend.md)
- [../../backend/README.md](../../backend/README.md)
- [README.md](README.md)
