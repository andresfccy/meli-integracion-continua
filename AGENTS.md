# Agent Operating Guide

This file is the normative entry point for AI agents working in this repository.

## Initial reading order

1. [docs/README.md](docs/README.md) for the documentation map and precedence.
2. [specs/README.md](specs/README.md) and the relevant spec before behavior changes.
3. [docs/decisions/README.md](docs/decisions/README.md) before reversing or extending prior choices.
4. [docs/architecture/README.md](docs/architecture/README.md) before changing boundaries.
5. [docs/plans/implementation-plan.md](docs/plans/implementation-plan.md) before claiming implementation status.
6. [docs/development/TESTING_GUIDELINES.md](docs/development/TESTING_GUIDELINES.md) before verification.

## Source-of-truth hierarchy

When documents disagree, use this order:

1. Current code and package metadata for implemented reality.
2. Active specs for approved behavior and contracts.
3. [docs/plans/implementation-plan.md](docs/plans/implementation-plan.md) for implementation status.
4. Decisions and ADRs for accepted direction and supersession history.
5. Architecture docs for boundaries and rationale.
6. README files for onboarding and local maps.
7. Archive files only for historical context.

## Documentation ownership

- Specs define requirements, contracts, and verifiable behavior.
- Decisions explain durable choices and superseded alternatives.
- Architecture docs describe current boundaries, approved direction, and future options.
- Plans track implementation status or future work.
- Development docs define workflow, style, contribution, and verification rules.
- Archive docs preserve history and deprecated guidance; they are not active sources of truth.

Do not duplicate definitions. Link to the owning document instead.

## Required workflow

1. Inspect repository state, relevant code, package metadata, docs, and tests.
2. Run `git status --short` and protect unrelated user work.
3. Plan the smallest scoped change.
4. Confirm the change fits approved specs and decisions.
5. Implement without unrelated refactors.
6. Run relevant verification.
7. Update the owning documentation if behavior, contracts, commands, or structure changed.
8. Review diffs and report changed files.

## Stop conditions

Stop and ask for direction when:

- Required behavior is missing from specs.
- A change would alter accepted requirements or contracts without approval.
- A new dependency is needed but not justified.
- Requested work requires modifying read-only paths.
- Verification fails for reasons not understood after focused investigation.

## Guardrails

- Do not modify read-only path patterns `backend/src/test/*` or `.vscode/tasks.json` when present.
- Do not treat [docs/plans/implementation-plan.md](docs/plans/implementation-plan.md) as a disposable standalone plan.
- Record durable plan history in [docs/archive/BINACLE.md](docs/archive/BINACLE.md) when a task explicitly creates standalone plan files.
- Move executed standalone plans to archive or delete them only after their content is preserved.
- Never mark work complete without checking implementation and tests.
- Keep REST as the command boundary and GraphQL as read-only composition unless specs and decisions change.
- Preserve backend dependency direction: `domain <- application <- adapters <- bootstrap/cmd`.

## Verification baseline

- Backend code changes: `gofmt`, `go test ./...`, and `go vet ./...` from [backend](backend).
- Frontend code changes: `corepack pnpm run test` and `corepack pnpm run build` from [frontend](frontend).
- Login, route protection, preview CORS, or smoke-flow changes: run Playwright E2E when available.
- Documentation-only changes: validate links/references and run code verification only when technical claims, commands, or documented routes changed.

## Historical documents

Files under [docs/archive](docs/archive) are historical. Do not rely on them for current behavior unless active specs, decisions, plans, or architecture docs explicitly promote the content.
