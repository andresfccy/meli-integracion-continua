# Test Plan

## 1. Purpose
Connect normalized requirements to acceptance scenarios and test cases.

## 2. Traceability matrix

| Requirement | Acceptance Scenario | Test Case id | Level | Expected result |
| ------ | ------ | ------ | ------ | ------ |
| FR-ORD-001 | AC-ORD-001 | TC-ORD-001 | HTTP + frontend | Seller sees recent seller-owned orders. |
| FR-ORD-002 | AC-ORD-002 | TC-ORD-002 | HTTP + frontend | Order details include buyer, status, date, items, quantities, purchase-time prices, and total. |
| FR-ORD-003 | AC-ORD-003 | TC-ORD-003 | HTTP + frontend | Status filter returns only matching orders. |
| FR-ORD-004 | AC-ORD-004 | TC-ORD-004 | HTTP + frontend | Buyer text search returns matching buyer orders. |
| FR-ORD-005 | AC-ORD-005 | TC-ORD-005 | HTTP + frontend | Date range returns orders within inclusive bounds. |
| FR-ORD-006 | AC-ORD-006 | TC-ORD-006 | Frontend | No-results empty state is displayed. |
| FR-ORD-007 | AC-ORD-002 | TC-ORD-007 | Unit + HTTP | Purchase-time unit prices are preserved in responses. |
| FR-QUE-001 | AC-QUE-001 | TC-QUE-001 | HTTP + frontend | Order questions are displayed in order context. |
| FR-QUE-002 | AC-QUE-001 | TC-QUE-002 | HTTP + frontend | Question details include required fields and optional product reference. |
| FR-QUE-003 | AC-QUE-002 | TC-QUE-003 | HTTP + frontend | Seller can reply to an open question. |
| FR-QUE-004 | AC-QUE-002 | TC-QUE-004 | HTTP + frontend | Seller reply stores the answer, keeps the question unresolved, and moves lifecycle to `seller_answered`. |
| FR-QUE-005 | AC-QUE-004 | TC-QUE-005 | Unit + HTTP + frontend | Resolved question rejects additional reply. |
| FR-QUE-003 | AC-QUE-004A | TC-QUE-011 | HTTP + frontend | A seller cannot see or reply to questions for orders assigned to another seller. |
| FR-QUE-006 | AC-QUE-006 | TC-QUE-006 | HTTP | Backend endpoint creates a valid open buyer question without buyer UI. |
| FR-QUE-007 | AC-QUE-005 | TC-QUE-007 | Frontend | Order with no questions displays an empty state. |
| FR-QUE-008 | AC-QUE-007 | TC-QUE-008 | Unit + HTTP | Backend-only creation rejects an existing open buyer/order question by coarse `status`. |
| FR-QUE-009 | AC-QUE-006 | TC-QUE-009 | Unit + HTTP | Created question is classified at creation time. |
| FR-PRI-001 | AC-PRI-001 | TC-PRI-001 | Unit + HTTP + frontend | Operations queue returns unresolved questions sorted by priority. |
| FR-PRI-002 | AC-PRI-003 | TC-PRI-002 | Unit | Priority score uses waiting time, order value, and urgency words. |
| FR-PRI-003 | AC-PRI-003 | TC-PRI-003 | Unit + HTTP | Score maps to Low, Medium, High Priority, or Critical. |
| FR-PRI-004 | AC-PRI-004 | TC-PRI-004 | Unit + HTTP | Tied scores use deterministic tie-breakers. |
| FR-PRI-005 | AC-PRI-005 | TC-PRI-005 | Frontend | Operations queue appears inside the existing dashboard. |
| FR-PRI-006 | AC-PRI-005 | TC-PRI-006 | HTTP + frontend | Seeded roles support Seller and Operations usage. |
| FR-PRI-007 | AC-PRI-002 | TC-PRI-007 | Unit + HTTP | Resolved questions are excluded from Operations queue. |
| FR-PRI-008 | AC-PRI-006 | TC-PRI-008 | Unit + HTTP | Created question stores priority score, classification, matched words, and dictionary version. |
| FR-PRI-009 | AC-PRI-007, AC-PRI-008 | TC-PRI-009 | Unit + HTTP | Daily review recategorizes only eligible open questions at least 2 days old. |
| FR-PRI-009 | AC-PRI-007, AC-PRI-008 | TC-PRI-012 | Unit | The inbound scheduler computes daily 00:00 UTC windows, invokes the review use case, recovers panics, and stops on cancellation. |
| FR-PRI-009 | AC-PRI-007, AC-PRI-008 | TC-PRI-013 | Unit | Scheduled review execution acquisition allows one execution per UTC review window and returns `already_executed` for duplicates. |
| FR-PRI-010 | AC-PRI-009 | TC-PRI-010 | HTTP + frontend | Operations filters narrow unresolved priority queue results, including seller search by name, email, or id. |
| FR-NOT-001 | AC-NOT-001 | TC-NOT-001 | Unit | High Priority question logs notification event. |
| FR-NOT-002 | AC-NOT-002 | TC-NOT-002 | Unit | Critical question logs notification event. |
| FR-NOT-003 | AC-NOT-001, AC-NOT-002 | TC-NOT-003 | Unit | Email notification is simulated as a logged event. |
| FR-NOT-004 | AC-NOT-001 | TC-NOT-004 | Unit/design review | Notification behavior is channel-extensible without changing priority classification. |
| FR-NOT-005 | AC-NOT-003 | TC-NOT-005 | Unit | Medium and Low questions do not log notification events. |
| FR-NOT-006 | AC-NOT-004 | TC-NOT-006 | Unit | Qualifying notifications are triggered by High Priority/Critical classifications and batched by notifiable target, channel, and review date. |
| FR-AUTH-001 | AC-AUTH-001 | TC-AUTH-001 | HTTP + frontend | Seeded user can log in with valid email and password. |
| FR-AUTH-002 | AC-AUTH-002 | TC-AUTH-002 | HTTP + frontend | Invalid credentials are rejected and do not authenticate the user. |
| FR-AUTH-003 | AC-AUTH-003 | TC-AUTH-003 | Frontend | Unauthenticated users are redirected to the login page for protected routes. |
| FR-AUTH-004 | AC-AUTH-004 | TC-AUTH-004 | Frontend | Seller account can access only seller dashboard UI. |
| FR-AUTH-005 | AC-AUTH-005 | TC-AUTH-005 | Frontend | Operations account can access only Operations role UI inside the dashboard application. |
| FR-AUTH-006 | AC-AUTH-006 | TC-AUTH-006 | Frontend | Logout clears browser-session authenticated state and returns to login. |
| GRAPHQL-READ-001 | N/A | TC-GQL-001 | HTTP | `viewer` query returns authenticated role-shaped viewer data. |
| GRAPHQL-READ-002 | N/A | TC-GQL-002 | HTTP + frontend | `sellerDashboard` query returns seller dashboard read model. |
| GRAPHQL-READ-003 | N/A | TC-GQL-003 | HTTP + frontend | `sellerOrderDetail` query returns order detail read model. |
| GRAPHQL-READ-004 | N/A | TC-GQL-004 | HTTP + frontend | `operationsPriorityQueue` query returns filtered Operations read model. |
| GRAPHQL-READ-005 | N/A | TC-GQL-005 | HTTP | GraphQL reads require a valid, unexpired JWT and reject invalid signatures. |
| GRAPHQL-READ-006 | N/A | TC-GQL-006 | HTTP | GraphQL mutations are rejected because REST remains the command API. |
| GRAPHQL-READ-007 | N/A | TC-GQL-007 | HTTP | `buyerDashboard` returns authenticated buyer-owned order list data. |
| GRAPHQL-READ-008 | N/A | TC-GQL-008 | HTTP | `buyerOrderDetail` returns authenticated buyer-owned order detail data and rejects wrong-role or mismatched-buyer reads. |
| FR-BUYER-001 | AC-BUYER-001 | TC-BUYER-001 | HTTP + frontend | Buyer account sees only buyer-owned orders through protected buyer routes. |
| FR-BUYER-002 | AC-BUYER-002 | TC-BUYER-002 | HTTP + frontend | Buyer question creation creates a `question_active` question. |
| FR-BUYER-002, FR-QUE-015 | AC-BUYER-002 | TC-BUYER-005 | Unit + HTTP + frontend | Buyer question creation rejects duplicate active or unclosed questions for the same buyer and order. |
| FR-BUYER-003 | AC-BUYER-003 | TC-BUYER-003 | Unit + HTTP + frontend | Buyer approval records approval metadata and displays the question lifecycle as closed. |
| FR-BUYER-003 | AC-BUYER-003 | TC-BUYER-006 | HTTP + frontend | Buyer approval rejects wrong-owner, missing-reply, non-`seller_answered`, already-approved, and already-closed conflicts. |
| FR-BUYER-004 | AC-BUYER-004 | TC-BUYER-004 | Frontend | UI hides question lifecycle labels when an order has no questions. |
| FR-QUE-012 | AC-QUE-010 | TC-QUE-010 | Unit + HTTP + frontend | Seller answer moves question lifecycle to `seller_answered`. |
| Buyer endpoint authorization | N/A | TC-BUYER-007 | HTTP + GraphQL | Missing token, wrong role, mismatched `buyerId`, non-owned order, and non-owned question are rejected deterministically. |
| FR-NOT-007 | AC-NOT-005 | TC-NOT-007 | Unit | Delivery selects a notification channel driver by channel type. |
| FR-NOT-008 | AC-NOT-006 | TC-NOT-008 | Unit | MVP console driver writes delivered payloads to stdout. |
| FR-NOT-009 | AC-NOT-007 | TC-NOT-009 | Unit/design review | Application services deliver through driver ports and do not write payloads directly to stdout. |
| FR-NOT-009 | AC-NOT-008 | TC-NOT-010 | Unit | Delivery skips disabled channels and returns deterministic errors for unsupported enabled channel types. |
| FR-NOT-011 | AC-NOT-009 | TC-NOT-011 | Unit + HTTP + E2E | Operations manual priority review can resend same-day seller notification batches through manual override grouping while scheduled duplicate windows remain blocked. |
| FR-PRI-010 | AC-PRI-009 | TC-PRI-011 | Unit + HTTP + frontend + E2E | Operations seller search works by seller name, email, or id from backend service through GraphQL/frontend and E2E smoke coverage. |

## 3. Unit tests
- Validate order total calculation and purchase-time item pricing.
- Validate question lifecycle rules for open and resolved states.
- Validate priority score factors and classification thresholds.
- Validate deterministic priority tie-breakers.
- Validate notification trigger thresholds.
- Validate seeded credential matching rules without external identity dependencies.
- Validate one-open-question-per-buyer/order rule.
- Validate creation-time urgency dictionary matching.
- Validate 2-day recategorization eligibility.
- Validate notification batching grouping keys by notifiable target, channel, and review date.
- Validate buyer account ownership rules for buyer order reads.
- Validate buyer question lifecycle transitions: `question_active`, `seller_answered`, `buyer_approved`, and `closed`.
- Validate that buyer approval records approval metadata and displays `closed` as the terminal lifecycle.
- Validate order-like states such as `created` and `accepted` are not valid question lifecycle values.
- Validate duplicate active or unclosed buyer questions for the same buyer and order are rejected.
- Validate notification channel driver selection, disabled-channel skipping, unsupported-channel deterministic errors, and console driver output payload shape.
- Validate Operations seller search semantics by seller name, email, and id in the application service.
- Validate scheduled priority review execution-window acquisition, duplicate scheduled-window skipping, and manual-trigger bypass of the scheduled lock.
- Validate manual priority review notification override grouping and repeated same-day delivery when a default or previous manual notification event already exists.
- Validate scheduler next-window calculation, cancellation, structured failure recovery, and scheduled review command payloads.

## 4. HTTP tests
- Verify login succeeds for seeded seller and Operations accounts.
- Verify login fails for invalid email or password.
- Verify question creation succeeds for a valid order, buyer, body, and optional order item.
- Verify backend-only question creation fails for empty body, missing order, missing buyer, invalid order item, and an existing open buyer/order question by coarse `status`.
- Verify question creation returns priority classification metadata.
- Verify seller orders endpoint filters by status, buyer text, and date range.
- Verify order detail endpoint includes items and questions.
- Verify reply endpoint accepts open questions and rejects resolved questions.
- Verify reply endpoint rejects or hides questions for orders assigned to another seller.
- Verify Operations queue endpoint excludes resolved questions and returns priority fields.
- Verify Operations queue filters by seller search, buyer text, order status, date range, and priority classification.
- Verify priority review endpoint recategorizes only eligible open questions.
- Verify priority review endpoint requires an Operations access token and rejects missing or wrong-role tokens.
- Verify manual priority review is not blocked by an already acquired scheduled execution window.
- Verify repeated manual priority review can log and deliver same-day notification batches through manual override grouping even when no additional questions are recategorized.
- Verify `POST /graphql` read queries return role-shaped read models and do not mutate state.
- Verify `POST /graphql` rejects missing, malformed, expired, or invalid-signature JWTs.
- Verify `POST /graphql` rejects mutation operations.
- Verify buyer order list and detail reads require authenticated Buyer ownership.
- Verify buyer order list and detail reads reject missing tokens, wrong roles, mismatched `buyerId` values, and non-owned orders.
- Verify buyer-created question commands reject non-owned orders, invalid item references, and duplicate active or unclosed questions.
- Verify buyer answer approval requires the owning buyer and a seller-answered question, and rejects missing replies, already-approved questions, and already-closed questions.
- Verify Operations seller search behavior by name, email, and id through REST and GraphQL read paths.

## 5. Frontend tests
- Verify login form success and validation states.
- Verify protected routes redirect unauthenticated users to the login page.
- Verify role-specific UI visibility for Seller and Operations accounts.
- Verify logout clears authenticated state and returns to login.
- Verify order list rendering, filters, and empty states.
- Verify order questions rendering and product references.
- Verify seller reply success and validation states.
- Verify Operations section rendering inside the dashboard.
- Verify Operations filter controls.
- Verify skeleton placeholders for login, order list, order detail, questions, replies, and Operations queue.
- Verify GraphQL-driven seller dashboard and Operations read states.
- Verify GraphQL requests include the current session JWT as a Bearer token.
- Verify Buyer role route protection and role-specific navigation.
- Verify buyer order list, buyer order detail, buyer question creation, duplicate active-question conflicts, seller-answer review, buyer approval, and approve-answer conflict UI states.
- Verify question lifecycle labels appear only when questions exist.
- Verify Operations seller search input submission and filtered results for seller name, email, and id.
- Verify Operations manual priority review trigger loading, success, execution status, notification batch summary, and recoverable error states.

## 6. E2E smoke tests
- Start with seeded data.
- Log in as a seeded seller.
- Open the dashboard as a seeded seller.
- Filter orders and inspect an order with questions.
- Create a question through the backend endpoint and confirm it appears in the seller read model.
- Reply to an open question and confirm it moves to `seller_answered` until buyer approval.
- Log out and log in as a seeded Operations account.
- Open the Operations section and confirm resolved questions are excluded.
- Trigger the Operations manual priority review from seeded initial data and confirm the REST command succeeds with completed execution feedback and a positive notification batch count from the email-only seeded channels.
- Log in as a seeded Buyer account, open buyer-owned orders, create a question, log in as seller, answer it, log back in as buyer, and approve the answer.
- Log in as Operations and apply the seller search input by seller name, email, or id to confirm filtered priority queue results through the visible Operations UI.
- Trigger notification delivery and confirm the MVP console channel driver receives the delivery payload in backend tests, including manual override resend payloads.

## 7. Final verification
- TASK-017 final verification runs `corepack pnpm run test` and `corepack pnpm run build` from `frontend/`.
- Because TASK-011 through TASK-014 introduced Go changes, TASK-017 final verification runs `go test ./...` and `go vet ./...` from `backend/`.
- Documentation diff review confirms buyer-facing flow, notification drivers, and Operations seller filter guidance are traceable to completed tasks.