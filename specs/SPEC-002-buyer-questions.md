# SPEC-002: Buyer Questions

## 1. Status
- State: Refined
- Last updated: 2026-07-14
- Related tasks:
- TASK-001
- TASK-003A
- TASK-003
- TASK-006

## 2. Objective
Allow buyer questions to be associated with orders and allow sellers to review and answer them.

## 3. Actors
- Seller
- Buyer

## 4. User Stories

### US-QUE-001
As a seller
I want to view questions associated with an order, so that I can understand buyer concerns in context.

### US-QUE-002
As a seller
I want to reply to open questions, so that buyers receive an answer from the seller dashboard.

### US-QUE-003
As a seller
I want my reply to move the question into a seller-answered state, so that buyers can approve the answer before the question closes.

### US-QUE-004
As a seller
I want questions to reference a specific product when applicable, so that product-specific concerns are clear.

## 5. Scope
### In scope
- Question linked to an order.
- Optional link to a product contained in that order.
- Question body.
- Open/resolved coarse `status` semantics.
- Seller reply.
- Question timestamps.
- Seeded buyer questions.
- Backend-only question creation endpoint for MVP testing and external-ingestion simulation.
- Empty state for orders with no questions.
- Backend buyer-owned order reads for seeded Buyer accounts.
- Buyer-facing question creation UI for authenticated buyer-owned order detail.
- Buyer approval of seller answers before buyer-facing question closure.
- Question lifecycle labels visible to all authorized roles only when a question exists.
- Coexistence of current MVP `status` values and buyer-facing `lifecycleStatus`.

### Out of scope
- Buyer-facing question history UI beyond order-context questions.
- Real-time chat.
- Attachments.
- Editing sent questions.
- Multiple threaded replies.
- Threaded message lists and conversation history beyond one question and one seller reply.
- Order-level states such as `created` or `accepted` inside the question lifecycle.

## 6. Functional requirements
### FR-QUE-001
The system shall display buyer questions associated with each seller order.

The seller view must include only questions for orders assigned to the authenticated seller.

### FR-QUE-002
The system shall display each question's body, status, creation timestamp, linked order, buyer, and optional linked order item or product.

### FR-QUE-003
The system shall allow a seller to reply to an open question.

The seller may reply only when the question belongs to an order assigned to that seller.

### FR-QUE-004
The system shall move the question lifecycle to `seller_answered` and keep the question unresolved until buyer approval when a seller reply is accepted.

### FR-QUE-005
The system shall prevent seller replies to resolved questions.

### FR-QUE-006
The system shall provide a backend-only endpoint to create buyer questions for MVP testing and external-ingestion simulation.

### FR-QUE-007
The system shall display an empty state when an order has no questions.

### FR-QUE-008
The system shall reject creation of a buyer question when the same buyer already has an open question for the same order.

### FR-QUE-009
The system shall classify a created question at creation time using the approved priority policy.

### FR-QUE-010
The system shall allow an authenticated seeded Buyer account to view orders belonging to that buyer.

### FR-QUE-011
The buyer-facing system shall allow an authenticated buyer to create a question from one of their order detail views.

### FR-QUE-012
The buyer-facing system shall move a question to a seller-answered lifecycle state when a seller reply is accepted.

### FR-QUE-013
The buyer-facing system shall allow the buyer who owns the order to approve the seller answer and close the question lifecycle.

### FR-QUE-014
The frontend shall display question lifecycle labels for existing questions to authorized Seller, Buyer, and Operations users, and shall not display question lifecycle labels for orders without questions.

### FR-QUE-015
The system shall reject buyer-created questions when the buyer already has an active or unclosed question for the same order.

## 7. Business rules
### BR-QUE-001
Each question must belong to exactly one order.

### BR-QUE-002
A question may reference one order item only when that item belongs to the same order as the question.

Because each order is assigned to exactly one seller, seller question visibility and seller reply permissions are derived from the order's seller owner.

### BR-QUE-003
Only open questions may receive seller replies.

Only the seller assigned to the question's order may submit the seller reply.

### BR-QUE-004
A seller reply moves the question lifecycle to `seller_answered`; buyer approval closes the question lifecycle.

### BR-QUE-005
Buyer-facing question creation UI is available for authenticated buyer-owned order detail. Direct backend question creation remains available only for MVP testing and external-ingestion simulation.

### BR-QUE-006
A buyer must not have more than one open question for the same order.

### BR-QUE-007
A created question must start in open status.

### BR-QUE-008
Question creation must validate that the order exists, the buyer exists, the body is non-empty, and any order item reference belongs to the same order.

### BR-QUE-009
The MVP data structure is one question to zero or one seller reply; threaded messages would require future `ConversationThread` and `Message` entities and are deferred outside MVP.

### BR-QUE-010
Current completed buyer-facing behavior moves a question to `seller_answered` when the seller reply is accepted, then requires buyer approval before closure.

The current MVP `status` values remain `open` and `resolved`. `seller_answered` questions keep `status=open` so Operations can continue to treat them as unresolved until buyer approval sets the terminal closed lifecycle.

### BR-QUE-011
Question lifecycle states are question-specific and include `question_active`, `seller_answered`, `buyer_approved`, and `closed`.

`buyer_approved` records that the owning buyer approved the seller answer. The terminal visible lifecycle after approval is `closed`; implementations may persist `buyer_approved` as approval metadata, but user-facing open/closed behavior must treat the question as closed after approval.

### BR-QUE-012
Question lifecycle states must not include order-like states such as `created` or `accepted`.

### BR-QUE-013
Question lifecycle state is visible in the UI only when a question exists. Orders without questions show only normal order status.

### BR-QUE-014
Questions in `question_active` or `seller_answered` states remain unresolved for Operations visibility until buyer approval and closure rules mark them closed.

### BR-QUE-015
Buyer-facing duplicate detection treats `question_active` and `seller_answered` as active or unclosed questions for the same buyer and order. Closed buyer-approved questions do not block a later question unless a future spec changes that rule.

### BR-QUE-016
The one-question-to-zero-or-one-seller-reply model remains in force after buyer-facing lifecycle is added. Additional seller replies, buyer follow-up messages, and threaded conversations remain outside scope.

## 8. Acceptance Scenarios

### AC-QUE-001: View order questions
Given a seeded order has buyer questions
When the seller views the order questions section
Then the system displays each question with status, timestamp, buyer, body, and optional product reference

### AC-QUE-002: Reply marks question seller answered
Given an open question exists for a seller order
When the seller submits a valid reply
Then the system stores the reply, keeps `status=open`, and marks `lifecycleStatus=seller_answered`

### AC-QUE-004: Prevent reply to resolved question
Given a question is resolved
When the seller attempts to reply to the question
Then the system rejects the reply with a validation error

### AC-QUE-004A: Prevent reply by another seller
Given a question belongs to an order assigned to one seller
When a different seller attempts to view the order question or submit a seller reply
Then the system does not expose the order question to that seller and rejects the reply as not found or unauthorized according to the API contract

### AC-QUE-005: Order with no questions
Given an order has no buyer questions
When the seller views the order questions section
Then the system displays an empty state

### AC-QUE-006: Create question through backend endpoint
Given an order and buyer exist
When a valid question creation request is submitted to the backend
Then the system creates an open question, classifies it at creation time, and returns the created question

### AC-QUE-007: Prevent duplicate open buyer question
Given a buyer already has an open question for an order
When another question creation request is submitted for the same buyer and order
Then the system rejects the request with a conflict error

### AC-QUE-008: Reject invalid product reference
Given an order exists
When a question creation request references an order item from another order
Then the system rejects the request with a validation error

### AC-QUE-009: Buyer creates a question from order detail
Given an authenticated seeded Buyer account owns an order
When the buyer submits a valid question from the order detail view
Then the system creates a question in `question_active` lifecycle state and displays it in the order context

### AC-QUE-010: Seller answer awaits buyer approval
Given a buyer question is in `question_active` lifecycle state
When the seller submits a valid answer
Then the system stores the seller reply and moves the question to `seller_answered`

### AC-QUE-011: Buyer approves seller answer
Given a seller answer exists for the buyer's question
When the buyer approves the answer
Then the system records buyer approval metadata and displays the question lifecycle as closed

### AC-QUE-012: Hide question lifecycle for orders without questions
Given an order has no buyer questions
When any authorized role views that order in context
Then the UI shows the normal order status and does not show a question lifecycle label

## 9. Edge Cases
- Question without product reference.
- Question with product reference to one item in a multi-item order.
- Order with no questions.
- Attempted reply to a resolved question.
- Attempted reply by a seller who is not assigned to the question's order.
- Empty reply body.
- Empty question body.
- Duplicate open question for the same buyer and order.
- Duplicate active or unclosed buyer-facing question for the same buyer and order.
- Question linked to an order item from another order must be invalid.
- Threaded message history is requested; it remains outside MVP.
- A seller answer exists but the buyer has not approved it yet.
- An order has no questions and must not display question lifecycle status.
- A question lifecycle transition is attempted by a user who does not own the buyer or seller role required by the transition.
- A second seller reply is attempted after the one reply already exists.

## 10. Success criteria
- Sellers can see buyer questions in order context.
- Sellers can reply to open questions.
- Seller replies move questions to `seller_answered` while buyer approval remains required for closure.
- Resolved questions cannot receive additional MVP replies.
- Backend question creation can be tested without a buyer-facing UI.
- Backend buyer-owned order reads are available for seeded Buyer accounts.
- Buyer-facing UI lets buyers create questions from their own orders.
- Buyer approval closes questions only after a seller answer exists.
- Buyer approval records approval metadata and makes `closed` the terminal visible lifecycle.
- Question lifecycle labels are visible and consistent across Seller, Buyer, and Operations views when questions exist.

## 11. Dependencies
- Shared data model.
- API contract.
- Frontend contract.
- Priority and notifications specification for unresolved Operations queue behavior.

## 12. Open Questions
- None for this MVP normalization pass.
