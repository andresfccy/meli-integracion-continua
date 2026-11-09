# ADR-005: Priority Review and Notification Delivery Boundary

## Status

Accepted

## Context

Operations needs deterministic priority review and seller notifications for High Priority and Critical questions. The MVP must avoid provider-specific dependencies while preserving future notification extensibility.

## Decision

Use a shared Operations priority review use case for manual and scheduled review. Scheduled review acquires one execution per UTC review window. Manual Operations review is an authenticated override path and may resend same-day seller notification batches through manual override grouping keys.

Notification event creation and delivery are separate. Delivery goes through outbound channel-driver ports selected by channel type. The MVP registers a console driver for email channels.

## Consequences

- Priority policy is independent of notification protocols.
- Scheduled review remains duplicate-safe.
- Manual resend behavior is explicit and auditable.
- Slack, SMS, and other providers can be added behind the driver boundary later.

## Alternatives considered

- Immediate notification during question creation: rejected; review-time batching is the approved behavior.
- Application services writing directly to stdout: rejected; stdout belongs to the console driver adapter.
- Real external providers in MVP: deferred.

## References

- [../../specs/SPEC-003-priority-notifications.md](../../specs/SPEC-003-priority-notifications.md)
- [../architecture/notifications.md](../architecture/notifications.md)
- [../plans/implementation-plan.md](../plans/implementation-plan.md)
