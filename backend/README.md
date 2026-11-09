# Backend

Go backend for the Seller Dashboard MVP. It exposes seeded authentication, seller/buyer/Operations REST APIs, authenticated read-only GraphQL composition, priority review, scheduler wiring, in-memory persistence, and notification-driver delivery.

## Boundaries

The backend follows lightweight hexagonal boundaries:

```text
domain <- application <- adapters <- bootstrap/cmd
```

See [../docs/architecture/backend.md](../docs/architecture/backend.md) for architecture details.

## Main structure

- `cmd/api`: executable entry point; starts HTTP on port 8080.
- `internal/bootstrap`: composition root for store, services, router, scheduler, and notification dispatcher.
- `internal/domain`: commerce, identity, questions, and notifications domain types and policies.
- `internal/application`: auth, orders, questions, Operations, GraphQL read services, and outbound ports.
- `internal/adapters/inbound/http`: REST and GraphQL HTTP handlers plus CORS edge behavior.
- `internal/adapters/inbound/scheduler`: daily priority review scheduler adapter.
- `internal/adapters/outbound/persistence/memory`: deterministic in-memory seed store.
- `internal/adapters/outbound/notifications`: channel dispatcher and console driver.

## API boundaries

- REST owns commands: login, seller replies, backend question creation, buyer question creation, buyer approval, and priority review.
- GraphQL owns authenticated read-only dashboard composition.
- Business rules belong in domain/application layers, not handlers.

## Key references

- API contract: [../specs/api-contract.md](../specs/api-contract.md)
- Data model: [../specs/data-model.md](../specs/data-model.md)
- Priority and notifications: [../specs/SPEC-003-priority-notifications.md](../specs/SPEC-003-priority-notifications.md)
- Implementation status: [../docs/plans/implementation-plan.md](../docs/plans/implementation-plan.md)

## Verification

From this directory:

```bash
go test ./...
go vet ./...
```

Run `gofmt` on changed Go files before committing backend changes.

## Local run

```bash
go run ./cmd/api
```

The server listens on port 8080. Stop any existing process on port 8080 before starting another backend instance.
