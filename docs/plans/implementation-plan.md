# Implementation Plan

This file is the primary entry point for agents to understand current implementation status, task readiness, and route coverage. Other docs provide context and contracts, but progress validation should start here.

## TASK-001
### References
`specs/SPEC-001-seller-orders.md`, `specs/SPEC-002-buyer-questions.md`, `specs/SPEC-003-priority-notifications.md`, `specs/data-model.md`
PHASE-2

### Goal
Define backend domain entities and seeded in-memory data for sellers, buyers, products, orders, order items, questions, roles, and notification events.

### Verification
`gofmt`, `go test ./...`, `go vet ./...`

### Checklist
- [x] Seed at least one seller and one Operations role.
- [x] Seed seller and Operations login accounts with email and password.
- [x] Seed orders with multiple products and multiple statuses.
- [x] Seed buyer questions with open and resolved statuses.
- [x] Preserve purchase-time item names and prices.
### Status
Completed

## TASK-001A
### References
`specs/data-model.md`, `specs/api-contract.md`, `specs/frontend-contract.md`, `specs/test-plan.md`, `docs/decisions/README.md`
PHASE-3

### Goal
Implement minimal backend seeded-account authentication without introducing durable sessions, token refresh, OAuth, registration, or external identity providers.

### Verification
`gofmt`, `go test ./...`, `go vet ./...`

### Checklist
- [x] Implement `POST /api/auth/login` credential validation only against in-memory seeded seller and Operations accounts.
- [x] Return authenticated user role context without JWT refresh tokens or external identity provider data.
### Status
Completed

## TASK-001B
### References
`specs/frontend-contract.md`, `specs/test-plan.md`, `docs/decisions/README.md`
PHASE-5

### Goal
Implement frontend login, logout, route protection, and role-specific UI visibility using current browser-session state.

### Verification
`corepack pnpm run test`, `corepack pnpm run build`, `corepack pnpm run test-e2e`

### Hotfix notes
- 2026-07-14: TASK-001B login validation includes a preview CORS hotfix so HackerRank frontend origins ending in `-3000.us-vmprovider.projects.hrcdn.net` receive `Access-Control-Allow-Origin` from the backend preflight response.
- Playwright E2E login coverage verifies seeded seller login, invalid credentials, and absence of browser CORS console errors during the successful login flow.

### Checklist
- [x] Add frontend login view.
- [x] Protect seller and Operations routes or sections from unauthenticated users.
- [x] Show UI components only for the authenticated user's seeded role.
- [x] Clear browser-session authenticated state on logout.
### Status
Completed

## TASK-002
### References
`specs/SPEC-001-seller-orders.md`, `specs/api-contract.md`, `specs/test-plan.md`
PHASE-3

### Goal
Expose seller order listing and order detail APIs with approved filters and deterministic sorting.

### Verification
`gofmt`, `go test ./...`, `go vet ./...`

### Checklist
- [x] Implement `GET /api/sellers/{sellerId}/orders`.
- [x] Implement `GET /api/sellers/{sellerId}/orders/{orderId}`.
- [x] Implement status filter.
- [x] Implement buyer text search.
- [x] Implement date range filter.
- [x] Include buyer, items, purchase-time prices, and totals.
### Status
Completed

## TASK-003
### References
`specs/SPEC-002-buyer-questions.md`, `specs/api-contract.md`, `specs/test-plan.md`
PHASE-3

### Goal
Expose seller reply behavior for open buyer questions.

### Verification
`gofmt`, `go test ./...`, `go vet ./...`

### Checklist
- [x] Return questions in order detail.
- [x] Implement `POST /api/sellers/{sellerId}/questions/{questionId}/reply`.
- [x] Allow reply to open questions.
- [x] Resolve question by default when seller reply is accepted.
- [x] Reject replies to resolved questions.
### Supersession note
TASK-013 supersedes the original closure behavior for the current buyer-facing implementation: accepted seller replies now move questions to `seller_answered` with `status=open`, and buyer approval closes the lifecycle.
### Status
Completed

## TASK-003A
### References
`specs/SPEC-002-buyer-questions.md`, `specs/SPEC-003-priority-notifications.md`, `specs/api-contract.md`, `specs/data-model.md`, `specs/test-plan.md`
PHASE-3, PHASE-4

### Goal
Expose backend-only question creation for MVP testing and external-ingestion simulation with validation, creation-time classification, and notification trigger integration.

### Verification
`gofmt`, `go test ./...`, `go vet ./...`

### Checklist
- [x] Implement `POST /api/orders/{orderId}/questions`.
- [x] Validate order, buyer, body, and optional order item relationship.
- [x] Enforce one open question per buyer and order.
- [x] Classify created questions using the approved urgency dictionary.
- [x] Return notification trigger metadata for High Priority and Critical questions.
### Status
Completed

## TASK-003B
### References
`specs/api-contract.md`, `specs/frontend-contract.md`, `specs/test-plan.md`
PHASE-3, PHASE-5

### Goal
Expose a scoped GraphQL read composition HTTP endpoint for authenticated role-shaped dashboard rendering.

### Verification
`gofmt`, `go test ./...`, `go vet ./...`, `corepack pnpm run test`, `corepack pnpm run build`

### Checklist
- [x] Implement `POST /graphql` as the GraphQL read HTTP route documented in `specs/api-contract.md`.
- [x] Define GraphQL read schema for `viewer`, `sellerDashboard`, `sellerOrderDetail`, and `operationsPriorityQueue`.
- [x] Keep GraphQL resolvers as read-only composition over backend services.
- [x] Keep login, question creation, replies, recategorization, and all mutations on REST.
- [x] Consume GraphQL reads in seller dashboard and Operations read views where useful.
### Status
Completed

## TASK-004
### References
`specs/SPEC-003-priority-notifications.md`, `specs/data-model.md`, `specs/test-plan.md`
PHASE-4

### Goal
Complete priority policy integration for unresolved questions without duplicating the `EvaluatePriority` policy introduced by TASK-003A for creation-time classification.

### Verification
`gofmt`, `go test ./...`, `go vet ./...`

### Checklist
- [x] Score waiting time in `EvaluatePriority` for creation-time classification.
- [x] Score order value in `EvaluatePriority` for creation-time classification.
- [x] Score urgency words from the approved dictionary in `EvaluatePriority` for creation-time classification.
- [x] Classify Low, Medium, High Priority, and Critical in `EvaluatePriority` for creation-time classification.
- [x] Apply deterministic tie-breakers.
- [x] Store dictionary version and matched urgency words for creation-time classification.
- [x] Reuse the same policy for Operations queue reads and daily or manual review recategorization.
- [x] Add focused coverage for unresolved queue sorting and review-time priority behavior.
### Status
Completed

## TASK-004B
### References
`specs/SPEC-003-priority-notifications.md`, `specs/api-contract.md`, `specs/test-plan.md`
PHASE-3, PHASE-4

### Goal
Expose the Operations unresolved question queue API with documented priority sorting and approved filters.

### Verification
`gofmt`, `go test ./...`, `go vet ./...`

### Checklist
- [x] Implement `GET /api/operations/questions`.
- [x] Return only unresolved questions.
- [x] Apply priority sorting using the approved deterministic tie-breakers.
- [x] Support filters for seller, buyer text, order status, date range, and priority classification.
- [x] Return priority score, classification, waiting time, order value, urgency words, seller, buyer, and order summary fields.
### Status
Completed

## TASK-004A
### References
`specs/SPEC-003-priority-notifications.md`, `specs/api-contract.md`, `specs/test-plan.md`
PHASE-4

### Goal
Implement daily 00:00 priority review behavior that recategorizes open questions only when the 2-day waiting-time threshold is reached.

### Verification
`gofmt`, `go test ./...`, `go vet ./...`

### Checklist
- [x] Add recategorization service behavior.
- [x] Skip resolved questions.
- [x] Skip open questions categorized less than 2 days ago.
- [x] Implement `POST /api/operations/priority-review` as the manual review trigger endpoint for verification or scheduler integration.
### Status
Completed

## TASK-005
### References
`specs/SPEC-003-priority-notifications.md`, `specs/api-contract.md`, `specs/test-plan.md`
PHASE-4

### Goal
Simulate batched notifications for High Priority and Critical questions using notifiable targets and notification channels.

### Verification
`gofmt`, `go test ./...`, `go vet ./...`

### Checklist
- [x] Trigger notification execution for High Priority classifications.
- [x] Trigger notification execution for Critical classifications.
- [x] Do not notify Medium or Low questions.
- [x] Keep notification behavior extensible for Slack and SMS.
- [x] Model notifiable targets and enabled notification channels.
- [x] Batch qualifying notifications by notifiable target, channel, and recategorization review date to avoid spam.
### Status
Completed

## TASK-005A
### References
`specs/api-contract.md`, `specs/frontend-contract.md`, `specs/test-plan.md`
PHASE-3, PHASE-5

### Goal
Expose seeded role options needed by the MVP dashboard experience.

### Verification
`gofmt`, `go test ./...`, `go vet ./...`, `corepack pnpm run test`, `corepack pnpm run build`

### Checklist
- [x] Implement `GET /api/roles/seeded-users` without exposing seeded passwords.
- [x] Use `POST /api/auth/login` as the source of authenticated frontend state once login exists.
### Status
Completed

## TASK-006
### References
`specs/SPEC-001-seller-orders.md`, `specs/SPEC-002-buyer-questions.md`, `specs/frontend-contract.md`, `specs/test-plan.md`
PHASE-5

### Goal
Build the seller dashboard UI for orders, filters, questions, and seller replies.

### Verification
`corepack pnpm run test`, `corepack pnpm run build`

### Checklist
- [x] Derive backend URL by replacing port 3000 with 8080.
- [x] Require authenticated seller state before showing the seller dashboard.
- [x] Display order list and approved filters.
- [x] Display order questions.
- [x] Support seller replies that move open questions to `seller_answered` on success.
- [x] Display loading, error, success, and empty states.
### Supersession note
TASK-013 supersedes the original seller-reply UI wording for the current buyer-facing implementation: successful seller replies now display `seller_answered` and await buyer approval before closure.
### Status
Completed

## TASK-007
### References
`specs/SPEC-003-priority-notifications.md`, `specs/frontend-contract.md`, `specs/test-plan.md`
PHASE-5

### Goal
Build the Operations priority section inside the existing dashboard.

### Verification
`corepack pnpm run test`, `corepack pnpm run build`

### Checklist
- [x] Read authenticated Operations role state.
- [x] Display Operations section for Operations role.
- [x] Display unresolved priority queue.
- [x] Add Operations filters for seller, buyer text, order status, date range, and priority classification.
- [x] Show score, classification, and priority factors.
- [x] Display empty queue state.
### Status
Completed

## TASK-008
### References
`README.md`, `docs/README.md`, `docs/decisions/README.md`, `specs/`
PHASE-6

### Goal
Update final project documentation after implementation is complete.

### Verification
`corepack pnpm run test`, `corepack pnpm run build`, `go test ./...`

### Checklist
- [x] Document backend setup and run steps.
- [x] Document frontend setup and run steps.
- [x] Document technical notes and MVP limitations.
- [x] Confirm reference map points to current source-of-truth files.
### Status
Completed

## TASK-009
### References
`README.md`, `backend/README.md`, `frontend/README.md`, `docs/architecture/README.md`, `docs/README.md`, `docs/decisions/README.md`
PHASE-7

### Goal
Correct current documentation inconsistencies so active docs match the implemented TASK-001 through TASK-008 code and describe the planned buyer-facing roadmap without changing deprecated platform guidance or personal notes.

### Verification
Documentation diff review

### Checklist
- [x] Fix backend dependency documentation to acknowledge JWT and GraphQL packages.
- [x] Update architecture notes so Operations queue integration, manual review, notification batching, and event persistence are documented as implemented.
- [x] Add buyer-facing scope notes to active README and reference-map docs.
- [x] Keep `docs/archive/PROJECT_FILES_INSTRUCTIONS.md` and `docs/archive/notes.md` unchanged.
### Status
Completed

## TASK-010
### References
`specs/SPEC-001-seller-orders.md`, `specs/SPEC-002-buyer-questions.md`, `specs/SPEC-003-priority-notifications.md`, `specs/data-model.md`, `specs/api-contract.md`, `specs/frontend-contract.md`, `specs/test-plan.md`
PHASE-7

### Goal
Normalize the buyer-facing question lifecycle before TASK-011 through TASK-015 implementation begins, because lifecycle names affect backend domain values, API response shapes, frontend labels, and test traceability.

### Verification
Documentation diff review and traceability review

### Checklist
- [x] Move the new lifecycle definition to questions, not orders.
- [x] Keep order status separate from question lifecycle status.
- [x] Exclude order-like states such as `created` and `accepted` from question lifecycle values.
- [x] Define question lifecycle values such as `question_active`, `seller_answered`, `buyer_approved`, and `closed`.
- [x] Define how `lifecycleStatus` coexists with the current MVP `open`/`resolved` status.
- [x] Define whether `buyer_approved` is persisted as approval metadata, a visible lifecycle label, or a transition immediately followed by terminal `closed` display.
- [x] Define that orders without questions show no question lifecycle labels.
- [x] Define buyer-created questions and buyer approval of seller answers.
- [x] Update API, frontend, data model, and test traceability docs.
### Status
Completed

## TASK-011
### References
`specs/data-model.md`, `specs/api-contract.md`, `specs/test-plan.md`
PHASE-8

### Goal
Add seeded Buyer account data and authenticated buyer order read support, scoped to backend identity, buyer-owned read services/routes, and authenticated read composition needed by later buyer UI work.

### Verification
`gofmt`, `go test ./...`, `go vet ./...`

### Checklist
- [x] Add typed Buyer role support to seeded accounts.
- [x] Seed at least one Buyer login account linked to an existing buyer.
- [x] Add buyer-owned order list and detail read behavior.
- [x] Prevent buyers from reading orders owned by other buyers.
- [x] Add buyer GraphQL read queries only for list/detail read composition; keep buyer question creation and approval as REST commands.
- [x] Keep GraphQL reads authenticated and role-shaped.
- [x] Do not add buyer-facing frontend routes in this task; TASK-015 owns buyer route implementation after TASK-011 through TASK-013 contracts are available.
### Status
Completed

## TASK-012
### References
`specs/SPEC-002-buyer-questions.md`, `specs/api-contract.md`, `specs/frontend-contract.md`, `specs/test-plan.md`
PHASE-8

### Goal
Implement buyer-facing question creation from buyer-owned order detail.

### Verification
`gofmt`, `go test ./...`, `go vet ./...`, `corepack pnpm run test`, `corepack pnpm run build`

### Checklist
- [x] Add buyer-authenticated REST command for creating a question from an owned order.
- [x] Reuse validation for body, order ownership, optional order item reference, and duplicate active or unclosed question constraints.
- [x] Create questions in `question_active` lifecycle state.
- [x] Add buyer order detail UI question form.
- [x] Display success, duplicate-conflict, validation, loading, error, and empty states.
- [x] Preserve REST as the command boundary; do not introduce GraphQL mutations.
### Status
Completed

## TASK-013
### References
`specs/SPEC-002-buyer-questions.md`, `specs/data-model.md`, `specs/api-contract.md`, `specs/frontend-contract.md`, `specs/test-plan.md`
PHASE-8

### Goal
Implement buyer approval question lifecycle for seller answers.

### Verification
`gofmt`, `go test ./...`, `go vet ./...`, `corepack pnpm run test`, `corepack pnpm run build`

### Checklist
- [x] Change seller reply behavior to move questions to `seller_answered` instead of immediately closing the buyer-facing lifecycle.
- [x] Add buyer approval command for the buyer who owns the order.
- [x] Close the question lifecycle only after buyer approval, using the TASK-010-defined `buyer_approved` and terminal `closed` semantics.
- [x] Keep `question_active` and `seller_answered` questions visible to Operations as unresolved until buyer approval and closure.
- [x] Show question lifecycle labels to authorized users only when questions exist.
- [x] Keep multiple threaded replies out of scope unless SPEC-002 is explicitly expanded.
### Status
Completed

## TASK-014
### References
`docs/architecture/README.md`, `specs/SPEC-003-priority-notifications.md`, `specs/api-contract.md`, `specs/data-model.md`, `specs/test-plan.md`
PHASE-8

### Goal
Deliver every notification through an outbound channel-driver abstraction whose MVP console adapter can log enabled channel payloads while preserving future Slack, SMS, or additional protocol support.

### Verification
`gofmt`, `go test ./...`, `go vet ./...`

### Checklist
- [x] Review the current notification architecture for hexagonal boundaries.
- [x] Define notification delivery ports for outbound channel drivers.
- [x] Keep domain and application logic independent from stdout and channel protocols.
- [x] Implement a channel driver registry or dispatcher selected by channel type.
- [x] Implement the MVP console driver adapter.
- [x] Route seller notification delivery during priority review through selected drivers.
- [x] Keep buyer notification delivery deferred unless finalized buyer notification rules are approved before implementation.
- [x] Skip disabled channels deterministically and return a deterministic delivery error for unsupported enabled channel types.
- [x] Verify High Priority and Critical notification triggers use the driver path during priority review delivery.
### Status
Completed

## TASK-015
### References
`specs/frontend-contract.md`, `specs/test-plan.md`
PHASE-8

### Goal
Add buyer role frontend routes and E2E smoke coverage after TASK-011 through TASK-013 backend/API contracts are available.

### Verification
`corepack pnpm run test`, `corepack pnpm run build`, `corepack pnpm run test-e2e`

### Checklist
- [x] Add Buyer role routing and role-specific navigation.
- [x] Add protected buyer order list route.
- [x] Add protected buyer order detail route.
- [x] Add buyer question creation and answer approval UI coverage.
- [x] Add duplicate active-question and approve-answer conflict UI coverage.
- [x] Add E2E smoke coverage for the buyer flow.
### Status
Completed

## TASK-016
### References
`specs/SPEC-003-priority-notifications.md`, `specs/api-contract.md`, `specs/frontend-contract.md`, `specs/test-plan.md`, `frontend/src/features/operations/OperationsQuestionsShell.tsx`, `backend/internal/application/operations/service.go`
PHASE-8

### Goal
Confirm the Operations priority queue seller search filter works end-to-end as a verification-first task, then add missing focused coverage only where service, HTTP, GraphQL, frontend, or E2E gaps exist.

Seller search must use one input that matches seller name, seller email, or seller identifier.

### Verification
`gofmt`, `go test ./...`, `go vet ./...`, `corepack pnpm run test`, `corepack pnpm run build`, `corepack pnpm run test-e2e`

### Checklist
- [x] Verify backend seller search semantics for seller name, email, and id.
- [x] Verify HTTP and GraphQL seller search filter forwarding.
- [x] Verify frontend seller search input submission from the single Operations seller field.
- [x] Add or update backend unit and HTTP tests only if coverage gaps exist.
- [x] Add or update frontend tests only if coverage gaps exist.
- [x] Add E2E smoke coverage for filtering by seller name, email, or id if no equivalent smoke coverage exists.
### Status
Completed

## TASK-017
### References
`README.md`, `docs/README.md`, `docs/decisions/README.md`, `specs/`, `docs/development/TESTING_GUIDELINES.md`
PHASE-9

### Goal
Run the final documentation and verification pass for the completed buyer-facing, notification-driver, and Operations seller-filter roadmap.

### Verification
`corepack pnpm run test`, `corepack pnpm run build`, `go test ./...`, `go vet ./...` when TASK-011 through TASK-016 introduced Go changes, documentation diff review

### Checklist
- [x] Confirm README and reference-map docs match implemented and deferred status.
- [x] Confirm specs and implementation tasks remain traceable.
- [x] Confirm testing guidance covers buyer-facing flow, notification drivers, and Operations seller filter.
- [x] Run final verification commands.
- [x] Include `go vet ./...` in the final pass because TASK-011 through TASK-014 introduced Go implementation changes.
### Status
Completed

## TASK-018
### References
`specs/SPEC-003-priority-notifications.md`, `specs/api-contract.md`, `specs/frontend-contract.md`, `specs/test-plan.md`, `docs/architecture/README.md`
PHASE-4

### Goal
Add automatic daily 00:00 UTC priority review execution inside the monolith through an inbound scheduler adapter while preserving the existing Operations review use case and notification driver boundaries.

### Verification
`gofmt`, `go test ./...`, `go vet ./...`, `corepack pnpm run test`

### Checklist
- [x] Keep review and recategorization logic in the Operations application use case.
- [x] Add an inbound scheduler adapter that invokes the same use case as the manual HTTP trigger.
- [x] Add outbound scheduled execution-window acquisition and thread-safe in-memory implementation.
- [x] Wire scheduler configuration and non-blocking startup in bootstrap.
- [x] Protect the manual Operations trigger with an Operations access token.
- [x] Allow Operations manual priority review to override notification event deduplication for explicit same-day seller batch resend.
- [x] Keep MVP seeded notification channels email-only while preserving future Slack/SMS driver extensibility.
- [x] Add Operations UI trigger feedback for manual review execution.

### Status
Completed
