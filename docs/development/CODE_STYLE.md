# Code Style

## Go

1. Use the standard library first unless an approved dependency already exists or is justified.
2. Keep packages focused by responsibility.
3. Preserve backend dependency direction: `domain <- application <- adapters <- bootstrap/cmd`.
4. Use explicit request and response structs for HTTP JSON contracts.
5. Return deterministic errors.
6. Format changed Go files with `gofmt`.
7. Avoid global mutable state outside the in-memory storage boundary.

## Frontend

1. Use Vite, React, TypeScript, and pnpm through Corepack.
2. Build by feature, not by technical layer.
3. Derive the backend base URL from `window.location` by replacing frontend port 3000 with backend port 8080.
4. Keep authenticated state in `sessionStorage` only.
5. Avoid unnecessary dependencies.
6. Keep UI state predictable and local unless shared state is required by the spec.

## Documentation

1. Specs are the source of truth for behavior and contracts.
2. Decisions/ADRs own durable technical choices.
3. Architecture docs own boundaries and rationale.
4. Plans own implementation status or future work.
5. Archive docs are historical and non-authoritative.
6. Link to source-of-truth content instead of copying it.
