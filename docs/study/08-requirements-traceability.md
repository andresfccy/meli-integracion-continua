# Requirements Traceability

## How to use this matrix

“Declared” means the implementation plan marks the work complete. “Observed” means code exists. “Verified” is limited to executed tests.

| Capability | Normative source | Decision/task | Backend | Frontend | Test evidence | Assessment |
| --- | --- | --- | --- | --- | --- | --- |
| Seed authentication | API/frontend contracts | ADR-003; TASK-001A/B | login, HS256 validation, protected routes | Login/session | application + HTTP + client tests | Verified focused paths |
| Seller order list/detail | SPEC-001 | TASK-002/006 | owned REST/composed reads | typed Seller list/detail | HTTP auth + domain totals | Aligned |
| Status/buyer/date filters | SPEC-001 | TASK-002/006 | validates and applies all | exposes and forwards all | application + HTTP invalid filters | Aligned |
| Purchase-time prices | SPEC-001/data model | TASK-001/002 | `OrderItem.UnitPrice`, `Total` | item totals rendered | domain test passes | Verified narrowly |
| Buyer-owned reads | SPEC-002/API | decisions 39/40; TASK-011 | `BuyerOrders`, `BuyerOrder` | `Buyer` | none visible | Observed |
| Backend question creation | SPEC-002/API | TASK-003A | route + `CreateQuestion(nil,...)` | none by design | HTTP contract test | Verified |
| Buyer question creation | SPEC-002 | decision 53; TASK-012 | authenticated route + service | active-question-aware form | application + HTTP lifecycle | Verified happy path |
| Duplicate prevention | SPEC-002 BR/FR | decision 50 | coarse open blocks | form checks open only | application conflicts | Aligned |
| Seller reply | SPEC-002 | decisions 41/54; TASK-003/013 | seller-authorized transition | typed Seller form | application + HTTP lifecycle | Verified |
| Buyer approval | SPEC-002 | decisions 49/54; TASK-013 | buyer-authorized closure | approval button | application + HTTP lifecycle | Verified |
| Priority calculation | SPEC-003 | TASK-004 | `EvaluatePriority` | score/class shown | thresholds, caps, keywords | Verified core rules |
| Operations queue | SPEC-003/API | TASK-004B/007 | role-protected REST/composed read | full filter UI | HTTP role tests | Aligned |
| Seller search | SPEC-003/test plan | TASK-016 | ID/name/email | displays and sends all | HTTP/application coverage | Aligned |
| Manual review | SPEC-003/API | decision 56 | protected `ReviewManual` | detailed feedback | application + HTTP | Verified |
| Scheduled review | SPEC-003 | decision 55; TASK-018 | inbound scheduler + bootstrap | n/a | scheduler tests | Observed and verified |
| Scheduled window dedupe | SPEC-003/data model | decision 55 | atomic memory acquisition | n/a | race + application tests | Verified |
| Manual notification resend | SPEC-003 | decision 56 | manual suffix | result count | application test | Verified |
| Notification driver | SPEC-003 | ADR-005; TASK-014 | email console driver | n/a | disabled/unsupported/email | Verified |
| Read-only GraphQL | API/frontend contracts | ADR-002; TASK-003B | embedded schema, engine, typed role resolvers | typed role queries and field selections | variables/aliases/fragments, validation, role and mutation tests | Aligned; exact original dependency not verifiable |
| Mutation rejection | API contract | ADR-002 | rejects mutations/non-reads | no mutations | HTTP test | Verified |
| JWT validation | API contract | ADR-003 | header + HMAC + claims | bearer injection | application + client tests | Verified MVP behavior |
| Preview CORS | frontend contract | decisions 36/37 | router edge allowlist | backend URL derivation | HTTP + client tests | Verified focused path |
| Loading/error/empty UI | frontend contract | multiple tasks | error shapes | partial states | absent | Partial |

## Reverse map: code to justification

| Code area | Justified by |
| --- | --- |
| `domain.Order.Total` | purchase-time price and total requirements |
| `domain.EvaluatePriority` | priority factors, thresholds, determinism |
| application ports | ADR-001 dependency direction |
| application ownership checks | authenticated role/ownership contracts |
| HTTP CORS middleware | decisions 36–37 |
| memory adapter | MVP persistence decision |
| console driver | notification delivery boundary |
| GraphQL branch | ADR-002 read composition |
| `sessionStorage` helper | ADR-003 browser-session-only state |
| feature folders | frontend architecture decision |

## Highest-value missing evidence

1. Deeper isolated React component tests for loading/error/empty/conflict states.
2. More malformed/expired/claim-specific JWT permutations.
3. Durable repository/outbox contract tests when those adapters are introduced.
4. Additional CORS denial and preview-host permutations.
5. Exact boundary tests for every priority threshold and tie-break combination.

## Traceability conclusion

The local reconstruction now covers and verifies the central interactive flows, scheduler, notification targeting/dispatch, and schema-backed GraphQL boundary. This still does not prove parity with the inaccessible original or exhaustively prove every plan checkbox; remaining evidence gaps are primarily depth and production-adapter concerns.
