# Presentation Guide

## Five-minute version

1. **Provenance (30s):** reconstructed from recovered documentation; not claimed identical to original.
2. **Architecture (60s):** React → REST commands/schema-backed GraphQL reads → Go application/domain → memory/drivers.
3. **Key decision (60s):** modular monolith balances simplicity and boundaries.
4. **Representative flow (90s):** Buyer asks → priority → Seller answers → Buyer approves.
5. **Trade-offs (45s):** memory, seeds, console are MVP constraints.
6. **Honest audit (15s):** browser-testing, security, persistence, and delivery gaps are identified.

## Fifteen-minute version

Suggested slides:

1. Evidence and reconstruction rules.
2. Requirement-to-code traceability chain.
3. System context diagram.
4. Hexagonal dependency direction.
5. REST commands versus GraphQL reads.
6. Question lifecycle.
7. Priority policy.
8. Notification drivers.
9. Frontend feature organization.
10. Testing strategy versus observed coverage.
11. Documentation/code gaps.
12. Production evolution.

## Thirty-minute version

Add:

- package walkthrough;
- login sequence;
- question sequence;
- Operations review sequence;
- technology decision matrix;
- alternatives and trade-offs;
- code excerpts for ports, `EvaluatePriority`, lifecycle, and client;
- live demo if services can be verified;
- audience Q&A using the question bank.

## Demonstration script

Only demonstrate behavior that has been run successfully in the current environment.

1. Start backend/frontend.
2. Login as Buyer and open owned order.
3. Create question.
4. Login as Seller and answer it.
5. Login as Buyer and approve it.
6. Login as Operations and inspect queue/review.

Explain the scheduler through its tests: next UTC midnight, cancellation, shared use case, one-winner window acquisition, and manual override. Avoid waiting for a real midnight during a presentation.

## Safe language

- “The recovered ADR states…”
- “The regenerated code implements…”
- “This can be defended technically because…”
- “The current snapshot has a gap…”
- “For production I would evolve it by…”
- “I cannot verify parity with the original cloud instance.”

## Language to avoid

- “The original definitely used this exact file.”
- “All tasks are complete” based only on checkboxes.
- “The original used this exact GraphQL library/version.” The recovered task requires a Go GraphQL dependency and schema, but the exact original package is not verifiable.
- “Tests prove the whole system.”
- “JWT makes it secure.”
- “Microservices are always more scalable.”

## Recommended visuals

- dependency-direction diagram;
- question lifecycle;
- end-to-end sequence;
- traceability matrix excerpt;
- documented-versus-observed audit table.

## Closing

> The strongest part of the design is the separation of business policy from delivery mechanisms. The reconstruction demonstrates the central flows, while the audit identifies where implementation evidence does not yet meet the recovered contract. That combination gives both a defensible architecture and a concrete evolution plan.
