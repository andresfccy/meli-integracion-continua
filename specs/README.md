# Specifications Index

Specs are the source of truth for approved functional requirements, contracts, and verifiable behavior.

## Active specs

- [SPEC-001: Seller Orders](SPEC-001-seller-orders.md)
- [SPEC-002: Buyer Questions](SPEC-002-buyer-questions.md)
- [SPEC-003: Priority and Notifications](SPEC-003-priority-notifications.md)

## Shared contracts

- [Data Model](data-model.md)
- [API Contract](api-contract.md)
- [Frontend Contract](frontend-contract.md)
- [Test Plan](test-plan.md)

## Implementation status

Implementation status is tracked outside specs in [../docs/plans/implementation-plan.md](../docs/plans/implementation-plan.md). Do not infer completion from a requirement being present in this directory; verify status against the plan, tests, and code.

## Scope rules

- Specs define what the system must do.
- Specs do not own agent workflow, historical reports, package maps, or implementation logs.
- Plans and architecture docs may reference specs, but they must not redefine contracts.
- When changing behavior, update the relevant spec before or with implementation changes.
