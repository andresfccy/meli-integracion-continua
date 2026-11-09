# Data Model

## 1. Purpose
Define the shared MVP domain objects, relationships, and invariants used by the seller orders, buyer questions, priority queue, notifications, API, frontend, and tests.

## 2. Entities

### Seller
- `id`: required stable identifier.
- `name`: required display name.
- `email`: required notification email address.
- `notificationChannels`: required collection of notification channel identifiers for the seller notifiable target.

### Buyer
- `id`: required stable identifier.
- `name`: required display name.
- `email`: required contact email.
- `notificationChannels`: planned collection of notification channel identifiers for buyer-facing notifications.

### Product
- `id`: required stable identifier.
- `sellerId`: required seller owner identifier.
- `name`: required display name.
- `category`: optional category used for display and future prioritization.
- `currentPrice`: optional current catalog price; order item pricing must not depend on this value.

### Order
- `id`: required stable identifier.
- `sellerId`: required seller owner identifier.
- `buyerId`: required buyer identifier.
- `status`: required lifecycle value: `new`, `processing`, `shipped`, `delivered`, or `cancelled`.
- `orderedAt`: required order timestamp.
- `items`: required non-empty collection of order items.
- `total`: derived sum of item quantity multiplied by purchase-time unit price.

### OrderItem
- `id`: required stable identifier.
- `orderId`: required order identifier.
- `productId`: required product identifier.
- `productName`: required purchase-time display name.
- `quantity`: required positive integer.
- `unitPrice`: required non-negative purchase-time unit price.

### Question
- `id`: required stable identifier.
- `orderId`: required order identifier.
- `buyerId`: required buyer identifier.
- `orderItemId`: optional order item identifier when the question concerns a specific item.
- `body`: required non-empty text.
- `status`: current coarse resolution state: `open` or `resolved`.
- `lifecycleStatus`: buyer-facing question lifecycle value used alongside `status`; values are question-specific and include `question_active`, `seller_answered`, `buyer_approved`, and `closed`. Buyer-created questions use `question_active`, seller replies move questions to `seller_answered` while `status` remains `open`, and buyer approval closes the lifecycle with `status=resolved` and terminal visible `lifecycleStatus=closed`.
- `buyerApprovedAt`: optional timestamp set when the owning buyer approves the seller answer.
- `createdAt`: required creation timestamp.
- `resolvedAt`: optional timestamp set when status is `resolved`.
- `reply`: optional seller reply for MVP.
- `priorityScore`: required integer set at creation and updated during recategorization.
- `priorityClass`: required value: `Low`, `Medium`, `High Priority`, or `Critical`.
- `matchedUrgencyWords`: required collection of urgency dictionary words or phrases matched during the latest categorization.
- `priorityDictionaryVersion`: required version used for latest categorization; MVP value is `mvp-2026-07-13`.
- `lastCategorizedAt`: required timestamp for latest creation-time categorization or recategorization.

### Reply
- `id`: required stable identifier.
- `questionId`: required question identifier.
- `sellerId`: required seller author identifier.
- `body`: required non-empty text.
- `createdAt`: required reply timestamp.
- Accepted seller replies move the question lifecycle to `seller_answered` while `status` remains `open`; no configurable reply-resolution flag is modeled.
- Buyer approval closes the lifecycle and sets the terminal visible state to `closed`.

### UserRole
- `userId`: required seeded user identifier.
- `role`: required current value: `seller`, `operations`, or `buyer`.
- `sellerId`: optional seller identifier when the role is `seller`.
- `buyerId`: required buyer identifier when the role is `buyer`.

### SeededUserAccount
- `id`: required stable identifier.
- `email`: required login email.
- `password`: required demo credential validated only against the in-memory seed dataset.
- Seeded demo passwords may be displayed by the frontend challenge login view as static hints, but safe role discovery responses must not include password fields.
- `displayName`: required display name.
- `roles`: required non-empty collection of `seller`, `operations`, and/or `buyer` role values.
- `sellerId`: required when the account has the `seller` role.
- `opsTeamMemberId`: required when the account has the `operations` role.
- `buyerId`: required value when the account has the `buyer` role.

### OpsTeamMember
- `id`: required stable identifier.
- `name`: required display name.
- `email`: required login email.
- `notificationChannels`: required collection of notification channel identifiers for the Operations notifiable target.

### NotificationChannel
- `id`: required stable identifier.
- `type`: required value; MVP value is `email`, future values may include `slack` or `sms`.
- MVP seed data uses enabled `email` channels only. Future `slack` or `sms` channels may be added behind the driver boundary.
- `destination`: required channel destination such as an email address, Slack identifier, or phone number.
- `enabled`: required boolean.
- `label`: optional display label.

### NotificationChannelDriver
- Outbound delivery abstraction selected by notification channel type.
- Drivers own channel-specific delivery protocols.
- The MVP console driver writes enabled channel payloads to stdout.
- Domain entities and application services must not depend on stdout, email, Slack, SMS, or provider-specific protocols.

### Notifiable
- `id`: required stable identifier.
- `objectType`: required value identifying the notifiable entity type, such as `seller`, `opsTeamMember`, `buyer`, or `merchant`.
- `objectId`: required identifier of the notifiable entity.
- `notificationChannels`: required collection of notification channel identifiers.

### NotificationEvent
- `id`: required stable identifier.
- `questionIds`: required non-empty collection of included question identifiers.
- `notifiableId`: required notified target identifier.
- `channelId`: required notification channel identifier.
- `batchDate`: required review date used to group notification events.
- `groupingKey`: required deterministic key composed from notifiable target, channel, and batch date. Manual Operations resend events retain that default prefix and add a manual override suffix.
- `includedPriorityClasses`: required collection of triggering classes included in the batch; values are `High Priority` and/or `Critical`.
- `message`: required logged notification message.
- `createdAt`: required event timestamp.

## 3. Relationships

| Source | Relationship | Target | Cardinality |
| ------ | ------ | ------ | ------ |
| Seller | owns | Product | One seller to many products |
| Seller | receives | Order | One seller to many orders |
| Buyer | places | Order | One buyer to many orders |
| SeededUserAccount | authenticates as | Buyer | Zero or one buyer identity per seeded account |
| Order | contains | OrderItem | One order to one or many order items |
| OrderItem | references | Product | Many order items to one product |
| Order | has | Question | One order to zero or many questions |
| Question | may reference | OrderItem | Zero or one order item per question |
| Question | may have | Reply | Zero or one reply per question in MVP |
| NotificationEvent | includes | Question | One batched notification event to one or many questions |
| UserRole | grants access for | Seller or Operations role | One seeded user to one or many roles |
| SeededUserAccount | authenticates as | Seller | Zero or one seller identity per seeded account |
| SeededUserAccount | authenticates as | OpsTeamMember | Zero or one Operations identity per seeded account |
| Seller | participates in | Notifiable | One seller to one notifiable target |
| OpsTeamMember | participates in | Notifiable | One Operations team member to one notifiable target |
| Buyer | participates in | Notifiable | Planned one buyer to one notifiable target |
| Notifiable | uses | NotificationChannel | One notifiable target to one or many channels |
| NotificationEvent | targets | Notifiable | Many notification events to one notifiable target |
| NotificationChannelDriver | delivers through | NotificationChannel | One driver to one or many channels of the same supported type |

## 4. Invariants

- An order must contain at least one order item.
- An order item quantity must be greater than zero.
- An order item unit price must be greater than or equal to zero.
- An order item must preserve purchase-time product name and unit price.
- A question must belong to exactly one order.
- A question body must not be empty.
- A buyer must not have more than one open question for the same order.
- A question order item reference is valid only when the order item belongs to the same order.
- Only open questions may receive a seller reply.
- A seller reply body must not be empty.
- An accepted seller reply moves the question lifecycle to `seller_answered` while preserving unresolved `open` status until buyer approval.
- Buyer-facing behavior must keep order status separate from question lifecycle status.
- Question lifecycle states must not include order-like states such as `created` or `accepted`.
- Question lifecycle labels must be visible to authorized users only when a question exists; orders without questions show only order status.
- Seller replies must move the question lifecycle to `seller_answered` until buyer approval closes the question.
- Buyer approval must record approval metadata and make `closed` the terminal visible lifecycle according to SPEC-002 rules. `buyer_approved` may be persisted as transition metadata or an audit-facing state, but open/closed UI behavior treats approved questions as closed.
- Duplicate detection rejects a new buyer-created question when the same buyer already has an active or unclosed question for the same order.
- The one-question-to-zero-or-one-seller-reply model remains in force for the buyer-facing lifecycle unless a later spec explicitly expands it.
- A resolved question must be excluded from the Operations priority queue.
- Buyer-facing question creation UI is available for authenticated buyer-owned order detail; backend question creation remains available for MVP testing and external-ingestion simulation.
- Each created question must be classified at creation time and store priority metadata.
- Daily recategorization may update priority metadata only for open questions whose last categorization is at least 2 days old.
- A notifiable target must use enabled notification channels for notification delivery simulation.
- Batched notification events must include at least one question identifier and one notification channel.
- Notification batching must group by notifiable target, channel, and review date; High Priority and Critical classifications trigger notification execution but do not split the batch grouping.
- Scheduled priority review notification events must remain duplicate-safe for the same notifiable target, channel, and UTC review date. Manual Operations priority review may create manual override grouping keys for the same review date to support explicit same-day resend.
- Seed data must include at least one Seller role, one Operations role, and one Buyer role.
- Authentication must validate email and password only against seeded user accounts.
- Seeded user accounts must not support registration, password recovery, password reset, email verification, social login, OAuth, refresh tokens, or external identity providers.
- Frontend authenticated state must not persist beyond the current browser session.
- Role-specific UI components must be shown only to authenticated users with the related seeded role.
- Notification delivery must use channel drivers selected by channel type.
- Disabled or unsupported notification channels must be handled deterministically without changing priority classification behavior.
- Application services must not write notification payloads directly to stdout; stdout output belongs to the MVP console channel driver.

## 5. Post-MVP Integration Assumptions

- The MVP assumes orders, products, and order items are available in the same backend context.
- Future integrations may move orders, products, or order items behind external systems or team-owned services.
- Threaded buyer/seller conversations would require additional `ConversationThread` and `Message` structures outside the MVP model.
- Future notification integrations may add Slack, SMS, or other channel drivers behind the same outbound driver selection boundary.
