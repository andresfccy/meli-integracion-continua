# Seller Dashboard: Orders and Questions MVP

Full-stack MVP for a seller dashboard with order visibility, buyer questions, buyer approval of seller answers, Operations priority review, seeded authentication, read-only GraphQL dashboard composition, and driver-backed seller notification delivery.

## Status

The implemented MVP includes:

- Go backend on port 8080.
- Vite React TypeScript frontend on port 3000.
- Seeded Seller, Buyer, and Operations accounts.
- REST commands for login, seller replies, buyer question creation, buyer answer approval, and Operations priority review.
- Authenticated read-only GraphQL for role-shaped dashboard reads.
- In-memory persistence and deterministic seed data.
- Daily scheduled priority review plus Operations manual review.
- Console-driver notification delivery for seller notification events.

Buyer notification delivery, durable persistence, external notification providers, GraphQL mutations, OAuth, refresh tokens, webhooks, gRPC, event streaming, and Infrastructure as Code remain outside the MVP unless later approved.

## Stack

- Backend: Go 1.22.
- Frontend: Vite, React, TypeScript, pnpm through Corepack.
- Tests: Go tests/vet, Vitest, Playwright smoke tests.

## Documentation

- Documentation index: [docs/README.md](docs/README.md)
- Agent operating guide: [AGENTS.md](AGENTS.md)
- Specs and contracts: [specs/README.md](specs/README.md)
- Architecture: [docs/architecture/README.md](docs/architecture/README.md)
- Decisions: [docs/decisions/README.md](docs/decisions/README.md)
- Implementation status: [docs/plans/implementation-plan.md](docs/plans/implementation-plan.md)
- Backend map: [backend/README.md](backend/README.md)
- Frontend map: [frontend/README.md](frontend/README.md)

## Demo credentials

- Seller: `seller@example.com` / `seller-demo-password`
- Operations: `ops@example.com` / `ops-demo-password`
- Buyer: `buyer@example.com` / `buyer-demo-password`

Additional deterministic seeded users are listed on the login page.

## Local setup

From [frontend](frontend):

```bash
corepack pnpm install
corepack pnpm run start
```

The `start` script runs the Go backend and Vite frontend concurrently. The frontend listens on `0.0.0.0:3000`; the backend listens on `:8080`.

If port 8080 is already in use, stop the existing backend process before running the combined start command.

## Verification

Backend from [backend](backend):

```bash
go test ./...
go vet ./...
```

Frontend from [frontend](frontend):

```bash
corepack pnpm run test
corepack pnpm run build
```

E2E smoke tests from [frontend](frontend), when relevant:

```bash
corepack pnpm run install-e2e
corepack pnpm run test-e2e
```

See [docs/development/TESTING_GUIDELINES.md](docs/development/TESTING_GUIDELINES.md) for task-specific verification guidance.
