# Frontend Architecture

## Current implementation

The frontend is a Vite React TypeScript app under [../../frontend](../../frontend). It uses feature-based organization under [../../frontend/src/features](../../frontend/src/features) and browser-session-only authentication state.

## Runtime constraints

- Vite runs on port 3000 and host `0.0.0.0`.
- Browser API calls derive the backend URL by replacing frontend port 3000 with backend port 8080.
- Authenticated state and access tokens live in `sessionStorage` only.

## Routing

[../../frontend/src/App.tsx](../../frontend/src/App.tsx) owns the minimal browser route handling:

- `/login`
- `/`
- `/seller/orders`
- `/seller/orders/:orderId`
- `/operations/questions`
- `/buyer/orders`
- `/buyer/orders/:orderId`

Routes are protected by seeded roles. Role-inappropriate access redirects to the authenticated user's default route.

## Feature boundaries

- [../../frontend/src/auth](../../frontend/src/auth): frontend auth types, API helper, seeded role options, backend URL derivation, and session storage.
- [../../frontend/src/features/auth](../../frontend/src/features/auth): login UI.
- [../../frontend/src/features/dashboard](../../frontend/src/features/dashboard): authenticated layout, navigation, and logout.
- [../../frontend/src/features/graphql](../../frontend/src/features/graphql): authenticated read-only GraphQL client.
- [../../frontend/src/features/seller](../../frontend/src/features/seller): seller orders, questions, and seller reply command.
- [../../frontend/src/features/buyer](../../frontend/src/features/buyer): buyer orders, buyer question creation, seller-answer review, and buyer approval.
- [../../frontend/src/features/operations](../../frontend/src/features/operations): Operations priority queue and manual review trigger.

## Data access

- GraphQL is used for authenticated role-shaped reads where useful.
- REST remains the command path for state changes.
- Frontend role visibility is enforced independently of GraphQL response shape.

## Verification

From [../../frontend](../../frontend):

```bash
corepack pnpm run test
corepack pnpm run build
```

Run E2E tests when login, route protection, preview CORS, or smoke flows change:

```bash
corepack pnpm run install-e2e
corepack pnpm run test-e2e
```
