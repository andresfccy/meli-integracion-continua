# SPEC-003: Priority and Notifications
## 1. Status
- State: Refined
- Last updated: 2026-07-14
- Related tasks:
- TASK-001
- TASK-004
- TASK-005
- TASK-007

## 2. Objective
Help Operations identify the most important unresolved buyer questions and notify sellers when urgency crosses a threshold.

## 3. Actors
- Operations team
- Seller

## 4. User Stories

### US-PRI-001
As an operations team member
I want to review unresolved questions ordered by importance.

### US-PRI-002
As a seller
I want to receive notifications for high-priority questions

### US-PRI-003
As an operations team member
I want to access the priority queue inside the existing dashboard, so that I do not need a separate application.

## 5. Scope
### In scope
- Deterministic priority score.
- Priority classification.
- Unresolved questions only.
- Stable ordering and tie-break rules.
- Simulate email notifications.
- MVP seed data uses enabled email notification channels only.
- Dedicated Operations section inside the existing dashboard.
- Seeded application roles for Seller and Operations.
- Future notification channel extensibility for Slack and SMS.
- Classification at question creation.
- Daily 00:00 recategorization checks for open questions that reached the 2-day waiting-time threshold.
- Notification batching to prevent seller notification spam.
- Operations filters by seller search, buyer text, order status, date range, and priority classification.
- Notification channel driver delivery selected by channel type.
- Buyer notification behavior is deferred until buyer notification rules are explicitly approved.

### Out of scope
- Real providers.
- Notification retries across distributed systems.
- Delivery tracking.
- Separate Operations application.
- Buyer-facing notifications.
- Operations-configurable notification thresholds.
- Reminder cadence configuration.
- Channel-provider implementation details inside domain or application services.

## 6. Functional requirements
### FR-PRI-001
The system shall return unresolved questions for Operations sorted by priority.

### FR-PRI-002
The system shall calculate priority using at least waiting time, order value, and urgency words from a simple dictionary.

### FR-PRI-003
The system shall classify unresolved questions as Low, Medium, High Priority, or Critical.

### FR-PRI-004
The system shall use deterministic tie-breakers when two unresolved questions have the same priority score.

### FR-PRI-005
The system shall expose the Operations priority queue as a dedicated section inside the existing dashboard.

### FR-PRI-006
The system shall support seeded application roles including Seller and Operations.

### FR-PRI-007
The system shall exclude resolved questions from the Operations priority queue.

### FR-PRI-008
The system shall classify a question at creation time using the approved urgency dictionary.

### FR-PRI-009
The system shall run a daily 00:00 UTC recategorization check through an inbound scheduler adapter for open questions that have reached at least 2 days since their last categorization or recategorization.

### FR-PRI-010
The system shall allow Operations to filter the priority queue by seller search, buyer text, order status, date range, and priority classification.

Seller search must use one input that matches seller name, seller email, or seller identifier case-insensitively where applicable.

### FR-NOT-001
The system shall trigger a notification when a question is classified as High Priority.

### FR-NOT-002
The system shall trigger a notification when a question is classified as Critical.

### FR-NOT-003
The system shall simulate the current email notification by logging the notification event.

### FR-NOT-004
The notification design shall allow future Slack and SMS channels without changing priority classification behavior.

### FR-NOT-005
The system shall not trigger priority notifications for Medium or Low questions.

### FR-NOT-006
The system shall batch qualifying High Priority and Critical notification events by notifiable target, channel, and review date while preserving High Priority and Critical classifications as notification triggers.

### FR-NOT-007
The system shall deliver notification events through outbound notification channel drivers selected by channel type.

### FR-NOT-008
The MVP console driver shall write delivered notification payloads to stdout for enabled channels without application services writing directly to stdout.

### FR-NOT-009
The notification driver boundary shall support future Slack, SMS, removed channels, disabled channels, and additional channel protocols without changing priority classification behavior.

### FR-NOT-010
The buyer-facing flow shall not notify buyer notifiable targets until finalized buyer notification rules explicitly require it.

### FR-NOT-011
The Operations manual priority review trigger shall allow same-day seller notification batches to be sent again through explicit manual override grouping, while scheduled review triggers remain duplicate-safe for one UTC review window.

## 7. Business rules
### BR-PRI-001
A question is eligible for the Operations priority queue only when its status is open.

When buyer-facing lifecycle states are active, `question_active` and `seller_answered` questions remain unresolved for Operations until buyer approval records closure. `buyer_approved` and terminal `closed` questions must not appear in the unresolved Operations queue.

### BR-PRI-002
Priority score must be deterministic for the same seeded data and current evaluation time.

### BR-PRI-003
Waiting time must increase importance as the question remains unresolved.

### BR-PRI-004
Higher order value must increase importance.

### BR-PRI-005
Urgency words must increase importance when the question body contains a word from the approved dictionary.

### BR-PRI-005A
The approved Critical words are `fraud`, `scam`, `chargeback`, `legal`, `lawsuit`, `stolen`, `unsafe`, `dangerous`, `injury`, `broken`, `missing`, `not delivered`, `never arrived`, `urgent refund`, and `cancel immediately`.

### BR-PRI-005B
The approved High Priority words are `urgent`, `refund`, `damaged`, `defective`, `wrong item`, `late`, `delayed`, `where is my order`, `replacement`, `complaint`, `expedite`, `address issue`, and `payment issue`.

### BR-PRI-005C
Urgency dictionary matching must be case-insensitive, phrase matches must be evaluated before single-word matches, and dictionary version `mvp-2026-07-13` must be recorded with the classification result.

### BR-PRI-005D
A Critical dictionary match sets the keyword signal to Critical; otherwise a High Priority dictionary match sets the keyword signal to High Priority; otherwise keywords do not increase classification by themselves.

### BR-PRI-006
Priority classifications use these thresholds: Critical is score greater than or equal to 80; High Priority is score greater than or equal to 60 and less than 80; Medium is score greater than or equal to 30 and less than 60; Low is score less than 30.

### BR-PRI-007
Tie-breakers must sort by priority score descending, question creation timestamp ascending, order value descending, and question identifier ascending.

### BR-PRI-008
Creation-time classification and daily recategorization must use the same deterministic priority policy.

### BR-PRI-009
Daily recategorization must skip resolved questions and open questions whose last categorization or recategorization is less than 2 days old.

Scheduled daily recategorization must acquire one execution per UTC review window before running. A duplicate scheduled trigger for the same window returns `already_executed` and must not recategorize questions or deliver notifications. The Operations manual trigger uses the same application review behavior but is not blocked by the scheduled execution-window acquisition.

### BR-NOT-001
Notifications are triggered only for High Priority and Critical classifications.

### BR-NOT-002
The MVP notification channel is simulated email through a logged notification event. Seeded enabled channels for the MVP are email-only; future Slack or SMS channels may be added behind the driver boundary without changing priority rules.

### BR-NOT-003
Notification delivery targets must be resolved from notifiable objects and their enabled notification channels.

### BR-NOT-004
A daily 00:00 review must group qualifying notification items into a batch per notifiable target, channel, and review date instead of emitting one seller notification per question. High Priority and Critical classifications trigger notification execution, but priority class must not split the batch grouping.

Default notification event idempotency remains based on review window, notifiable target, and channel. Scheduled review triggers must preserve this duplicate-safe event grouping. Operations manual review is the approved override path and may create manual override grouping keys for the same review date so sellers can receive a repeated same-day batch when Operations explicitly triggers it.

### BR-NOT-005
Notification event creation and notification delivery are separate responsibilities. Event creation records what should be delivered; channel drivers own how a channel payload is delivered.

### BR-NOT-006
Domain logic must not depend on stdout, email, Slack, SMS, or provider-specific protocols.

### BR-NOT-007
Application services may orchestrate notification delivery, but they must call outbound driver ports instead of writing channel payloads directly.

### BR-NOT-008
The MVP console driver output must include the notification target, channel type, channel identifier, question identifiers, priority classes, delivery trigger context, and timestamp.

### BR-NOT-009
Disabled channels must be skipped deterministically without producing delivery payloads. Unsupported enabled channel types must return deterministic delivery errors at the delivery boundary and must not change priority classification or notification event creation.

### BR-NOT-010
Driver delivery is scoped to seller notification events emitted by priority review. Buyer notification delivery remains deferred until a later approved rule defines buyer targets, triggers, and copy.

### BR-NOT-011
Manual override notification grouping keys must retain the default notifiable target, channel, and review-date grouping prefix and add a manual override identifier so repeated manual sends are auditable without changing scheduled idempotency.

### BR-NOT-012
Console driver output may be asserted in adapter-level tests, but domain and application tests should verify delivery through ports and payload values without depending on direct stdout writes from domain or application services.

## 8. Acceptance Scenarios

### AC-PRI-001: View unresolved priority queue
Given unresolved seeded questions exist
When Operations opens the priority section in the dashboard
Then the system displays unresolved questions sorted by priority

### AC-PRI-002: Exclude resolved questions
Given resolved and unresolved questions exist
When Operations views the priority queue
Then the system excludes resolved questions

### AC-PRI-003: Classify by score
Given unresolved questions have calculated priority scores
When the priority queue is returned
Then each question includes its priority score and classification

### AC-PRI-004: Apply deterministic tie-breakers
Given two unresolved questions have the same priority score
When Operations views the priority queue
Then the system applies the approved tie-breaker order

### AC-PRI-005: Operations section in dashboard
Given a seeded user has the Operations role
When the user opens the dashboard
Then the user can access a dedicated Operations priority section without opening a separate application

### AC-PRI-006: Classify question at creation
Given a valid question creation request contains urgency dictionary words
When the question is created
Then the system stores the initial priority score, classification, matched words, and dictionary version

### AC-PRI-007: Daily recategorization threshold
Given an open question was last categorized at least 2 days ago
When the daily 00:00 recategorization check runs
Then the system recalculates the question priority using the approved policy

### AC-PRI-008: Skip questions below recategorization threshold
Given an open question was categorized less than 2 days ago
When the daily 00:00 recategorization check runs
Then the system does not recategorize that question

### AC-PRI-009: Filter Operations queue
Given unresolved questions exist across seller names, seller emails, seller identifiers, buyers, statuses, dates, and priority classes
When Operations applies filters
Then the system returns only unresolved questions matching the selected filters in priority order

### AC-NOT-001: Trigger High Priority notification
Given an unresolved question is classified as High Priority
When the question is evaluated for priority
Then the system logs a simulated email notification event for the seller

### AC-NOT-002: Trigger Critical notification
Given an unresolved question is classified as Critical
When the question is evaluated for priority
Then the system logs a simulated email notification event for the seller

### AC-NOT-003: Do not notify for lower priority
Given an unresolved question is classified as Medium or Low
When the question is evaluated for priority
Then the system does not log a priority notification event

### AC-NOT-004: Batch daily notifications
Given multiple unresolved questions qualify for notification for the same notifiable target and channel on the same review date
When notification events are generated
Then the system logs one batched notification event containing the qualifying question identifiers

### AC-NOT-005: Deliver through channel driver
Given a notification event targets an enabled channel
When notification delivery is triggered
Then the system selects the channel driver by channel type and delegates delivery to that driver

### AC-NOT-006: Console driver writes payload
Given the MVP console notification driver receives a notification payload
When it delivers the payload
Then it writes the target, channel, question identifiers, priority classes, trigger context, and timestamp to stdout

### AC-NOT-007: Application service avoids direct stdout writes
Given notification delivery is triggered by priority review or buyer-facing notification behavior
When the application service orchestrates delivery
Then the service calls a notification driver port and does not write notification payloads directly to stdout

### AC-NOT-008: Disabled and unsupported channels
Given notification delivery includes disabled channels or enabled channels without registered drivers
When delivery is triggered
Then disabled channels are skipped and unsupported enabled channels produce deterministic delivery errors at the driver boundary

### AC-NOT-009: Manual review resends same-day batches
Given a qualifying seller notification batch already exists for a review date
When an authenticated Operations user manually triggers priority review again for that same review date
Then the system creates a manual override notification event and delegates delivery for the batch again without changing scheduled duplicate-window behavior

## 9. Edge Cases
- No unresolved questions exist.
- Two questions have the same priority score.
- Question body has no urgency words.
- Question body has multiple urgency words.
- Order value is zero.
- Question becomes resolved after previously appearing in the priority queue.
- Notification event is simulated and has no external delivery confirmation.
- A question contains both Critical and High Priority words; Critical takes precedence.
- Multiple qualifying questions target the same seller and channel on the same day.
- Multiple qualifying High Priority and Critical questions target the same notifiable target and channel on the same daily review.
- A notifiable target has no enabled notification channels.
- Daily recategorization runs when no questions have reached the 2-day threshold.
- A notification channel has no registered driver.
- A notification channel is disabled.
- A future Slack or SMS channel is added behind the driver registry.
- Seeded MVP notification channels are email-only even though future SMS and Slack types remain supported by the driver boundary.
- Operations manually triggers priority review after a same-day scheduled or manual notification batch already exists.

## 10. Success criteria
- Operations can see unresolved questions ordered by documented importance.
- Sellers receive simulated email notification events for High Priority and Critical questions.
- Priority behavior is explainable and testable.
- The dashboard can represent both Seller and Operations roles from seed data.
- Notification volume is batched by notifiable target to avoid spam while High Priority and Critical remain the send triggers.
- Priority classification is deterministic at creation and recategorization.
- Notification delivery remains extensible through channel drivers and does not leak provider protocols into domain rules.
- Scheduled review remains duplicate-safe per UTC review window, while Operations manual review can resend same-day seller notification batches through explicit manual override grouping.

## 11. Dependencies
- Buyer Questions specification.
- Shared data model.
- API contract.
- Frontend contract.
- Test plan.

## 12. Open Questions
- None for this MVP normalization pass.
