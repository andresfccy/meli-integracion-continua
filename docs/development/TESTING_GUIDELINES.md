# Testing Guidelines

## Backend

Run from [../../backend](../../backend):

```bash
go test ./...
go vet ./...
```

Format changed Go files with `gofmt`.

## Frontend

Run from [../../frontend](../../frontend):

```bash
corepack pnpm run test
corepack pnpm run build
```

`corepack pnpm run test-frontend` runs the Vitest suite. `corepack pnpm run test` runs frontend tests and backend Go tests through the configured script.

## E2E

Run when login, protected navigation, preview CORS, buyer flow, Operations seller search, or manual priority-review smoke behavior changes:

```bash
corepack pnpm run install-e2e
corepack pnpm run test-e2e
```

## Documentation

For documentation-only changes:

- Validate links and moved path references.
- Run code verification when technical claims, scripts, routes, or command guidance changed.
- Record any command that cannot run and why.

## Current full roadmap verification

The broad verification set is:

- Backend: `go test ./...`, `go vet ./...`.
- Frontend: `corepack pnpm run test`, `corepack pnpm run build`.
- E2E when relevant: `corepack pnpm run test-e2e`.
