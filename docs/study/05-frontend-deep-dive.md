# Frontend Deep Dive

## Runtime choices

The frontend uses React, TypeScript, and Vite. Vite serves `0.0.0.0:3000`; the client derives the backend origin by replacing the current port with `8080`.

JavaScript executes in the browser. TypeScript checks source types before build, then disappears from the runtime output. React converts component state into UI. Vite transforms TSX and creates the browser bundle.

## Feature organization

```text
src/auth                    shared session and HTTP client
src/features/auth           login
src/features/dashboard      authenticated shell
src/features/graphql        read client
src/features/seller         seller orders and replies
src/features/buyer          buyer orders, questions, approval
src/features/operations     queue and manual review
```

**[DEFENSA]** Feature organization keeps the UI, actions, and local types for one user capability close together. A global “components/services/pages” split would make each change cross more directories.

## Routing and role visibility

`App.tsx` reads `location.pathname` and the stored session. It:

- redirects unauthenticated users to `/login`,
- chooses a default route from roles,
- renders Seller, Buyer, or Operations only for the matching role,
- redirects role-inappropriate paths.

This is route selection rather than a router library. It is adequate for a few static paths but lacks nested route composition, history abstractions, route-level error boundaries, and test utilities.

Frontend role checks improve usability; backend checks provide security.

## Session and client

`session.ts` owns:

- `User` and `Session` types,
- JSON serialization to `sessionStorage`,
- backend origin derivation,
- bearer header injection,
- JSON parsing,
- shared error conversion.

Centralization avoids repeated headers and error logic.

Limitations:

- parsed session JSON is not runtime-validated,
- all responses are assumed to be JSON,
- 204 or non-JSON responses would fail parsing,
- no abort/cancellation handling,
- token is accessible to injected JavaScript,
- no automatic expiration/logout handling.

## React data flow

The screens follow:

```text
initial state → effect → request → setState → render
```

`useState` stores data and feedback. `useEffect` starts reads after render. Form handlers prevent browser submission, issue REST commands, and reload the page.

Reloading is simple and guarantees a fresh read, but discards local state and is less efficient than updating a cache.

## Seller

Seller list/detail reads use schema-backed GraphQL queries. Replies use REST. Status, buyer text, and date filters are sent to the application through typed GraphQL variables.

## Buyer

Buyer list/detail reads use schema-backed GraphQL composition. Question creation and approval use REST, preserving the command/read boundary.

The form is shown whenever no question has coarse status `open`, so closed historical questions no longer block a new one.

## Operations

The queue read uses schema-backed GraphQL composition. Manual review uses REST. Seller ID/name/email, buyer, status, date, and priority filters are forwarded to the backend. Feedback includes execution status and ID.

## GraphQL client

The client sends standard GraphQL documents and variables to `/graphql`; the role screens specify the fields they need instead of receiving fixed server maps. Backend schema validation catches unknown fields before resolver execution. TypeScript types describe the expected result inside the frontend, while the GraphQL schema is the runtime transport contract.

`gql<T>` posts query text and variables through the shared API client, checks a top-level `errors` array, and returns `data`.

Principal role screens use typed results and `Record<string, unknown>` variables. The endpoint still ignores individual selected fields, so query field shapes are not schema-enforced.

## Package tooling

`package.json` declares pnpm 10.12.1 through Corepack. Scripts preserve:

- `start`: backend and Vite concurrently,
- `build`: TypeScript project build plus Vite,
- `test`: frontend then backend tests,
- `test-e2e`: Playwright.

**[RECOMENDACIÓN]** Replace `latest` dependency ranges with intentional versions and update them through reviewed changes. The lockfile helps current reproducibility, but `latest` weakens future manifest-based installs.

## Code quality observations

**[OBSERVADO]** Seller, Buyer, Operations, session, and GraphQL-client files are formatted and typed. Some small legacy components and the stylesheet remain compact.

Recommended evolution:

- format source,
- split view and data hooks when complexity grows,
- define response types,
- validate API payloads at runtime,
- add accessible status announcements and labels,
- avoid full reloads,
- add error boundaries and request cancellation.

## Verification

The frontend helper tests cover session persistence, backend URL derivation, Bearer injection, and documented role redirect priority.

**[VERIFICADO 2026-07-30]** Installed Vitest passed 5 session/routing assertions; TypeScript plus Vite built the production bundle; and Playwright passed 3 Chromium flows covering authentication, Buyer→Seller→Buyer lifecycle, and Operations seller filtering/manual review. Isolated component coverage remains a future depth improvement.
