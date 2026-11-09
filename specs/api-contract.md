# API Contract

## 1. Conventions
- Style: RESTful HTTP JSON API.
- Request and response bodies use `application/json`.
- Timestamps use ISO 8601 strings.
- Money values use decimal numbers in the seed currency.
- Collections are returned in deterministic order.
- Error responses use a shared shape: `{"error":{"code":"string","message":"string","details":{}}}`.

## 2. Status codes
- `200 OK`: successful read or update.
- `201 Created`: successful creation of a reply, buyer question, or notification event when exposed.
- `400 Bad Request`: invalid filter, malformed request body, or validation error.
- `401 Unauthorized`: credentials are missing or invalid for authentication.
- `403 Forbidden`: authenticated user has the wrong role or does not match the protected buyer, seller, or Operations context.
- `404 Not Found`: seller, order, or question does not exist.
- `409 Conflict`: action conflicts with lifecycle state, such as replying to a resolved question.
- `500 Internal Server Error`: unexpected server error.
- Each endpoint below states its success status explicitly when the general status-code guidance could be ambiguous.

## 3. Endpoints

### POST `/api/auth/login`
Purpose: validate MVP login credentials against the in-memory seeded Seller, Operations, and Buyer accounts.

Success status: `200 OK`.

Request body:
- `email`: required login email.
- `password`: required login password.

Success response:
- `user`: seeded account summary with display name, roles, optional seller identifier, optional Operations team member identifier, and optional buyer identifier.
- `accessToken`: JWT access token for authenticated GraphQL reads, authenticated buyer order reads, and authenticated buyer question creation.

Error behavior:
- `400 Bad Request` for malformed request body.
- `401 Unauthorized` for unknown email or invalid password.

Notes:
- Registration, password recovery, password reset, email verification, social login, OAuth, refresh tokens, and external identity providers are out of scope.
- Logout is handled by the frontend by clearing the current browser-session authenticated state.
- Backend authentication infrastructure is limited to seeded credential validation and short-lived JWT access-token validation for GraphQL reads, authenticated buyer order reads, and authenticated buyer question creation required by the MVP.

### GET `/api/sellers/{sellerId}/orders`
Purpose: return recent orders belonging to one seller.

Success status: `200 OK`.

Query parameters:
- `status`: optional order status filter.
- `buyer`: optional case-insensitive buyer text search.
- `dateFrom`: optional inclusive order date lower bound.
- `dateTo`: optional inclusive order date upper bound.

Success response:
- `orders`: array of orders with buyer summary, status, ordered date, item summaries, and total.
- Sorting: order date descending, then order identifier ascending.

Error behavior:
- `400 Bad Request` for unsupported status or invalid date range.
- `404 Not Found` when the seller is unknown.

### GET `/api/sellers/{sellerId}/orders/{orderId}`
Purpose: return order detail including items and buyer questions.

Success status: `200 OK`.

Success response:
- `order`: order fields, buyer summary, items, total, and questions.
- Questions include body, `status`, `lifecycleStatus`, created timestamp, buyer, optional order item/product reference, and optional reply.

Error behavior:
- `404 Not Found` when the seller or order is unknown, or the order does not belong to the seller.

### POST `/api/sellers/{sellerId}/questions/{questionId}/reply`
Purpose: allow a seller to reply to an open question.

Success status: `201 Created` because the request creates a seller reply and returns the updated question state.

Request body:
- `body`: required non-empty reply text.
- Accepted seller replies move the question lifecycle to `seller_answered` and keep the question unresolved until buyer approval; API consumers cannot choose to close the question from the seller reply command.

Success response:
- `question`: updated question including reply, `status=open`, and `lifecycleStatus=seller_answered`.

Error behavior:
- `400 Bad Request` for empty reply body.
- `404 Not Found` when the seller or question is unknown, or the question is not attached to an order assigned to that seller.
- `409 Conflict` when the question is already answered, already approved, or already closed.

### POST `/api/orders/{orderId}/questions`
Purpose: create a buyer question directly through the backend for MVP testing and external-ingestion simulation. This endpoint does not imply a buyer-facing UI.

Success status: `201 Created`.

Request body:
- `buyerId`: required buyer identifier.
- `body`: required non-empty question text.
- `orderItemId`: optional order item identifier when the question concerns a specific item in the order.

Success response:
- `question`: created open question with body, buyer, order, optional order item reference, priority score, priority classification, matched urgency words, and priority dictionary version.
- `notification`: optional metadata indicating whether the created High Priority or Critical question is a candidate for the next daily 00:00 notification review.

Error behavior:
- `400 Bad Request` for empty body or order item not belonging to the order.
- `404 Not Found` when the order or buyer is unknown.
- `409 Conflict` when the buyer already has an open question for the same order; this backend-only endpoint uses the current coarse `status` semantics.

Side effects:
- The question is classified at creation time using the approved priority dictionary.
- High Priority or Critical classification marks the question as notification-eligible for the next daily 00:00 review; creation does not send notification messages immediately.

### GET `/api/operations/questions`
Purpose: return unresolved questions sorted by priority for Operations.

Success status: `200 OK`.

Query parameters:
- `seller`: optional case-insensitive seller search filter matching seller identifier, seller name, or seller email from one input.
- `buyer`: optional case-insensitive buyer text filter.
- `status`: optional order status filter.
- `dateFrom`: optional inclusive order date lower bound.
- `dateTo`: optional inclusive order date upper bound.
- `priorityClass`: optional priority classification filter.

Success response:
- `questions`: array of unresolved questions with order summary, buyer summary, seller summary, waiting time, order value, matched urgency words, priority score, and priority classification.
- Sorting: priority score descending, question creation timestamp ascending, order value descending, question identifier ascending.

Rules:
- The seller filter matches seller identifier, seller name, or seller email using one search value.

Error behavior:
- `500 Internal Server Error` for unexpected priority evaluation failures.

### POST `/api/operations/priority-review`
Purpose: allow an authenticated Operations user to trigger the MVP daily recategorization review manually.

Success status: `200 OK`.

Request body:
- `reviewedAt`: optional ISO 8601 timestamp; defaults to server time.

Authentication:
- Requires a valid Bearer token for a seeded account with the Operations role.

Success response:
- `review`: review timestamp, review window, execution identifier, execution status, number of eligible open questions, number recategorized, and number of notification batches logged.

Rules:
- Only open questions whose latest categorization or recategorization is at least 2 days old are recategorized.
- Resolved questions and recently categorized open questions are skipped.
- Notification sending happens only during this 00:00 daily review flow or a manual invocation of the same review behavior.
- The review checks newly created and recategorized questions, skips resolved or non-qualifying questions before sending, batches qualifying High Priority and Critical notifications by seller or notifiable user plus assigned enabled channel, and sends one batched message per grouping.
- Scheduled review triggers acquire one execution per UTC review window before running. Duplicate scheduled triggers return `already_executed` and do not send notifications for that scheduled window. Manual Operations triggers are not blocked by this scheduled execution-window acquisition.
- Manual Operations triggers intentionally override notification delivery deduplication for seller batches. A repeated same-day manual call can return `notificationBatchesLogged` greater than zero even when no questions were recategorized, because the event uses a manual override grouping key and delivery is delegated again.

Error behavior:
- `401 Unauthorized` when the access token is missing or invalid.
- `403 Forbidden` when the authenticated user does not have the Operations role.

### GET `/api/roles/seeded-users`
Purpose: expose seeded role options needed by the MVP dashboard experience.

Success status: `200 OK`.

Success response:
- `users`: array of seeded users with role values `seller`, `operations`, and/or `buyer` and the matching optional seller, Operations team member, or buyer identifier.

Notes:
- This endpoint must not expose seeded passwords.
- The frontend may show static demo credential hints for the in-memory challenge environment, but those hints must not come from this safe role-options response.
- Once the login flow is available, frontend navigation should rely on `POST /api/auth/login` for authenticated user state.

### GET `/api/buyers/{buyerId}/orders`
Purpose: return orders belonging to the authenticated seeded Buyer account.

Success status: `200 OK`.

Query parameters:
- `status`: optional order status filter.
- `seller`: optional seller text filter.
- `dateFrom`: optional inclusive order date lower bound.
- `dateTo`: optional inclusive order date upper bound.

Success response:
- `orders`: array of buyer-owned orders with seller summary, status, ordered date, item summaries, total, and question lifecycle summaries only when questions exist.
- Question lifecycle labels belong to question summaries. An order may include aggregate question summary counts for display, but orders without questions must not include a synthetic order-level lifecycle label.

Rules:
- The authenticated user must have the Buyer role and must match `buyerId`.
- Orders without questions return only normal order status and no question lifecycle label.

Error behavior:
- `401 Unauthorized` when the access token is missing or invalid.
- `403 Forbidden` when the authenticated user is not a Buyer or the authenticated buyer identity does not match `buyerId`.
- `400 Bad Request` for unsupported status or invalid date range.

### GET `/api/buyers/{buyerId}/orders/{orderId}`
Purpose: return buyer-owned order detail, items, and questions.

Success status: `200 OK`.

Success response:
- `order`: buyer-owned order detail including seller, buyer, items, total, questions, seller replies, and question lifecycle status on each existing question.
- Orders with no questions include the normal order status and an empty questions collection; they do not include a question lifecycle label.

Rules:
- The authenticated Buyer account may read only orders belonging to that buyer.

Error behavior:
- `401 Unauthorized` when the access token is missing or invalid.
- `403 Forbidden` when the authenticated user is not a Buyer or the authenticated buyer identity does not match `buyerId`.
- `404 Not Found` when the order is unknown or is not owned by the authenticated buyer.

### POST `/api/buyers/{buyerId}/orders/{orderId}/questions`
Purpose: allow an authenticated buyer to create a question from their own order detail view.

Success status: `201 Created`.

Request body:
- `body`: required non-empty question text.
- `orderItemId`: optional order item identifier when the question concerns a specific item in the order.

Success response:
- `question`: created open question with `lifecycleStatus` set to `question_active` and priority metadata.

Rules:
- The authenticated Buyer account must own the order.
- Order-like states such as `created` and `accepted` are not valid question lifecycle states.
- Duplicate active or unclosed buyer questions for the same buyer and order remain invalid unless a later spec changes that rule.

Error behavior:
- `400 Bad Request` for empty body or order item not belonging to the order.
- `401 Unauthorized` when the access token is missing or invalid.
- `403 Forbidden` when the authenticated user is not a Buyer, the authenticated buyer identity does not match `buyerId`, or the order is not owned by that buyer.
- `404 Not Found` when the order is unknown.
- `409 Conflict` when the buyer already has an active or unclosed question for the same order.

### POST `/api/buyers/{buyerId}/questions/{questionId}/approve-answer`
Purpose: allow the buyer who owns the question to approve the seller answer and close the question lifecycle.

Success status: `200 OK`.

Success response:
- `question`: updated question with buyer approval metadata and closed lifecycle state.

Rules:
- The question must belong to an order owned by the authenticated buyer.
- The question must already have a seller reply and be in the `seller_answered` lifecycle state.
- Approval records buyer approval metadata and returns a terminal closed lifecycle display.

Error behavior:
- `401 Unauthorized` when the access token is missing or invalid.
- `403 Forbidden` when the authenticated user is not a Buyer, the authenticated buyer identity does not match `buyerId`, or the question is not owned by that buyer.
- `404 Not Found` when the question is unknown.
- `409 Conflict` when the question has no seller reply, is not in `seller_answered`, is already closed, or has already been approved.

## 4. Authentication and authorization behavior
- Current authentication is based exclusively on seeded Seller, Operations, and Buyer accounts.
- Buyer authentication uses the same seeded account model without registration, refresh tokens, OAuth, durable sessions, or external identity providers.
- Credentials are validated only against the in-memory seed dataset.
- Protected API consumers must use the authenticated seeded user context returned by login to choose role-appropriate requests.
- The MVP does not require backend session persistence beyond validating login credentials.
- The frontend is responsible for redirecting unauthenticated users to the login page and hiding role-inappropriate UI components.

## 5. Notification behavior
- Priority notification eligibility is triggered by priority evaluation when a question is classified as High Priority or Critical.
- Notification messages are not sent immediately during question creation; sending occurs during the daily 00:00 review flow after recategorization checks are applied.
- The daily review skips resolved and non-qualifying questions, then batches qualifying notifications by seller or notifiable user plus assigned enabled channel.
- The MVP email notification is simulated by logging one batched notification event per grouping. Current MVP seed data uses enabled email channels only.
- Medium and Low questions must not trigger notification events.
- Future Slack and SMS channels must not require API consumers to change priority queue behavior.
- Notification targets are resolved through notifiable objects and their enabled notification channels.
- Notification events are batched by notifiable target, channel, and review date to avoid spam; High Priority and Critical classifications remain the trigger for notification eligibility. Manual Operations review may append a manual override suffix to the grouping key to resend a same-day seller batch explicitly.
- Notification delivery separates event creation from delivery. Event creation records notification events; outbound channel drivers deliver enabled channel payloads.
- Channel drivers are selected by channel type. The MVP console driver writes delivered payloads to stdout, while future Slack or SMS drivers can be added without changing priority classification behavior.
- Driver delivery is triggered by manual or scheduled priority review for seller notification events. Disabled channels are skipped without payload output, and unsupported enabled channel types produce deterministic delivery errors at the driver boundary.
- Buyer notification delivery remains deferred until buyer notification rules explicitly define targets and triggers.
- Application services must call notification delivery ports and must not write channel payloads directly to stdout.

## 6. GraphQL read composition

### HTTP route
- `POST /graphql`: accepts authenticated GraphQL read queries for frontend dashboard composition.
- Requests must include `Authorization: Bearer <accessToken>` using the JWT returned by `POST /api/auth/login`.

Success status: `200 OK` for successful read query execution.

Authentication errors use `401 Unauthorized` with the shared error shape when the JWT is missing, malformed, expired, signed with an unsupported algorithm, or does not map to a seeded account.

JWT behavior:
- Signing algorithm: HS256 for the MVP seed environment.
- Required claims: subject seeded account identifier, issuer `seller-dashboard-mvp`, audience `graphql-read`, issued-at, and expiration.
- The JWT does not introduce refresh tokens, durable backend sessions, OAuth, social login, or external identity-provider behavior.

GraphQL is approved only as a frontend-facing read composition layer for authenticated role-shaped rendering. GraphQL resolvers must compose existing backend services into UI-ready read models and must not own domain or business rules.

### Scope
- In scope: authenticated role-shaped reads for seller dashboard, buyer order reads, and Operations priority queue screens.
- Out of scope: mutations, subscriptions, federation, schema stitching, public/external consumers, business logic ownership, and replacement of REST command endpoints.

### Initial queries
- `viewer`: returns the authenticated seeded account summary and roles needed by the frontend.
- `sellerDashboard(sellerId, filters)`: returns seller order list data shaped for the seller dashboard.
- `sellerOrderDetail(sellerId, orderId)`: returns order detail, order items, questions, replies, and priority metadata for the selected order.
- `operationsPriorityQueue(filters)`: returns unresolved priority queue data shaped for Operations, including seller, buyer, order, priority, and filter metadata.
- `buyerDashboard(buyerId, filters)`: returns buyer-owned order list data shaped for the buyer dashboard.
- `buyerOrderDetail(buyerId, orderId)`: returns buyer-owned order detail, questions, seller replies, and question lifecycle labels when questions exist.

### Command boundary
- REST remains the MVP command API for login, backend question creation, seller replies, recategorization triggers, and all state-changing behavior.
- GraphQL mutations are rejected.