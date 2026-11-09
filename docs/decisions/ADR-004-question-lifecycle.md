# ADR-004: Buyer Approval Question Lifecycle

## Status

Accepted

## Context

Initial seller-reply behavior closed questions immediately. The buyer-facing slice requires seller answers to remain unresolved until the owning buyer approves the answer.

## Decision

Keep coarse question `status` values as `open` and `resolved`, and add question-specific `lifecycleStatus` values for buyer-facing lifecycle display. Seller replies move questions to `seller_answered` while `status` remains `open`. Buyer approval records approval metadata and closes the lifecycle with terminal visible `lifecycleStatus=closed` and `status=resolved`.

Order status remains separate from question lifecycle status.

## Consequences

- Operations continues to see `question_active` and `seller_answered` questions as unresolved.
- Orders without questions do not show question lifecycle labels.
- Earlier immediate-closure decisions are superseded.

## Alternatives considered

- Immediate seller-reply closure: superseded because buyer approval is now required.
- Order-like question lifecycle states such as `created` or `accepted`: rejected because question lifecycle must remain question-specific.

## References

- [../../specs/SPEC-002-buyer-questions.md](../../specs/SPEC-002-buyer-questions.md)
- [../../specs/data-model.md](../../specs/data-model.md)
- [../../specs/api-contract.md](../../specs/api-contract.md)
