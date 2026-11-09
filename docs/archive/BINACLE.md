# Binacle: Agent plans

## Plan file policy

Every agent-created `*-PLAN.md` must be recorded in this file. Once the plan is executed, the standalone plan file must be deleted so `docs/archive/BINACLE.md` remains the durable plan history.

## Date
2026-07-15

## Objective
Execute `daily-priority-scheduler-PLAN.md` for automatic daily priority review scheduling, scheduled execution-window acquisition, and Operations manual trigger support, then remove the standalone plan file after completion.

## Plan
1. Keep Operations priority review as the single application use case for eligibility, recategorization, batching, and notification driver delivery.
2. Add an outbound execution acquisition port and a thread-safe in-memory implementation for scheduled review windows.
3. Add an inbound scheduler adapter that runs daily at 00:00 UTC and calls the same use case as the manual endpoint.
4. Wire scheduler configuration and non-blocking startup from bootstrap.
5. Protect the manual priority-review endpoint with an Operations access token and expose a manual Operations UI trigger.
6. Add backend and frontend tests for scheduler timing, duplicate execution windows, authorization, and trigger feedback.
7. Update architecture, specs, backend docs, test plan, and implementation status.

## Completion status
Done.

## Verification run
- Passed: `cd /projects/challenge/backend && gofmt -w ./cmd ./internal && go test ./... && go vet ./...`.
- Passed: `cd /projects/challenge/frontend && corepack pnpm run test-frontend && corepack pnpm run build`.

## Notes
- Scheduler is an inbound adapter and does not contain priority, batching, or notification-channel rules.
- Scheduled duplicate windows return `already_executed`; manual Operations triggers remain unrestricted by the scheduled execution lock.
- The standalone plan file was deleted after verification.

---

## Date
2026-07-14

## Objective
Execute `task-016-operations-seller-search-PLAN.md` for TASK-016 Operations seller search verification, then remove the standalone plan file after completion.

## Plan
1. Replace exact Operations seller-id filtering with one seller search value matching seller id, name, or email.
2. Align REST and GraphQL Operations queue filter names with the approved `seller` contract.
3. Update the frontend Operations seller field to submit the single `seller` search value.
4. Add focused backend service, HTTP, GraphQL, frontend, and E2E coverage for seller name, email, and id filtering.
5. Mark TASK-016 completed after verification.
6. Run `gofmt`, `go test ./...`, `go vet ./...`, `corepack pnpm run test`, `corepack pnpm run build`, and `corepack pnpm run test-e2e`.
7. Review diffs and delete the standalone plan file.

## Completion status
Done.

## Verification run
- `gofmt -w` on changed Go files from `backend/`: passed.
- `go test ./...` from `backend/`: passed.
- `go vet ./...` from `backend/`: passed.
- `corepack pnpm run test` from `frontend/`: passed.
- `corepack pnpm run build` from `frontend/`: passed.
- `corepack pnpm run test-e2e` from `frontend/`: passed.

## Notes
- Operations seller search now uses one `seller` filter value matching seller name, email, or identifier.
- REST, GraphQL, frontend, and E2E coverage now verify the seller search path.
- No read-only files were modified.

---

## Date
2026-07-14

## Objective
Execute `task-015-buyer-frontend-PLAN.md` for TASK-015 buyer frontend routes and smoke coverage, then remove the standalone plan file after completion.

## Plan
1. Audit current buyer frontend route, navigation, list, detail, creation, approval, duplicate-conflict, and approval-conflict behavior against TASK-015.
2. Fill any buyer route, navigation, or UI display gaps while preserving REST commands for question creation and approval.
3. Add Playwright smoke coverage for buyer login, buyer question creation, seller answer, buyer re-login, and buyer approval.
4. Update frontend and root documentation so buyer approval UI and E2E coverage are current behavior.
5. Mark TASK-015 completed after verification.
6. Run `corepack pnpm run test`, `corepack pnpm run build`, and `corepack pnpm run test-e2e` from `frontend/`.
7. Review diffs and delete the standalone plan file.

## Completion status
Done.

## Verification run
- `corepack pnpm run test` from `frontend/`: passed.
- `corepack pnpm run build` from `frontend/`: passed.
- `corepack pnpm run test-e2e` from `frontend/`: passed.

## Notes
- Buyer routes and navigation were already present and remain role-protected.
- Buyer order detail now tolerates omitted empty `questions` from the API and displays the empty state without lifecycle labels.
- Buyer order UI shows item quantity and unit price context in list/detail views.
- Playwright now covers buyer question creation, seller answer, buyer re-login, and answer approval.
- No read-only files were modified.

---

## Date
2026-07-14

## Objective
Execute TASK-013 buyer approval lifecycle from `task-013-buyer-approval-PLAN.md` and remove the standalone plan file after completion.

## Plan
1. Change seller replies to move questions to `seller_answered` while keeping them unresolved for Operations.
2. Add the authenticated buyer approval REST command for the owning buyer.
3. Close the lifecycle only after buyer approval, recording approval metadata and returning terminal `closed` lifecycle state.
4. Expose lifecycle and approval metadata through REST and GraphQL read models.
5. Update Seller, Buyer, and Operations UI lifecycle labels and buyer approval controls.
6. Add backend and frontend coverage for seller-answer and buyer-approval transitions.
7. Update specs and implementation status, run listed verification, review diffs, and delete the standalone plan file.

## Completion status
Done.

## Verification run
- `gofmt -w ./cmd ./internal` from `backend/`: passed.
- `go test ./...` from `backend/`: passed.
- `go vet ./...` from `backend/`: passed.
- `corepack pnpm run test` from `frontend/`: passed.
- `corepack pnpm run build` from `frontend/`: passed.

## Notes
- Seller replies now set `lifecycleStatus=seller_answered` and keep `status=open` until buyer approval.
- Buyer approval records `buyerApprovedAt`, sets `status=resolved`, and returns terminal `lifecycleStatus=closed`.
- Operations reads continue to include `question_active` and `seller_answered` open questions until approval closes them.
- No read-only files were modified.

---

## Date
2026-07-14

## Objective
Execute TASK-008 final project documentation from `task-008-docs-PLAN.md` and remove the standalone plan file after completion.

## Plan
1. Update root documentation to reflect the completed MVP implementation, setup, run, verification, and limitations.
2. Correct backend setup, run, folder map, package responsibility, GraphQL, Operations, and notification documentation.
3. Correct frontend setup, run, runtime constraints, GraphQL Bearer-token, Seller, and Operations documentation.
4. Confirm `docs/README.md` reference-map guidance points to current source-of-truth files.
5. Mark TASK-008 completed after verification.
6. Run `corepack pnpm run test`, `corepack pnpm run build`, and `go test ./...`, then delete the standalone plan file.

## Completion status
Done.

## Verification run
- `corepack pnpm run test` from `frontend/`: passed.
- `corepack pnpm run build` from `frontend/`: passed.
- `go test ./...` from `backend/`: passed.

## Notes
- Final documentation now reflects completed MVP backend and frontend behavior through TASK-008.
- Root setup and run guidance includes backend port 8080, frontend port 3000, host `0.0.0.0`, Corepack pnpm, and backend URL derivation.
- MVP limitations now distinguish simulated/manual behavior from out-of-scope post-MVP capabilities.
- No read-only files were modified.

---

## Date
2026-07-14

## Objective
Execute TASK-007 Operations priority UI from `task-007-operations-priority-PLAN.md` and remove the standalone plan file after completion.

## Plan
1. Complete the existing Operations priority queue UI inside the authenticated dashboard.
2. Reuse the authenticated GraphQL `operationsPriorityQueue` read path and role-gated `/operations/questions` route.
3. Add Operations filters for seller, buyer text, order status, date range, and priority classification.
4. Render priority score, classification, waiting time, order value, urgency words, dictionary version, seller, buyer, and order metadata.
5. Add focused frontend coverage for role visibility, protected route behavior, filters, priority factors, and empty queue state.
6. Update TASK-007 implementation status.
7. Run `corepack pnpm run test` and `corepack pnpm run build`, then delete the standalone plan file.

## Completion status
Done.

## Verification run
- `corepack pnpm run test` from `frontend/`: passed.
- `corepack pnpm run build` from `frontend/`: passed.

## Notes
- Operations users can access the role-gated priority queue inside the dashboard; seller-only users are redirected away from the Operations route.
- The Operations queue shows unresolved question priority score, classification, waiting time, order value, urgency words, dictionary version, seller, buyer, order, and timestamp metadata.
- The Operations filters cover seller, buyer text, order status, date range, and priority classification and are sent through authenticated GraphQL reads.
- The empty unresolved queue state is covered by frontend tests.
- No read-only files were modified.

---

## Date
2026-07-14

## Objective
Execute TASK-006 seller dashboard UI from `task-006-seller-dashboard-PLAN.md` and remove the standalone plan file after completion.

## Plan
1. Add seller reply REST command support using the existing backend URL derivation helper.
2. Complete seller order, filter, detail, buyer question, reply, loading, error, success, and empty UI states.
3. Add approved order status controls and clear-filter behavior.
4. Add focused frontend tests for seller orders, filters, empty states, questions, replies, validation, and command URL derivation.
5. Update frontend documentation and TASK-006 implementation status.
6. Run `corepack pnpm run test` and `corepack pnpm run build`, then delete the standalone plan file.

## Completion status
Done.

## Verification run
- `corepack pnpm run test` from `frontend/`: passed.
- `corepack pnpm run build` from `frontend/`: passed.

## Notes
- Seller dashboard now shows approved status filters, clear-filter behavior, order summaries, order detail, item pricing, buyer questions, and empty/loading/error/success states.
- Seller replies are submitted through `POST /api/sellers/{sellerId}/questions/{questionId}/reply` using the shared frontend-to-backend port derivation helper.
- Successful replies update the selected question as resolved and display the seller reply without adding GraphQL mutations.
- No read-only files were modified.

---

## Date
2026-07-14

## Objective
Execute TASK-005A seeded role options from `task-005a-seeded-roles-PLAN.md` and remove the standalone plan file after completion.

## Plan
1. Add clone-safe seeded account listing without leaking passwords.
2. Add a safe seeded role option read model with deterministic ordering.
3. Expose `GET /api/roles/seeded-users` and return no seeded passwords or access tokens.
4. Display informational seeded role options on the login page while keeping authenticated state sourced from `POST /api/auth/login`.
5. Add backend and frontend coverage for response safety, role option loading, and session-state boundaries.
6. Update TASK-005A status and documentation.
7. Run `gofmt`, `go test ./...`, `go vet ./...`, `corepack pnpm run test`, and `corepack pnpm run build`, then delete the standalone plan file.

## Completion status
Done.

## Verification run
- `gofmt -w` on changed Go files from `backend/`: passed.
- `go test ./...` from `backend/`: passed.
- `go vet ./...` from `backend/`: passed.
- `corepack pnpm run test` from `frontend/`: passed.
- `corepack pnpm run build` from `frontend/`: passed.

## Notes
- Added `GET /api/roles/seeded-users` for safe seeded role option discovery.
- The endpoint returns account id, email, display name, roles, and optional role-specific identifiers without passwords or access tokens.
- The login page displays the role options as informational content only; authenticated session state still comes from `POST /api/auth/login`.
- No read-only files were modified.

---

## Date
2026-07-14

## Objective
Execute TASK-005 simulated batched priority notifications from `task-005-notifications-PLAN.md` and remove the standalone plan file after completion.

## Plan
1. Extend Operations storage access with clone-safe notification channel, notifiable, and notification event methods.
2. Trigger notification logging during priority review for open High Priority and Critical questions only.
3. Resolve seller notifiable targets and enabled channels, preserving future Slack and SMS extensibility.
4. Batch qualifying notifications by notifiable target, channel, and UTC review date with duplicate-safe grouping keys.
5. Add service, HTTP, and seed store coverage for batching, lower-priority exclusion, disabled-channel skipping, and duplicate review idempotency.
6. Update TASK-005 implementation status and backend documentation.
7. Run `gofmt`, `go test ./...`, and `go vet ./...` from `backend/`, then delete the standalone plan file.

## Completion status
Done.

## Verification run
- `gofmt -w ./cmd ./internal` from `backend/`: passed.
- `go test ./...` from `backend/`: passed.
- `go vet ./...` from `backend/`: passed.

## Notes
- Priority review now logs simulated batched notification events for open High Priority and Critical questions.
- Notification batches group by seller notifiable target, enabled channel, and UTC review date and skip duplicate grouping keys on repeated manual reviews.
- Medium and Low questions, resolved questions, missing notifiable targets, and disabled channels do not produce notification events.
- No read-only files were modified.

---

## Date
2026-07-14

## Objective
Execute TASK-004A priority review trigger from `task-004a-priority-review-PLAN.md` and remove the standalone plan file after completion.

## Plan
1. Reuse the existing Operations service review behavior without duplicating `EvaluatePriority`.
2. Add `POST /api/operations/priority-review` with optional `reviewedAt` request parsing and default service-clock behavior.
3. Return review counts including `notificationBatchesLogged`, defaulting to zero until TASK-005.
4. Add service and HTTP coverage for 2-day eligibility, resolved-question skipping, default timestamps, and request validation.
5. Run `gofmt`, `go test ./...`, and `go vet ./...` from `backend/`.
6. Update TASK-004A implementation status and backend documentation.
7. Record completion details and delete the standalone plan file.

## Completion status
Done.

## Verification run
- `gofmt -w ./cmd ./internal` from `backend/`: passed.
- `go test ./...` from `backend/`: passed.
- `go vet ./...` from `backend/`: passed.

## Notes
- Added `POST /api/operations/priority-review` as the manual trigger for daily priority review behavior.
- The endpoint recategorizes only open questions whose latest categorization is at least 2 days old and skips resolved or recent questions.
- The response includes review timestamp, eligible question count, recategorized question count, and zero notification batches until TASK-005 implements notification dispatch.
- No read-only files were modified.

---

## Date
2026-07-14

## Objective
Execute TASK-004B Operations unresolved question queue API from `task-004b-operations-queue-PLAN.md` and remove the standalone plan file after completion.

## Plan
1. Add a `GET /api/operations/questions` HTTP handler backed by the existing Operations priority queue service.
2. Wire the Operations route through the router and bootstrap composition root.
3. Add focused HTTP coverage for unresolved-only results, documented priority fields, deterministic order, filters, and endpoint errors.
4. Run `gofmt`, `go test ./...`, and `go vet ./...` from `backend/`.
5. Update TASK-004B implementation status and backend documentation.
6. Record completion details and delete the standalone plan file.

## Completion status
Done.

## Verification run
- `gofmt -w ./cmd ./internal` from `backend/`: passed.
- `go test ./...` from `backend/`: passed.
- `go vet ./...` from `backend/`: passed.

## Notes
- Added `GET /api/operations/questions` with unresolved-only Operations queue reads backed by the existing Operations service.
- The endpoint supports seller, buyer text, order status, date range, and priority classification filters.
- Responses include priority score, class, waiting hours, order value, matched urgency words, seller, buyer, order, and optional order item summaries.
- TASK-004A priority review endpoint and notification dispatch remain out of scope.
- No read-only files were modified.

---

## Date
2026-07-14

## Objective
Execute TASK-004 priority integration from `task-004-priority-PLAN.md` and remove the standalone plan file after completion.

## Plan
1. Add focused priority policy unit coverage for scoring factors, classifications, matched urgency words, and dictionary metadata.
2. Recompute unresolved Operations queue priority with `EvaluatePriority` at read time.
3. Preserve documented deterministic Operations queue ordering by score, question creation time, order value, and question identifier.
4. Add application-service review behavior that recategorizes only eligible open questions using the same priority policy.
5. Add Operations service coverage for unresolved sorting, read-time evaluation, and review eligibility.
6. Update TASK-004 implementation status after backend verification.
7. Run `gofmt`, `go test ./...`, and `go vet ./...`, then delete the standalone plan file.

## Completion status
Done.

## Verification run
- `gofmt -w ./cmd ./internal` from `backend/`: passed.
- `go test ./...` from `backend/`: passed.
- `go vet ./...` from `backend/`: passed.

## Notes
- Scope excludes `TASK-004A` and `TASK-004B` HTTP endpoints.
- Operations queue reads now recompute unresolved question priority with the same `EvaluatePriority` policy used at creation time.
- Review-time recategorization behavior is available at the Operations application-service layer for future endpoint wiring.
- No read-only files were modified.

---

## Date
2026-07-14

## Objective
Execute TASK-003B GraphQL read composition from `task-003b-graphql-PLAN.md` and remove the standalone plan file after completion.

## Plan
1. Update contracts and decisions for JWT-authenticated GraphQL reads.
2. Add backend GraphQL and JWT dependencies.
3. Extend seeded authentication to issue and validate short-lived JWT access tokens for GraphQL reads.
4. Add read-only Operations priority queue composition with documented filters and deterministic sorting.
5. Add a read-only GraphQL schema and `POST /graphql` handler for `viewer`, `sellerDashboard`, `sellerOrderDetail`, and `operationsPriorityQueue`.
6. Wire GraphQL through backend bootstrap and CORS `Authorization` support while keeping REST as the command API.
7. Connect seller and Operations frontend shells to authenticated GraphQL reads.
8. Add backend and frontend coverage for JWT-protected GraphQL reads and mutation rejection.
9. Update implementation status and documentation, then delete the standalone plan file.

## Completion status
Done.

## Verification run
- `gofmt -w ./cmd ./internal` from `backend/`: passed.
- `go test ./...` from `backend/`: passed.
- `go vet ./...` from `backend/`: passed.
- `corepack pnpm run test` from `frontend/`: passed.
- `corepack pnpm run build` from `frontend/`: passed.

## Notes
- JWT scope is limited to GraphQL reads.
- REST remains the MVP command API.
- No read-only files were modified.

---

## Date
2026-07-14

## Objective
Execute TASK-003A backend-only buyer question creation from `task-003a-PLAN.md` and remove the standalone plan file after completion.

## Plan
1. Extend the question store port and in-memory store for duplicate-open validation.
2. Add reusable creation-time priority dictionary matching and classification metadata.
3. Add question creation service behavior with validation for order, buyer, body, optional order item, and duplicate open questions.
4. Add the `POST /api/orders/{orderId}/questions` REST handler and route.
5. Return created question priority metadata and next-review notification eligibility for High Priority and Critical questions.
6. Add service and HTTP tests for success and error cases.
7. Run Go formatting, tests, and vetting.
8. Update TASK-003A implementation status and delete the executed standalone plan file.

## Completion status
Done.

## Verification run
- `gofmt -w ./cmd ./internal` from `backend/`: passed.
- `go test ./...` from `backend/`: passed.
- `go vet ./...` from `backend/`: passed.

## Notes
- Scope is backend-only for TASK-003A.
- Question creation marks notification eligibility only; it does not send notification events immediately.
- No read-only files were modified.

---

## Date
2026-07-14

## Objective
Execute TASK-003 seller reply implementation from `task-003-seller-replies-PLAN.md` and remove the standalone plan file after completion.

## Plan
1. Add question read and write ports for seller order detail and reply behavior.
2. Add in-memory store helpers to list, find, and save cloned questions.
3. Populate seller order detail responses with buyer questions, optional order item references, and replies.
4. Implement seller reply behavior for open questions with trimmed body validation.
5. Resolve accepted seller replies by default and reject resolved questions.
6. Add the seller reply REST route and route it alongside existing seller order endpoints.
7. Wire the questions service through bootstrap and router composition.
8. Add service and HTTP tests for order questions, successful replies, validation, unknown resources, non-owned questions, and resolved-question conflicts.
9. Update backend documentation and TASK-003 implementation status after verification.
10. Delete the executed standalone TASK-003 plan file.

## Completion status
Done.

## Verification run
- `gofmt -w ./internal` from `backend/`: passed.
- `go test ./...` from `backend/`: passed.
- `go vet ./...` from `backend/`: passed.

## Notes
- Implementation stayed backend-only for TASK-003.
- No buyer question creation, priority scoring, notifications, GraphQL, or frontend dashboard behavior was introduced.
- No read-only files were modified.

---

## Date
2026-07-14

## Objective
Execute TASK-002 seller order API implementation from `task-002-seller-orders-PLAN.md` and remove the standalone plan file after completion.

## Plan
1. Add a seller order read port for sellers, buyers, and orders.
2. Add in-memory store finder/list methods that preserve cloned order item slices.
3. Implement an orders application service for seller-owned order list and detail reads.
4. Apply status, buyer text, and inclusive date range filters with deterministic sorting.
5. Add REST handlers for seller order list and detail endpoints.
6. Wire the orders service through bootstrap and router composition.
7. Add service and HTTP tests for success, filters, validation, unknown sellers, and non-owned orders.
8. Mark TASK-002 completed, record verification, and delete the standalone plan file.

## Completion status
Done.

## Verification run
- `gofmt` on changed Go files: passed.
- `go test ./...` from `backend/`: passed.
- `go vet ./...` from `backend/`: passed.

## Notes
- Implementation stayed backend-only for TASK-002.
- Order detail returns an empty `questions` array; question population remains in TASK-003.
- No read-only files were modified.

---

## Date
2026-07-15

## Objective
Execute `priority-seed-mismatch-PLAN.md` to make seed data coherent with current priority definitions, expand demo coverage, and show seeded demo passwords on the login page.

## Plan
1. Expand in-memory sellers, buyers, Operations members, products, orders, questions, notification channels, notifiables, accounts, and seeded notification events.
2. Derive every seeded question priority from `EvaluatePriority()` at the stored categorization time.
3. Add seed consistency assertions for account references, order/product ownership, question lifecycle and priority metadata, and notification references.
4. Keep `GET /api/roles/seeded-users` safe while updating deterministic role-option expectations.
5. Display frontend-only demo password hints for known seeded profiles on the login page.
6. Update docs and run focused backend/frontend verification.
7. Delete the executed standalone plan file.

## Completion status
Done.

## Verification run
- `gofmt -w` on changed backend seed and handler test files: passed.
- `go test ./internal/adapters/outbound/persistence/memory ./internal/adapters/inbound/http ./internal/application/orders ./internal/application/operations` from `backend/`: passed.
- `npm test` from `frontend/`: passed, including frontend unit tests and `go test ./...` from `backend/` via the configured script.
- `npm run build` from `frontend/`: passed.
- `go vet ./...` from `backend/`: passed.

## Notes
- Core demo credentials remain unchanged: `seller@example.com`, `ops@example.com`, and `buyer@example.com` with their existing demo passwords.
- Safe seeded role options continue to omit passwords and access tokens.
- No read-only files were modified.

---

## Date
2026-07-13

## Objective
Normalize the extracted requirements before final specification generation for the Seller Dashboard: Orders & Questions MVP.

## Plan
1. Review the current state of all files under `specs/` and the repository context documents before editing.
2. Group requirements into the approved business capabilities: seller orders, buyer questions, Operations priority queue, and priority-triggered notifications.
3. Ensure each requirement belongs to exactly one feature specification.
4. Move cross-cutting concerns into shared artifacts: `specs/data-model.md`, `specs/api-contract.md`, `specs/frontend-contract.md`, `specs/test-plan.md`, and `docs/plans/implementation-plan.md`.
5. Remove implementation concerns from feature specifications while keeping implementation sequencing in `docs/plans/implementation-plan.md`.
6. Apply the resolved MVP decisions to the related specs and shared artifacts.
7. Update `docs/decisions/README.md` with the approved MVP decisions.
8. Update `docs/README.md` so the reference map points to the actual final specification structure.
9. Review changed files and document completion.

## Resolved MVP decisions applied
1. Order filters are status, buyer text search, and date range.
2. Buyer-facing UI is out of scope; buyer questions are provided through seeded data for the initial normalization pass.
3. Priority factors include waiting time, order value, and urgency words from a simple dictionary.
4. Priority thresholds are Critical >= 80, High Priority >= 60, Medium >= 30, and Low < 30.
5. Seller reply resolution was later refined to default resolution in the next-cycle pass.
6. Operations is a dedicated section within the existing dashboard, not a separate application.
7. Seed data must include roles for Seller and Operations usage.

## Files updated
- `specs/SPEC-001-seller-orders.md`
- `specs/SPEC-002-buyer-questions.md`
- `specs/SPEC-003-priority-notifications.md`
- `specs/data-model.md`
- `specs/api-contract.md`
- `specs/frontend-contract.md`
- `specs/test-plan.md`
- `docs/plans/implementation-plan.md`
- `docs/decisions/README.md`
- `docs/README.md`
- `docs/archive/BINACLE.md`

## Completion status
Done.

## Notes
- No application code was implemented.
- No read-only files were modified.
- No new material ambiguity was found that affects architecture or business behavior for this normalization pass.

---

## Date
2026-07-13

## Objective
Add the approved lightweight MVP authentication decision to the normalized specifications and shared artifacts.

## Plan
1. Review the current state of affected files before editing because `docs/README.md` and `docs/archive/BINACLE.md` changed after the previous request.
2. Add seller-order authentication preconditions only to `specs/SPEC-001-seller-orders.md`.
3. Add seeded authentication identity requirements to `specs/data-model.md`.
4. Add credential validation behavior to `specs/api-contract.md` without adding durable backend auth infrastructure.
5. Add login, route protection, role-specific UI, and logout behavior to `specs/frontend-contract.md`.
6. Add authentication traceability and validation coverage to `specs/test-plan.md`.
7. Add an implementation task for minimal seeded-account authentication to `docs/plans/implementation-plan.md`.
8. Record the approved decision in `docs/decisions/README.md` and update `docs/README.md` reference guidance.

## Completion status
Done.

## Notes
- No application code was implemented.
- Authentication remains intentionally lightweight and replaceable after MVP.
- No read-only files were modified.

---

## Date
2026-07-13

## Objective
Execute the next-cycle specification refinement plan before application implementation.

## Plan
1. Refine buyer-question lifecycle to one question and one reply for MVP.
2. Add backend-only question creation for MVP testing and external-ingestion simulation.
3. Resolve seller replies by default.
4. Add the approved urgency dictionary and creation-time priority classification.
5. Add daily 00:00 recategorization for open questions that reached the 2-day threshold.
6. Add notifiable targets, notification channels, and batched notification events.
7. Add Operations filters and frontend loading/skeleton expectations.
8. Add scoped GraphQL read composition for authenticated role-shaped dashboard rendering only.
9. Split mixed backend/frontend implementation tasks.
10. Document post-MVP nice-to-haves in the backlog and architecture notes.

## Completion status
Done.

## Notes
- No application code was implemented in this refinement cycle.
- REST remains the command API; GraphQL is read-only frontend composition for MVP.
- gRPC, buyer-facing UI, public webhooks, pagination, IaC, event streaming, and configurable thresholds remain post-MVP.

---

## Date
2026-07-13

## Objective
Execute the documentation reference cleanup plan and remove superseded standalone plan files.

## Plan
1. Validate that the bootstrap plan content is already represented in current docs.
2. Create the root `README.md` from current project context, decisions, contracts, and verification guidance.
3. Replace stale references to the superseded bootstrap plan with current `docs/plans/implementation-plan.md`, `docs/plans/implementation-backlog.md`, and context docs.
4. Prefix ambiguous references to files under `specs/`.
5. Make `docs/plans/implementation-plan.md` the entry point for current development status and route coverage.
6. Clarify REST and GraphQL route implementation tasks without implementing routes.
7. Align seller reply resolution docs so accepted replies always resolve open questions.
8. Align notification batching docs so High Priority and Critical trigger notification execution, then messages batch by notifiable target, channel, and review date.
9. Document that `docs/archive/PROJECT_FILES_INSTRUCTIONS.md` is initial guidance overridden by current decisions and specs when contradictory.
10. Delete superseded and already executed standalone plan files.

## Completion status
Done.

## Notes
- Documentation-only cleanup; no application code was implemented.
- The bootstrap plan is superseded by `README.md`, `docs/README.md`, `docs/decisions/README.md`, `docs/plans/implementation-backlog.md`, and `docs/plans/implementation-plan.md`.
- No read-only files were modified.

---

## Date
2026-07-13

## Objective
Execute the documentation audit plan and remove the standalone plan file after completion.

## Plan
1. Add explicit implementation-plan route coverage for `GET /api/operations/questions`.
2. Confirm API routes from `specs/api-contract.md` have matching implementation-plan checklist coverage.
3. Clarify per-endpoint API success status codes, especially reply and question creation.
4. Clarify that High Priority and Critical notifications are evaluated and sent during the daily 00:00 review, not immediately at question creation.
5. Add frontend browser route guidance for login, seller orders, seller order detail, and Operations priority queue access.
6. Reconcile read-only path guardrail wording when protected path patterns are absent from a local workspace snapshot.
7. Mark the standalone audit plan complete, record verification, and delete it.

## Completion status
Done.

## Verification run
- `corepack pnpm run test` from `frontend/`: passed.
- `corepack pnpm run build` from `frontend/`: passed; current build script reports the frontend build is not implemented yet.
- `go test ./...` from `backend/`: passed.

## Notes
- Documentation-only cleanup; no application code was implemented.
- Route contracts remain target contracts until implementation tasks are completed.
- `docs/archive/PROJECT_FILES_INSTRUCTIONS.md` was changed outside this cleanup after the audit implementation and was left as-is for this commit.
- No read-only files were modified.

---

## Date
2026-07-13

## Objective
Execute TASK-001 domain seed data implementation using the approved pragmatic hexagonal backend structure.

## Plan
1. Create cohesive backend domain packages for identity, commerce, questions, and notifications.
2. Add typed lifecycle constants and deterministic value types, including integer minor-unit money values.
3. Add purchase-time order item snapshot modeling and derived order totals.
4. Add deterministic in-memory seed data for sellers, buyers, products, orders, questions, accounts, Operations members, notification channels, notifiable targets, and notification events.
5. Add an outbound in-memory persistence adapter that owns seeded state and returns protected copies.
6. Add minimal bootstrap composition and move the command entry point to `backend/cmd/api/main.go`.
7. Add seed consistency and domain invariant tests.
8. Add backend reference documentation.
9. Update only TASK-001 implementation status after verification.
10. Remove the executed standalone TASK-001 plan file.

## Completion status
Done.

## Verification run
- `gofmt -w ./cmd ./internal` from `backend/`: passed.
- `go test ./...` from `backend/`: passed.
- `go vet ./...` from `backend/`: passed.

## Notes
- TASK-001A and later tasks were not marked started or completed.
- No HTTP handlers, GraphQL resolvers, authentication use cases, priority scoring algorithms, notification dispatch, scheduler behavior, or later application services were implemented.
- Plain seeded passwords remain deterministic in-memory demo credentials only.
- No read-only files were modified.

---

## Date
2026-07-13

## Objective
Execute the OpsTeamMember architecture alignment plan within the existing TASK-001 backend structure.

## Plan
1. Move `OpsTeamMember` and `NewOpsTeamMember()` from `backend/internal/domain/identity/account.go` to `backend/internal/domain/identity/ops_team_member.go`.
2. Keep the identity package as the owner of seeded accounts, roles, and Operations team member identity.
3. Preserve existing seed dataset and in-memory store behavior without changing authentication, Operations use cases, routes, GraphQL, priority scoring, notification dispatch, or schedulers.
4. Update backend architecture documentation to list the explicit Operations identity file.
5. Run Go formatting, tests, and vetting.
6. Delete the executed standalone plan file.

## Completion status
Done.

## Verification run
- `gofmt -w ./internal/domain/identity` from `backend/`: passed.
- `go test ./...` from `backend/`: passed.
- `go vet ./...` from `backend/`: passed.

## Notes
- This was an architectural clarity refactor only.
- No later task behavior was implemented.
- `backend/internal/domain/questions/priority.go` was inspected but not modified.
- No read-only files were modified.

---

## Date
2026-07-13

## Objective
Execute TASK-001B frontend authentication implementation using Vite, React, pnpm, and the latest available TypeScript version.

## Plan
1. Add a Vite React TypeScript frontend foundation.
2. Add an auth API client that derives the backend URL by replacing frontend port 3000 with backend port 8080.
3. Store authenticated user state in `sessionStorage` only.
4. Add a login view with missing-credential and invalid-credential feedback.
5. Add route protection for `/login`, `/`, `/seller/orders`, `/seller/orders/:orderId`, and `/operations/questions`.
6. Add role-specific seller and Operations placeholder shells without implementing later dashboard data behavior.
7. Add logout behavior that clears current browser-session state and returns to `/login`.
8. Add frontend auth tests for login, invalid credentials, protected redirects, role-specific visibility, and logout.
9. Add frontend reference documentation and update TASK-001B implementation status after verification.
10. Delete the executed standalone TASK-001B plan file.

## Completion status
Done.

## Verification run
- `corepack pnpm run test` from `frontend/`: passed.
- `corepack pnpm run build` from `frontend/`: passed.

## Notes
- TASK-002 and later dashboard data behavior were not implemented.
- No GraphQL reads, seller orders data, seller replies, Operations queue data, priority behavior, or notification behavior were implemented.
- Authenticated state is stored only in browser `sessionStorage`.
- No read-only files were modified.

---

## Date
2026-07-14

## Objective
Apply a TASK-001B login hotfix for HackerRank preview CORS during authentication validation.

## Plan
1. Update backend CORS origin matching so preview frontend origins ending in `-3000.us-vmprovider.projects.hrcdn.net` receive `Access-Control-Allow-Origin` on preflight responses.
2. Preserve local `localhost:3000` preflight support and keep unrelated origins rejected without wildcard CORS.
3. Remove temporary plaintext credential logging from the login handler.
4. Add backend regression coverage for the reported preview Origin and existing local preflight behavior.
5. Add Playwright E2E login coverage for seeded seller login, invalid credentials, and absence of browser CORS console errors.
6. Document the hotfix in decisions, architecture notes, README guidance, backend reference notes, and TASK-001B verification notes.

## Completion status
Done.

## Verification run
- `cd /projects/challenge/backend && gofmt -w ./cmd ./internal && go test ./... && go vet ./...`: passed.
- `cd /projects/challenge/frontend && corepack pnpm run test && corepack pnpm run build`: passed.
- `cd /projects/challenge/frontend && corepack pnpm run install-e2e`: completed to install Chromium for Playwright.
- `cd /projects/challenge/frontend && corepack pnpm run test-e2e`: passed.

## Notes
- The hotfix remains scoped to TASK-001B login transport validation.
- No durable authentication sessions, JWTs, OAuth, external identity providers, GraphQL reads, seller order data, seller replies, Operations queue behavior, priority behavior, or notification behavior are introduced.
- No read-only files were modified.

---

## Date
2026-07-14

## Objective
Execute documentation and roadmap updates from `documentation-buyer-ui-roadmap-PLAN.md` and remove the standalone plan file after completion.

## Plan
1. Correct active README, backend, frontend, architecture, context, and decision-log inconsistencies without changing deprecated platform guidance or personal notes.
2. Update buyer-facing question, seller-order, priority-notification, data-model, API, frontend, and test-plan specs for the planned roadmap.
3. Add TASK-009 through TASK-017 to `docs/plans/implementation-plan.md` for documentation correction, buyer-facing work, notification channel drivers, Operations seller-filter verification, and final validation.
4. Keep question lifecycle states scoped to questions, not orders, and make lifecycle UI visible only when questions exist.
5. Define notification delivery through channel drivers selected by channel type, with an MVP console adapter and future Slack/SMS extensibility.
6. Run documentation-scope verification commands and delete the standalone plan file.

## Completion status
Done.

## Verification run
- `corepack pnpm run test` from `frontend/`: passed.
- `corepack pnpm run build` from `frontend/`: passed.
- `go test ./...` from `backend/`: passed.

## Notes
- Active docs now distinguish completed TASK-001 through TASK-008 behavior from planned buyer-facing and notification-driver work.
- Specs now define planned buyer order reads, buyer-created questions, buyer approval of seller answers, question-only lifecycle states, and notification channel drivers.
- `docs/archive/PROJECT_FILES_INSTRUCTIONS.md` and `docs/archive/notes.md` were intentionally left unchanged.
- No read-only files were modified.

---

## Date
2026-07-14

## Objective
Execute `TASK-012-PLAN.md` for authenticated buyer question creation from buyer-owned order detail, then remove the standalone plan file after completion.

## Plan
1. Add question lifecycle status support while preserving current MVP `open`/`resolved` semantics.
2. Add buyer-owned question creation service behavior and authenticated REST command handling.
3. Add focused backend unit and HTTP coverage for success, authorization, validation, ownership, item, and duplicate constraints.
4. Add minimal buyer frontend role routing, order detail read, and create-question form behavior.
5. Add frontend coverage for buyer login, protected route handling, creation success, duplicate conflict, validation, and empty question lifecycle behavior.
6. Update implementation status/docs, run the listed verification commands, review the diff, and remove the standalone plan file.

## Completion status
Done.

## Verification run
- `gofmt` on changed Go files: passed.
- `go test ./...` from `backend/`: passed.
- `go vet ./...` from `backend/`: passed.
- `corepack pnpm run test` from `frontend/`: passed.
- `corepack pnpm run build` from `frontend/`: passed.

## Notes
- Implemented authenticated buyer question creation as a REST command with `lifecycleStatus=question_active` and preserved GraphQL as read-only.
- Scope intentionally excludes buyer answer approval and E2E expansion owned by later tasks.
- The standalone `TASK-012-PLAN.md` file was removed after execution.

---

## Date
2026-07-14

## Objective
Execute the refinement pass from `task-009-017-refinement-PLAN.md` for TASK-009 through TASK-017 and remove the standalone plan file after completion.

## Plan
1. Clarify TASK-010 through TASK-017 wording for lifecycle dependencies, buyer read scope, duplicate active-question detection, approval terminal-state behavior, notification driver timing/error policy, frontend dependency order, Operations seller-filter verification, and final verification expectations.
2. Align SPEC-002, SPEC-003, API, frontend, data model, and test-plan docs where task wording exposed ambiguity.
3. Preserve current TASK-001 through TASK-008 implementation claims and mark buyer-facing/notification-driver behavior as planned.
4. Keep deprecated `docs/archive/PROJECT_FILES_INSTRUCTIONS.md` and personal `docs/archive/notes.md` unchanged.
5. Run documentation-scope verification and review the final diff.

## Completion status
Done.

## Verification run
- `git diff --check`: passed.
- Targeted traceability grep for `FR-QUE-015`, `TC-BUYER-007`, `AC-NOT-008`, TASK-014 notification-driver scope, and active/unclosed duplicate-question wording: passed.
- `git status --short`: reviewed modified files and confirmed `docs/archive/PROJECT_FILES_INSTRUCTIONS.md` and `docs/archive/notes.md` are unchanged.

## Notes
- This is a documentation and planning refinement only; no buyer-facing or notification-driver implementation code is introduced.
- The standalone `task-009-017-refinement-PLAN.md` file was removed after execution.

---

## Date
2026-07-14

## Objective
Execute `task-009-docs-PLAN.md` for TASK-009 documentation corrections and remove the standalone plan file after completion.

## Plan
1. Audit active README, backend README, frontend README, architecture notes, project context, and decisions against TASK-009.
2. Add only necessary backend dependency documentation for current JWT and GraphQL packages.
3. Leave already-clear buyer-facing roadmap and Operations implementation notes unchanged.
4. Keep `docs/archive/PROJECT_FILES_INSTRUCTIONS.md` and `docs/archive/notes.md` unchanged.
5. Mark TASK-009 complete and review the documentation diff.

## Completion status
Done.

## Verification run
- Documentation diff review: passed.
- `git diff --check`: passed.
- `git status --short`: reviewed modified files and confirmed `docs/archive/PROJECT_FILES_INSTRUCTIONS.md` and `docs/archive/notes.md` are unchanged.

## Notes
- Active docs already clearly documented Operations queue integration, manual priority review, notification batching/event persistence, and the planned buyer-facing roadmap, so no redundant detail was added there.
- The standalone `task-009-docs-PLAN.md` file was removed after execution.

---

## Date
2026-07-14

## Objective
Execute `task-010-011-PLAN.md` for TASK-010 lifecycle normalization and TASK-011 backend buyer account/order read support, then remove the standalone plan file after completion.

## Plan
1. Normalize question lifecycle terminology across specs and traceability docs while keeping order status separate.
2. Add seeded Buyer role/account identity and expose `buyerId` through login, seeded-role, and GraphQL viewer read models.
3. Add authenticated buyer-owned order list/detail services, REST routes, and GraphQL read composition without adding buyer frontend routes or buyer command endpoints.
4. Add backend tests for buyer login, buyer order ownership, REST authorization, and GraphQL buyer reads.
5. Run `gofmt`, `go test ./...`, and `go vet ./...`; review diffs and delete the standalone plan file.

## Completion status
Done.

## Verification run
- `gofmt` on changed backend Go files: passed.
- `go test ./...` from `backend/`: passed.
- `go vet ./...` from `backend/`: passed.

## Notes
- Scope intentionally excludes buyer question creation, buyer answer approval, and frontend buyer routes, which remain assigned to later tasks.
- The standalone `task-010-011-PLAN.md` file was removed after execution.

---

## Date
2026-07-14

## Objective
Execute `task-014-notification-drivers-PLAN.md` for TASK-014 notification driver delivery, then remove the standalone plan file after completion.

## Plan
1. Add outbound notification delivery ports and a deterministic payload/error model.
2. Add a channel-type dispatcher and MVP console driver adapter.
3. Inject the dispatcher into Operations priority review and bootstrap.
4. Route newly logged seller notification events from manual priority review through selected drivers while keeping buyer delivery deferred.
5. Add backend coverage for delivery payloads, disabled-channel skips, duplicate-event skips, console output, and unsupported channel errors.
6. Run `gofmt`, `go test ./...`, and `go vet ./...`; review diffs and delete the standalone plan file.

## Completion status
Done.

## Verification run
- `gofmt` on changed backend Go files: passed.
- `go test ./...` from `backend/`: passed.
- `go vet ./...` from `backend/`: passed.

## Notes
- Notification event creation remains separate from driver delivery.
- The MVP console driver is registered for email channels; unsupported enabled channel types return deterministic delivery errors from the driver boundary.
- Buyer notification delivery remains deferred.
- The standalone `task-014-notification-drivers-PLAN.md` file was removed after execution.

---

## Date
2026-07-14

## Objective
Execute `task-017-final-pass-PLAN.md` for TASK-017 final documentation and verification pass, then remove the standalone plan file after completion.

## Plan
1. Audit active root README, project context, decisions, specs, and testing guidance for stale planned/current wording.
2. Update only TASK-017-scoped documentation so completed buyer-facing flow, notification drivers, and Operations seller filter remain traceable.
3. Mark TASK-017 complete in `docs/plans/implementation-plan.md` after verification.
4. Run frontend test/build and backend test/vet verification commands.
5. Review the documentation diff and delete the standalone plan file.

## Completion status
Done.

## Verification run
- `corepack pnpm run test` from `frontend/`: passed.
- `corepack pnpm run build` from `frontend/`: passed.
- `go test ./...` from `backend/`: passed.
- `go vet ./...` from `backend/`: passed.
- `git diff --check`: passed.
- Documentation diff review: passed.

## Notes
- Buyer notification delivery remains deferred until buyer notification rules are approved.
- The standalone `task-017-final-pass-PLAN.md` file was removed after execution.
