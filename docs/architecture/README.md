# Architecture Overview

This directory describes current architecture, approved direction, and future options. It does not redefine product requirements; use [../../specs/README.md](../../specs/README.md) for contracts.

## Current implementation

The system is a full-stack monorepo:

- Backend: Go 1.22 monolith on port 8080.
- Frontend: Vite React TypeScript app on port 3000.
- Persistence: in-memory MVP store.
- Commands: REST.
- Reads: REST plus authenticated read-only GraphQL dashboard composition.
- Background work: inbound daily priority review scheduler.
- Notifications: seller notification events delivered through outbound channel drivers, with an MVP console email driver.

## Backend dependency direction

```text
domain <- application <- adapters <- bootstrap/cmd
```

Domain and application layers must not depend on HTTP, GraphQL, stdout, provider protocols, scheduler implementation, or concrete persistence.

## Architecture documents

- [backend.md](backend.md): backend layers, routes, ports, and verification.
- [frontend.md](frontend.md): frontend runtime, routes, feature boundaries, and verification.
- [notifications.md](notifications.md): priority review and notification delivery boundaries.
- [evolution.md](evolution.md): current implementation, approved direction, and future options.

## Approved direction

- Keep architecture intentionally simple.
- Keep REST as the command boundary.
- Keep GraphQL read-only and frontend-facing for role-shaped reads.
- Keep notification providers behind outbound channel drivers.
- Keep scheduled review as an inbound adapter invoking the Operations use case.
- Keep future options separate from current behavior.

## Future options

Future options require specs, decisions, implementation-plan updates, and tests before implementation. See [evolution.md](evolution.md).
