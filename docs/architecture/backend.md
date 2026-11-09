# Backend Architecture

## Current implementation

The backend is a Go 1.22 monolith under [../../backend](../../backend). It uses lightweight hexagonal boundaries with dependencies pointing inward:

```text
domain <- application <- adapters <- bootstrap/cmd
```

The executable entry point is [../../backend/cmd/api/main.go](../../backend/cmd/api/main.go). Bootstrap wiring lives in [../../backend/internal/bootstrap](../../backend/internal/bootstrap). HTTP routes are registered in [../../backend/internal/adapters/inbound/http/router.go](../../backend/internal/adapters/inbound/http/router.go).

## Layers

### Domain

[../../backend/internal/domain](../../backend/internal/domain) owns business entities and deterministic policies:

- Commerce: buyers, sellers, products, orders, order items, and order status.
- Identity: seeded accounts, roles, and Operations team members.
- Questions: question lifecycle, replies, priority metadata, and `EvaluatePriority`.
- Notifications: notification channels, notifiable targets, and notification events.

Domain code must not depend on HTTP, GraphQL, stdout, schedulers, persistence adapters, or provider protocols.

### Application

[../../backend/internal/application](../../backend/internal/application) owns use cases:

- Seeded authentication and access-token validation.
- Seller and buyer order reads.
- Seller replies, buyer question creation, and buyer approval.
- Operations priority queue reads and manual/scheduled priority review.
- Read-only GraphQL composition.

Application code depends on outbound ports, not concrete adapters.

### Adapters

[../../backend/internal/adapters](../../backend/internal/adapters) contains inbound and outbound adapters:

- HTTP handlers map REST and GraphQL requests to application services.
- The scheduler adapter triggers the priority review use case.
- In-memory persistence implements current storage ports.
- Notification drivers deliver channel payloads at the boundary.

### Bootstrap and command

[../../backend/internal/bootstrap](../../backend/internal/bootstrap) composes the store, services, notification dispatcher, scheduler, and router. [../../backend/cmd/api/main.go](../../backend/cmd/api/main.go) starts the server on port 8080 and starts configured background processes.

## API boundaries

- REST is the command boundary for login, seller replies, buyer question creation, buyer approval, and priority review commands.
- GraphQL is authenticated, read-only dashboard composition.
- GraphQL must not own business rules or mutations for the MVP.

## Persistence

The MVP uses in-memory persistence. Durable storage is a future option and should be added behind existing ports or new ports justified by specs.

## Verification

From [../../backend](../../backend):

```bash
go test ./...
go vet ./...
```

Run `gofmt` on changed Go files when backend code changes.
