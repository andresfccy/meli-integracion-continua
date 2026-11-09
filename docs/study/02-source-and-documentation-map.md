# Source and Documentation Map

## Authority by purpose

| Source | Owns | Does not own |
| --- | --- | --- |
| `specs/SPEC-*.md` | Functional requirements, rules, acceptance scenarios | Implementation status |
| `specs/data-model.md` | Entities, fields, relationships, invariants | Concrete Go structs as an implementation detail |
| `specs/api-contract.md` | HTTP/GraphQL contract and error behavior | Handler organization |
| `specs/frontend-contract.md` | Routes, role visibility, UI states | Component internals |
| `specs/test-plan.md` | Required verification scenarios | Proof that tests exist or pass |
| `docs/decisions` | Accepted durable choices and alternatives | Current implementation truth |
| `docs/architecture` | Boundaries, dependency direction, evolution rules | Product requirements |
| `docs/plans` | Declared task status | Executable evidence |
| Code and package metadata | Regenerated implementation reality | Original-cloud history |
| Tests and command output | Executed coverage and current result | Untested behavior |
| `docs/archive` | Historical context | Active authority |

## Reading a requirement into code

Example: buyer approval.

1. [SPEC-002](../../specs/SPEC-002-buyer-questions.md) defines the lifecycle.
2. [Data model](../../specs/data-model.md) defines approval metadata.
3. [API contract](../../specs/api-contract.md) defines the command.
4. [ADR-002](../decisions/ADR-002-rest-commands-graphql-reads.md) keeps state changes on REST.
5. The implementation plan declares TASK-013 complete.
6. `Service.Approve` checks role, buyer ownership, reply existence, and lifecycle.
7. The HTTP adapter exposes `/api/buyers/{buyerId}/questions/{questionId}/approve-answer`.
8. `Buyer.tsx` invokes the command and updates its local typed question state.
9. **[VERIFICADO]** application, HTTP, and Playwright tests exercise the complete Buyer→Seller→Buyer flow.

## Documentation precedence

When documents disagree, use the repository’s declared hierarchy:

```text
current code and package metadata
→ active specifications
→ implementation plan
→ accepted decisions
→ architecture
→ README/development guidance
→ archive
```

This does not mean code is automatically correct. It means code tells us what the regenerated application currently does, while specifications tell us whether that behavior is approved.

## Important source clusters

### Orders

- `SPEC-001`
- data model
- API contract
- frontend contract
- TASK-001, TASK-002, TASK-006
- domain `Order.Total`
- application order reads and filters
- seller/buyer components

### Questions

- `SPEC-002`
- ADR-004
- decisions 40–54
- TASK-003, TASK-003A, TASK-012, TASK-013
- `CreateQuestion`, `Reply`, `Approve`

### Priority and notifications

- `SPEC-003`
- ADR-005
- decisions 28–31, 44–46, 51–57
- TASK-004, TASK-005, TASK-014, TASK-016, TASK-018
- `EvaluatePriority`, `Queue`, `Review`, console driver

### Transport and authentication

- API and frontend contracts
- ADR-002 and ADR-003
- decisions 23, 32–40
- router, token functions, session client, route selection

## Study rule

Never copy a contract into a derived document and then treat the copy as authoritative. Link to the owner, explain how it is implemented, and record any mismatch in the audit.
