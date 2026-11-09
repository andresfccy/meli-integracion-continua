# Decisions

This directory preserves durable project decisions. Relevant architecture and product decisions are summarized as ADRs; the preserved decision log below keeps the original numbered decision history and supersession context.

## ADR index

- [ADR-001: Lightweight Hexagonal Monolith](ADR-001-lightweight-hexagonal-monolith.md)
- [ADR-002: REST Commands and GraphQL Reads](ADR-002-rest-commands-graphql-reads.md)
- [ADR-003: Seeded Authentication for MVP](ADR-003-seeded-authentication.md)
- [ADR-004: Buyer Approval Question Lifecycle](ADR-004-question-lifecycle.md)
- [ADR-005: Priority Review and Notification Delivery Boundary](ADR-005-priority-notification-delivery.md)

## Decision status model

- `accepted`: active decision.
- `superseded`: replaced by a newer accepted decision.
- `deprecated`: preserved for history and no longer recommended.
- `proposed`: not yet accepted.

## How to use this file

- Check ADRs first for current architectural direction.
- Use the preserved log to retain traceability of all prior decisions.
- Do not delete superseded decisions; add supersession notes instead.
- If a future change reverses a durable decision, add a new ADR or update this index with explicit status and references.

## Preserved numbered decision log

1. Keep the backend in Go 1.22 using the standard library unless an existing dependency is already present and necessary.
2. Keep persistence in memory for the MVP.
3. Preserve backend port 8080 and frontend port 3000 constraints from the project instructions.
4. Migrate frontend package management from npm lockfiles to pnpm lockfiles.
5. Preserve `start`, `build`, and `test` script names for platform compatibility.
6. Do not implement domain behavior before specs are reviewed.
7. Build backend packages by responsibility: domain, services, handlers, storage, and notifications.
8. Build frontend by feature, not by technical layer.
9. Do not modify read-only path patterns `backend/src/test/*` and `.vscode/tasks.json` when present; treat them as protected guardrails even if absent from a local workspace snapshot.
10. Treat unresolved question priority as a documented product policy, not hidden code behavior.
11. Simulate current email notifications with console output until another channel is specified.
12. Use seeded in-memory data to demonstrate orders, products, buyers, questions, and statuses.
13. Track implementation work in atomic, independently committable tasks.
14. Verify Go changes with `gofmt`, `go test ./...`, and `go vet ./...`.
15. Verify frontend changes with pnpm scripts once the frontend is implemented.
16. MVP order filters are status, buyer text search, and date range.
17. Superseded by decisions 39, 40, 48, and 53 for the current implementation: buyer-facing question creation from owned order detail is in scope; broader buyer question history UI remains out of scope.
18. MVP priority factors are waiting time, order value, and urgency words from a simple dictionary.
19. Priority thresholds are Critical >= 80, High Priority >= 60, Medium >= 30, and Low < 30.
20. Superseded by decisions 41 and 54 for the current implementation: seller replies now leave questions unresolved in `seller_answered` until buyer approval.
21. Operations is a dedicated section within the existing dashboard, not a separate application.
22. Seed data must include application roles sufficient to represent Seller and Operations usage.
23. MVP authentication uses only seeded seller and Operations accounts with email/password validation against in-memory seed data; registration, recovery, reset, verification, social login, persistent backend sessions, JWT refresh tokens, OAuth, and external identity providers are out of scope.
24. Backend-only question creation is included for MVP testing and external-ingestion simulation; buyer-facing question creation UI remains out of scope.
25. MVP questions use a one-to-one question/reply model; threaded message lists are deferred outside MVP.
26. Superseded by decisions 41 and 54 for the current implementation: accepted seller replies move questions to `seller_answered` until buyer approval closes the lifecycle.
27. A buyer may not have more than one open question for the same order.
28. MVP priority classification uses dictionary version `mvp-2026-07-13` with approved Critical and High Priority urgency words.
29. Questions are classified at creation time and open questions are eligible for daily 00:00 recategorization only after at least 2 days since the latest categorization.
30. Notification delivery simulation uses notifiable targets and enabled notification channels instead of seller email-only fields.
31. MVP notifications are triggered only by High Priority and Critical classifications, then batched by notifiable target, channel, and review date to avoid notification spam in default scheduled and non-override flows.
32. GraphQL is approved only as a frontend-facing read composition layer for authenticated role-shaped UI data; GraphQL must not own business rules, expose public contracts, replace REST commands, or perform mutations in the MVP.
33. REST remains the MVP command API for login, backend question creation, seller replies, and priority review triggers.
34. gRPC, GraphQL mutations/subscriptions/federation, buyer-facing APIs, webhooks, pagination, IaC, event streaming, read-optimized stores, configurable thresholds, and reminder cadence are post-MVP nice-to-haves unless later approved.
35. `docs/archive/PROJECT_FILES_INSTRUCTIONS.md` is initial platform guidance; current decisions, specs, and package metadata override it when contradictory.
36. Local/preview CORS preflight support for frontend port 3000 to backend port 8080 belongs at the inbound HTTP router edge.
37. TASK-001B login validation supports HackerRank preview CORS by allowing frontend origins ending in `-3000.us-vmprovider.projects.hrcdn.net` to call the paired backend preview host while still rejecting unrelated origins and avoiding wildcard CORS.
38. TASK-003B adds short-lived JWT access tokens only for authenticated GraphQL reads; REST remains the command API, and refresh tokens, OAuth, external identity providers, and durable backend sessions remain out of scope.
39. Buyer-facing work adds seeded Buyer accounts and authenticated buyer order visibility without introducing registration, password recovery, OAuth, refresh tokens, or external identity providers.
40. Buyer-created questions move buyer-facing question creation into implemented scope for owned order detail; backend-only question creation remains available for MVP testing and external-ingestion simulation.
41. Seller replies no longer close buyer-facing questions immediately; they move the question lifecycle to a seller-answered state until the buyer approves the answer.
42. Order status remains separate from question lifecycle status. Orders without questions show only order status, while existing questions expose question lifecycle labels to authorized roles.
43. Question lifecycle states must be question-specific, such as `question_active`, `seller_answered`, `buyer_approved`, and `closed`; order-like states such as `created` and `accepted` do not belong to the question lifecycle.
44. Notification delivery separates notification event creation from delivery through outbound notification channel driver ports.
45. Notification delivery uses a registry or dispatcher to select drivers by channel type so Slack, SMS, disabled channels, removed channels, or future protocols do not require changes to priority classification behavior.
46. The MVP console notification driver may write enabled channel payloads to stdout, but domain and application logic must remain independent from stdout and channel-specific protocols.
47. Shipping, warranty, and returns workflows remain outside the next buyer-facing slice.
48. Buyer-facing implementation keeps current MVP question `open`/`resolved` status semantics while exposing `lifecycleStatus` for buyer-created active questions and buyer-approval lifecycle states.
49. Buyer approval records buyer approval metadata and makes `closed` the terminal visible question lifecycle; `buyer_approved` may be an audit or transition state rather than the final user-facing label.
50. Buyer-created question duplicate detection applies to active or unclosed questions for the same buyer and order, including `question_active` and `seller_answered`.
51. TASK-014 notification driver delivery is scoped to seller notification events from priority review; buyer notification delivery remains deferred until buyer notification rules are separately approved.
52. Disabled notification channels are skipped deterministically, while unsupported enabled channel types return deterministic delivery errors at the driver boundary.
53. TASK-012 implements authenticated buyer-created questions for buyer-owned order detail as a REST command, stores `lifecycleStatus=question_active` while preserving `status=open`, and keeps GraphQL read-only.
54. Decisions 41, 48, 49, 50, and 53 supersede earlier seller-reply closure decisions 20 and 26 for the current buyer-facing implementation: seller replies move questions to `seller_answered` and buyer approval performs closure.
55. TASK-018 adds an inbound daily 00:00 UTC priority review scheduler that invokes the same Operations application use case as the manual HTTP trigger, with scheduled execution-window acquisition to prevent duplicate scheduled delivery for the same UTC review window.
56. Operations manual priority review is the explicit human override path: it is protected by an Operations access token, bypasses scheduled execution-window acquisition, and may resend same-day seller notification batches using manual override notification grouping keys.
57. MVP seeded notification channels are email-only. Slack and SMS remain supported by the domain/driver boundary as future channel types, but SMS is not seeded for the current MVP.