# Frontend Contract

## 1. Purpose
Define the MVP dashboard behavior for seller order visibility, seller replies to buyer questions, and the Operations priority section.

## 2. Backend URL resolution
Browser API calls must derive the backend URL from the frontend location by replacing frontend port `3000` with backend port `8080`. The frontend must not hard-code a different host for local preview.

## 3. Views and sections

### Login view
- Allows users to authenticate with email and password from the seeded seller, Operations, and Buyer accounts.
- Seeded Buyer accounts use the same login flow and can access buyer order/question creation routes.
- Displays deterministic demo password hints for seeded profiles in the in-memory challenge environment.
- Displays validation feedback for missing credentials or invalid credentials.
- Does not expose registration, password recovery, password reset, email verification, social login, OAuth, or external identity provider options.
- Stores authenticated user state only for the current browser session.
- Stores the JWT access token returned by login only in the current browser session and sends it as a Bearer token for authenticated GraphQL reads.
- Redirects authenticated users to the route allowed for their seeded role.
- If future seed data grants multiple roles to one account, landing-route priority must remain deterministic: Seller first, then Operations, then Buyer, unless a later decision changes the order.

### Seller dashboard
- Displays recent orders for the active seeded seller role.
- Shows buyer, order date, status, items, quantities, purchase-time unit prices, and order total.
- Provides filters for status, buyer text, and date range.
- Allows selecting an order to view associated questions.
- May consume the GraphQL `sellerDashboard` and `sellerOrderDetail` read queries for role-shaped display data.

### Order questions section
- Displays questions linked to the selected order.
- Shows question body, status, timestamp, buyer, optional product reference, and existing reply when present.
- Allows replying to open questions.
- Submitting an accepted seller reply moves the question to `seller_answered`; the buyer must approve the answer before the question lifecycle closes.
- Prevents reply interaction for already answered, approved, or closed questions.
- Question lifecycle labels are displayed for existing questions only. Orders without questions do not show question lifecycle labels.

### Buyer dashboard
- Appears for seeded users with the Buyer role.
- Displays orders belonging to the authenticated buyer.
- Shows seller, order date, order status, items, quantities, purchase-time unit prices, and order total.
- Allows selecting an order to view existing questions and seller replies.
- Allows creating a question from order detail through the REST command endpoint.
- Allows approving a seller answer when the question is in `seller_answered` lifecycle state.
- Shows question lifecycle labels such as `question_active`, `seller_answered`, `buyer_approved`, and `closed` only when questions exist.
- Does not show order-like question states such as `created` or `accepted`.
- Displays a duplicate active-question conflict when the buyer already has an active or unclosed question for the same order.
- Displays an approve-answer conflict when the question has no seller reply, is not in `seller_answered`, is already approved, or is already closed.

### Operations priority section
- Appears inside the same dashboard application for seeded users with the Operations role.
- Displays unresolved questions sorted by priority.
- Shows priority score, priority classification, waiting time signal, order value signal, urgency word signal, seller, buyer, and order summary.
- Provides filters for seller, buyer text, order status, date range, and priority classification.
- The seller filter is a single search input that matches seller name, seller email, or seller identifier without requiring separate fields.
- Allows Operations users to manually trigger priority review and notification event dispatch through the REST command endpoint. This manual action is the approved same-day notification resend path and requires the current Operations access token.
- Shows loading, success, and recoverable error states for the manual priority review trigger, including notification batch counts and execution status when the trigger succeeds.
- Does not require a separate application.
- May consume the GraphQL `operationsPriorityQueue` read query for role-shaped display data.

### Authenticated role state
- The MVP must use seeded account data sufficient to represent Seller and Operations usage.
- Seeded account data supports Buyer usage and buyer order/question creation routes.
- UI components are shown only when the authenticated seeded account has the related role.
- A seller account can access only the seller dashboard.
- An Operations account can access the Operations priority section inside the existing dashboard application.
- A Buyer account can access only buyer order and question UI.
- Unauthenticated users must be redirected to the login page before accessing protected routes.
- Logout clears the authenticated browser-session state and returns the user to the login page.

## 4. User interactions
- Filter orders by status.
- Search orders by buyer text.
- Filter orders by date range.
- Clear filters.
- Select an order to inspect details and questions.
- Submit a seller reply for an open question, moving it to `seller_answered` while it remains unresolved until buyer approval closes it.
- Navigate to or reveal the Operations priority section when the active seeded role allows it.
- Log in with seeded email and password.
- Log out and clear the current browser-session authenticated state.
- Apply Operations priority queue filters.
- Trigger Operations priority review manually from the Operations profile.
- Buyer interactions include viewing buyer-owned orders, creating an order question, reviewing a seller answer, and approving the answer to close the question lifecycle.

## 5. Route protection
- The login route is public.
- Seller dashboard routes are protected and require an authenticated seeded seller account.
- Operations routes or sections are protected and require an authenticated seeded Operations account.
- Buyer routes are protected and require an authenticated seeded Buyer account.
- Unauthorized users must be redirected to the login page.
- Role-inappropriate routes or UI sections must not be shown to the authenticated user.

## 6. Browser route map

- `/login`: public login view.
- `/`: authenticated landing route that redirects the current seeded user to the first route allowed by their role.
- `/seller/orders`: protected seller dashboard order list.
- `/seller/orders/:orderId`: protected seller order detail and order questions view.
- `/operations/questions`: protected Operations priority queue route or route-backed section inside the dashboard application.
- `/buyer/orders`: protected buyer order list.
- `/buyer/orders/:orderId`: protected buyer order detail, question creation, seller-answer review, and buyer approval view.
- If the UI renders Operations as a tab or section instead of a standalone page, the browser route must still resolve to the protected Operations priority section and enforce the same role checks.

## 7. UI states

### Loading
- Show a loading state while login, orders, order details, replies, roles, or Operations questions are being fetched.
- Show skeleton placeholders for order lists, order detail, question lists, reply submission, and Operations priority queue while data is loading.

### Error
- Show a recoverable error state when an API request fails.
- Show validation feedback for empty reply body or invalid filters.
- Show login validation feedback for missing or invalid credentials.
- Buyer question creation shows duplicate active-question conflicts as recoverable form errors. Buyer approval shows approve-answer lifecycle conflicts as recoverable form errors.

### Success
- Redirect to the authorized dashboard area after successful login.
- Show the seller reply and updated question lifecycle label after a successful reply submission; `status` remains `open` until buyer approval closes the question.
- Show updated filtered results after filter changes.
- Show priority review summary feedback after an Operations manual notification trigger succeeds, including same-day resend results when notification batches are delivered again.
- Return to the login page after logout.

### Empty
- Show an empty orders state when the seller has no orders.
- Show a no-results state when filters match no orders.
- Show an empty questions state when an order has no questions.
- Show an empty Operations queue state when no unresolved questions exist.
- Buyer views show an empty buyer-orders state when the buyer has no orders.
- Buyer views do not show question lifecycle labels when an order has no questions.

## 8. GraphQL read usage

- GraphQL is used only for authenticated, role-shaped frontend read models.
- GraphQL requests must include the current session JWT in the `Authorization` header.
- GraphQL must not be used for login, backend question creation, seller replies, recategorization triggers, or any state-changing behavior.
- GraphQL responses should reduce frontend overfetching and multiple REST read calls for dashboard rendering.
- The frontend must continue to respect route protection and role-specific UI visibility regardless of GraphQL response shape.
- GraphQL usage here is contract guidance; implementation progress is tracked in `docs/plans/implementation-plan.md`.
- Buyer order reads may use GraphQL read composition, but buyer question creation and buyer answer approval remain REST commands.
- Buyer read GraphQL queries must remain role-shaped and authenticated; GraphQL mutations remain out of scope for buyer question creation and answer approval.

## 9. Authentication requirements and acceptance scenarios

### FR-AUTH-001
The system shall allow a seeded seller or Operations user to authenticate with email and password.

### FR-AUTH-002
The system shall reject authentication when email or password does not match a seeded account.

### FR-AUTH-003
The system shall redirect unauthenticated users to the login page when they attempt to access protected routes.

### FR-AUTH-004
The system shall allow an authenticated seller account to access only seller dashboard UI.

### FR-AUTH-005
The system shall allow an authenticated Operations account to access only Operations role UI inside the dashboard application.

### FR-AUTH-006
The system shall clear the current browser-session authenticated state on logout and return the user to the login page.

### FR-AUTH-007
The system shall allow an authenticated seeded Buyer account to access only buyer order and question UI.

### FR-BUYER-001
The system shall display orders belonging to the authenticated buyer.

### FR-BUYER-002
The system shall allow the buyer to create a question from an owned order detail view.

### FR-BUYER-003
The system shall allow the buyer to approve a seller answer and close the question lifecycle.

### FR-BUYER-004
The system shall display question lifecycle labels only for existing questions.

### AC-AUTH-001: Login with seeded credentials
Given a seeded seller or Operations account exists
When the user submits the matching email and password
Then the system authenticates the user for the current browser session

### AC-AUTH-002: Reject invalid credentials
Given no seeded account matches the submitted credentials
When the user attempts to log in
Then the system displays a login error and does not authenticate the user

### AC-AUTH-003: Redirect unauthenticated access
Given the user is not authenticated
When the user attempts to access a protected route
Then the system redirects the user to the login page

### AC-AUTH-004: Seller role access
Given the user is authenticated as a seeded seller
When the dashboard is displayed
Then only seller dashboard UI is available

### AC-AUTH-005: Operations role access
Given the user is authenticated as a seeded Operations account
When the dashboard application is displayed
Then only Operations role UI is available

### AC-AUTH-006: Logout
Given the user is authenticated
When the user logs out
Then the system clears browser-session authenticated state and returns the user to the login page

### AC-AUTH-007: GraphQL read authentication
Given the user is authenticated through seeded login
When the dashboard requests GraphQL read data
Then the frontend sends the current session JWT as a Bearer token

### AC-BUYER-001: Buyer role access
Given the user is authenticated as a seeded Buyer account
When the dashboard is displayed
Then only buyer order and question UI is available

### AC-BUYER-002: Buyer creates question
Given the buyer owns an order
When the buyer submits a valid question from order detail
Then the question appears with `question_active` lifecycle status

### AC-BUYER-003: Buyer approves answer
Given a seller answer exists for the buyer's question
When the buyer approves the answer
Then the question lifecycle closes and the updated status is shown

### AC-BUYER-004: No question lifecycle without questions
Given a buyer order has no questions
When the order is displayed
Then only normal order status is shown and no question lifecycle label is displayed