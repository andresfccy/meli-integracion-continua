# Frontend

Vite React TypeScript frontend for the Seller Dashboard MVP. It implements seeded login, browser-session auth, protected role routes, GraphQL-backed dashboard reads, REST commands, buyer flows, seller replies, Operations priority review UI, and smoke-test coverage.

## Boundaries

The frontend is organized by feature rather than technical layer. See [../docs/architecture/frontend.md](../docs/architecture/frontend.md) for architecture details.

## Main structure

- `src/App.tsx`: minimal browser routing and role protection.
- `src/auth`: auth types, API helper, backend URL derivation, seeded role options, and session storage.
- `src/features/auth`: login view.
- `src/features/dashboard`: authenticated layout and navigation.
- `src/features/graphql`: authenticated GraphQL read client.
- `src/features/seller`: seller order/question screens and seller reply command.
- `src/features/buyer`: buyer orders, question creation, seller-answer review, and approval.
- `src/features/operations`: Operations priority queue and manual priority review trigger.
- `e2e`: Playwright smoke tests.

## Runtime rules

- Vite listens on `0.0.0.0:3000`.
- Backend calls derive port 8080 from the current browser location.
- Authenticated user state and access token are stored in `sessionStorage` only.
- GraphQL is read-only; REST remains the command path.

## Key references

- Frontend contract: [../specs/frontend-contract.md](../specs/frontend-contract.md)
- API contract: [../specs/api-contract.md](../specs/api-contract.md)
- Testing guidance: [../docs/development/TESTING_GUIDELINES.md](../docs/development/TESTING_GUIDELINES.md)
- Implementation status: [../docs/plans/implementation-plan.md](../docs/plans/implementation-plan.md)

## Commands

From this directory:

```bash
corepack pnpm install
corepack pnpm run test
corepack pnpm run build
corepack pnpm run start
```

E2E when relevant:

```bash
corepack pnpm run install-e2e
corepack pnpm run test-e2e
```
