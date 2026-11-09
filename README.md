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

## Academic Docker delivery

This is a shareable academic case study, not a Mercado Libre production service.
Only two runtime services are used:

```text
Browser -> 127.0.0.1:3000 -> web (Nginx and static React)
                              -> api:8080 (Go, in-memory data)
```

The `application` bridge network provides service-name DNS. API port 8080 is not
published to the host. Nginx forwards `/api/` and `/graphql` without rewriting the
path and forwards the Authorization header. Both runtime processes are non-root.
Base images are pinned by digest and JavaScript dependencies use the frozen lockfile.
The web build opts into `VITE_API_MODE=same-origin`; local Vite keeps port 8080.

Prerequisites: Docker Desktop/Engine with Compose, and a free local port 3000.
From the repository root:

```bash
docker compose config --quiet
docker compose up --build -d --wait
docker compose ps
curl --fail http://127.0.0.1:3000/api/health
docker compose exec -T web wget -qO- http://api:8080/api/health
```

Open http://127.0.0.1:3000 and use the demo accounts above.
Both health requests return `{"status":"ok"}`. The second command executes
inside `web` and reaches `api` by its Compose service name: this proves
container-to-container communication rather than two independent host requests.

Run browser tests against the already running containers, without starting Vite
or a host Go server:

```bash
cd frontend
corepack pnpm install --frozen-lockfile
corepack pnpm exec playwright install chromium
corepack pnpm exec playwright test --config playwright.container.config.ts
```

The tests cover invalid login, seller reads, buyer question creation, seller reply,
buyer approval, and Operations filtering/manual review. For deterministic reruns,
reset the demo data first with `docker compose restart api web` and wait until
both services are healthy. All data is synthetic and in memory; restarting `api`
resets it. This is intentional for this case study, not durable storage.

Stop the demo with `docker compose down`. Do not expose this demonstration to the
public Internet; it intentionally contains well-known seeded demo credentials.
See [verified evidence](evidence/README.md) for the observed run.
