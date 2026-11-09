# SPEC-001: Seller Orders

## 1. Status
- State: Normalized
- Last updated: 2026-07-13
- Related tasks:
- TASK-001
- TASK-002
- TASK-006

## 2. Objective
Provide sellers with a single dashboard view of their recent orders so they can identify who purchased, what was purchased, the purchase-time prices, and the current order status.

## 3. Actors
- Seller

## 4. User Stories

### US-ORD-001
As a seller
I want to view my recent orders, so that I can understand what customers purchased.

### US-ORD-002
As a seller
I want to search and filter orders by approved criteria, so that I can quickly find orders that need attention.

## 5. Scope
### In scope
- Recent seller orders belonging to the seller.
- Buyer information.
- Order date.
- Order status.
- Products contained in each order.
- Quantity and purchase-time unit price.
- Order total.
- Search and filtering by status, buyer text, and date range.
- Empty and no-results states.
- Current order status display remains separate from any question lifecycle labels.

### Authentication preconditions
- A seller must authenticate with a seeded seller account before accessing the seller orders dashboard.
- The seller orders dashboard uses the authenticated seeded seller identity to determine the current seller.
- Authentication exists only to enable the MVP seller experience and is defined in the shared frontend, API, and data model artifacts.

### Out of scope
- Creating orders.
- Payments.
- Shipping-provider integrations.
- Inventory management.
- Authentication behavior beyond the shared lightweight MVP flow.
- Question lifecycle states as order status values.

## 6. Functional requirements
### FR-ORD-001
The system shall display recent orders belonging to the current seller.

### FR-ORD-002
The system shall display buyer name, buyer contact, order date, order status, order items, quantities, purchase-time unit prices, and order total for each order.

### FR-ORD-003
The system shall allow sellers to filter orders by status.

### FR-ORD-004
The system shall allow sellers to search orders by buyer text.

### FR-ORD-005
The system shall allow sellers to filter orders by date range.

### FR-ORD-006
The system shall display an empty state when the seller has no matching orders.

### FR-ORD-007
The system shall display purchase-time unit prices for order items instead of current catalog prices.

## 7. Business rules
### BR-ORD-001
An order must contain one or more products.

### BR-ORD-002
An order item must preserve quantity and unit price at purchase time.

### BR-ORD-003
Recent orders must be returned in deterministic order, sorted by order date descending and then order identifier ascending when dates are equal.

### BR-ORD-004
Order status must describe the order lifecycle only. Question lifecycle labels must not be absorbed into order status.

### BR-ORD-005
Orders without buyer questions must not display question lifecycle labels. When a question exists, its lifecycle label may be displayed alongside the order context for authorized users.

## 8. Acceptance Scenarios

### AC-ORD-001: View recent orders
Given seeded orders exist for the seller
When the seller opens the orders view
Then the system displays the seller's recent orders

### AC-ORD-002: View order item details
Given a seeded order contains multiple products
When the seller views the order
Then the system displays each item quantity, purchase-time unit price, and the order total

### AC-ORD-003: Filter orders by status
Given seeded orders exist in different statuses
When the seller filters orders by one status
Then the system displays only orders with that status

### AC-ORD-004: Search orders by buyer text
Given seeded orders exist for multiple buyers
When the seller searches by buyer text
Then the system displays only matching buyer orders

### AC-ORD-005: Filter orders by date range
Given seeded orders exist across multiple dates
When the seller applies a date range
Then the system displays only orders within that range

### AC-ORD-006: No matching orders
Given no orders match the selected filters
When the seller views the filtered result
Then the system displays a no-results empty state

## 9. Edge Cases
- Seller has no orders
- Order contains multiple products
- Two orders have the same date
- Search produces no results
- Seller has a long list of orders; pagination is deferred outside MVP unless later approved

## 10. Success criteria
- The seller can identify who purchased, what was purchased and the current status.
- Results are deterministic for the same data and filters.
- Approved filters work independently and in combination.
- Empty states are visible and understandable.

## 11. Dependencies
- Shared data model
- API contract
- Frontend contract

## 12. Buyer-Facing Extension

- Buyer-facing order listing is implemented as a separate role-specific view using the same shared order data and buyer ownership constraints.
- Buyer-facing order reads must not change seller dashboard ownership rules.
- Buyer-facing question lifecycle labels are question-level UI metadata, not order status values.

## 13. Open Questions
- None for this MVP normalization pass.
