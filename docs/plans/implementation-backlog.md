# Implementation Backlog

## Phase 1: Bootstrap Specs And Tooling

1. Objective: establish source-of-truth specs and pnpm package metadata.
2. Files affected: `frontend/package.json`, `frontend/pnpm-lock.yaml`, root docs, `specs/`, `docs/architecture/`.
3. Acceptance criteria: pnpm lockfile exists, npm lockfile is removed, specs and operating docs exist.
4. Verification commands: `corepack pnpm install`, `corepack pnpm run test`, `go test ./...`.
5. Dependencies: none.

## Phase 2: Backend Domain And Seed Storage

1. Objective: define domain types and seed in-memory data approved by specs.
2. Files affected: backend domain and storage packages.
3. Acceptance criteria: seeded sellers, products, buyers, orders, items, and questions are available to services.
4. Verification commands: `gofmt`, `go test ./...`, `go vet ./...`.
5. Dependencies: approved domain model spec.

## Phase 3: Backend API

1. Objective: expose approved REST order, backend question creation, reply, operations, authentication, seeded-role endpoints, and the GraphQL read HTTP endpoint.
2. Files affected: backend handlers, services, main entrypoint.
3. Acceptance criteria: REST endpoints and the GraphQL read endpoint return documented JSON and error behavior from `specs/api-contract.md`.
4. Verification commands: `gofmt`, `go test ./...`, `go vet ./...`.
5. Dependencies: Phase 2 and approved API contract.

## Phase 4: Priority And Notifications

1. Objective: implement approved priority policy, 2-day recategorization, notifiable channels, and batched notification abstraction.
2. Files affected: backend services and notifications packages.
3. Acceptance criteria: unresolved questions are sorted by policy, High Priority/Critical classifications trigger notification execution, and qualifying messages are batched by notifiable target, channel, and review date.
4. Verification commands: `gofmt`, `go test ./...`, `go vet ./...`.
5. Dependencies: Phase 3 and approved priority/notification specs.

## Phase 5: Frontend Dashboard

1. Objective: implement Vite React seller dashboard, GraphQL read consumption, and operations view.
2. Files affected: frontend source files and package scripts.
3. Acceptance criteria: dashboard lists orders, filters/searches, shows questions, supports replies, and displays unresolved priority.
4. Verification commands: `corepack pnpm run test`, `corepack pnpm run build`.
5. Dependencies: Phase 3 and approved frontend contract.

## Post-MVP Nice-to-Haves

1. Pagination for large seller order and Operations queues.
2. Expanded buyer question history UI beyond the implemented order-detail question flow.
3. Public buyer-question ingestion APIs and webhook subscriptions for external applications.
4. Operations-configurable priority thresholds.
5. Reminder cadence configuration for unresolved questions.
6. GraphQL mutations, subscriptions, federation, schema stitching, or external/public GraphQL contracts.
7. gRPC internal service communication if the monolith is split into services.
8. Infrastructure as Code for deployable environments.
9. Event streaming with SQS, Kafka, or equivalent peak-load strategies.
10. Read-optimized storage for immutable orders and resolved questions.
11. Self-healing infrastructure strategies.
12. Backend-for-Frontend or micro-frontend evolution for more complex role-specific UI.

## Phase 6: Final Documentation

1. Objective: document setup, run, test, and technical notes.
2. Files affected: root `README.md` and related docs.
3. Acceptance criteria: backend and frontend instructions are complete and accurate.
4. Verification commands: `corepack pnpm run test`, `corepack pnpm run build`, `go test ./...`.
5. Dependencies: Phases 1 through 5.