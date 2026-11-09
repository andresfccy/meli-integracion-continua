# Backend Deep Dive

## Runtime and entry

**[DOCUMENTADO]** The backend uses Go 1.22 and port 8080. TASK-003B requires a GraphQL dependency and read-only schema. **[OBSERVADO]** `go.mod` declares `github.com/graph-gophers/graphql-go` v1.6.0; `main.go` starts `http.ListenAndServe` with the handler built by bootstrap.

Go fits this reconstruction because static types, small interfaces, built-in HTTP/JSON/testing packages, and a single compiled process support a low-dependency monolith. This is a technical defense; only the Go constraint and standard-library preference are explicitly documented.

## Go concepts in the code

| Concept | Example | Why it is used |
| --- | --- | --- |
| Package | `domain`, `application`, `httpadapter` | Defines responsibility and import direction |
| Struct | `Order`, `Question`, `Service` | Groups typed state |
| Method | `Order.Total`, `Service.Reply` | Associates behavior with data/coordinator |
| Interface | `Store`, `Delivery` | Lets application depend on capability, not implementation |
| Pointer | `*Service`, `*Reply`, optional account | Shares service state and represents optional/mutable values |
| Slice | orders, questions, roles | Ordered collections |
| Map | review notification groups | Groups items by seller |
| Error value | `ErrForbidden`, `ErrConflict` | Stable mapping from application outcome to HTTP |
| Function value | `now func() time.Time` | Makes time replaceable in tests |
| Mutex | memory `Store` | Protects concurrent reads/writes |

## Domain model

### Orders and money

An `Order` contains `OrderItem` snapshots. `ProductName` and `UnitPrice` are copied into each item so later catalog changes do not rewrite purchase history. `Total` multiplies quantity by purchase-time unit price.

**[OBSERVADO]** `float64` is used. **[RECOMENDACIÓN]** Production commerce systems should use integer cents or decimal arithmetic.

### Questions

`Question` holds:

- coarse `Status` (`open`/`resolved`),
- detailed `LifecycleStatus`,
- optional `Reply`,
- priority evidence,
- approval/resolution timestamps.

The two state dimensions support compatibility while exposing the buyer-approval lifecycle. They must remain separate from `Order.Status`.

### Priority

`EvaluatePriority` is pure relative to its inputs. It does not read storage or current time itself. The application provides timestamps and order value. This makes classification repeatable.

Limitations:

- substring matching may produce false positives,
- English dictionaries do not cover all user language,
- scores are hard-coded,
- overlapping phrases are recorded but only one keyword tier contributes,
- policy configuration requires deployment.

## Application ports

`Store` exposes collections and save operations. `Delivery` exposes one delivery command. The application does not know `memory.Store` or `notifications.Console`.

**[DEFENSA]** Narrow ports reduce the impact of infrastructure change. **[RECOMENDACIÓN]** A durable store would benefit from task-specific methods and transactional semantics rather than returning full collections.

### Notifiable factory and delivery registry

`NotifiableFactory` has two explicit creation methods: `FromSeller` and `FromOpsTeamMember`. Each validates the source identity, assigns a deterministic generic target ID, records the source type/ID, and copies the source channel-ID slice. Keeping the supported conversions explicit is appropriate for this MVP: the compiler and tests reveal which entity families may participate instead of accepting arbitrary objects through reflection.

During priority review, application code resolves the Seller’s `Notifiable`, follows its channel IDs, skips missing or disabled channels, builds an event addressed to the generic target, and invokes the `Delivery` port. The outbound `Dispatcher` then selects a registered driver using `NotificationChannel.Type`. The initial registry is `email → Console`; Console writes the simulated-email payload to an injected writer. This separates three decisions that should evolve independently:

```text
who can be notified
→ which logical channels that target selected
→ which infrastructure driver currently implements each channel type
```

Ops members are built and seeded through the same factory so future Operations notifications do not require a new targeting abstraction. The current approved High/Critical flow still sends seller batches only.

## Authentication

`Login`:

1. trims and compares email case-insensitively;
2. compares raw seed passwords with `hmac.Equal`;
3. returns the account and a token.

`token` constructs three JWT segments with a JSON HS256 header. `Authenticate` verifies header algorithm/type, signature, issuer, audience, expiration, and seeded subject.

**[VERIFICADO]** Application tests decode the issued header as JSON and authenticate the signed token.

**[RECOMENDACIÓN]**

- use a maintained JWT library,
- move the secret out of source,
- hash passwords,
- keep algorithm/header validation covered by tests,
- prefer secure cookie/session designs appropriate to the deployment threat model.

## Authorization and data access

Buyer reads require a buyer role and matching `BuyerID`. Seller and Operations REST routes require an exact `Bearer ` authorization scheme, authenticate the JWT, validate the role, and verify the matching seller identity where applicable. GraphQL performs the equivalent role and identity checks in typed resolvers.

Filters are applied in memory. Dates are compared as `YYYY-MM-DD` strings, which works for the normalized format but does not validate arbitrary input rigorously.

## Question lifecycle

### Create

- finds the order,
- checks authenticated buyer ownership when an account is supplied,
- rejects blank body,
- validates buyer and optional item,
- rejects any same buyer/order question with coarse status `open`,
- classifies at creation,
- stores `status=open`, `lifecycleStatus=question_active`.

The backend-only route passes no account; it intentionally simulates external ingestion. Buyer-facing creation passes the authenticated account.

### Reply

- rejects blank body,
- hides a question belonging to another seller as not found,
- requires `open` and no prior reply,
- attaches one reply,
- changes lifecycle to `seller_answered`,
- keeps coarse status open.

### Approve

- requires matching Buyer identity,
- checks order ownership,
- requires a reply and `seller_answered`,
- sets coarse status to `resolved`,
- sets lifecycle to `closed`,
- records approval and resolution time.

## Queue and review

Queue excludes non-open questions and applies filters. Sorting is deterministic:

1. priority score descending,
2. question age ascending by creation instant (older first),
3. order total descending,
4. question ID ascending.

Review counts open questions and reclassifies those at least 48 hours from `LastCategorizedAt`. High/Critical entries are grouped by seller and enabled channel. Scheduled execution uses a deterministic grouping key and atomically acquires one UTC date window; manual execution adds an auditable override suffix and can resend.

The scheduler calls the same use case through `ReviewScheduled`; duplicate windows return `already_executed` without creating another event.

## HTTP adapter

The REST router uses an explicit switch. For this scale it is transparent. Seller, Buyer, and Operations routes authenticate and authorize server-side. `/graphql` delegates to a dedicated schema-backed adapter. As REST routes grow, a specialized router could reduce manual path parsing.

### Real read-only GraphQL adapter

`internal/adapters/inbound/graphql/schema.graphql` is embedded into the Go binary and parsed at startup. A schema error therefore fails immediately instead of becoming a hidden runtime mismatch. `handler.go`:

1. requires and validates the bearer token;
2. places the authenticated account in request context;
3. lets the GraphQL engine parse and validate the document;
4. invokes typed resolvers for the requested fields;
5. verifies Seller, Buyer, or Operations ownership;
6. calls application read use cases;
7. maps domain/application values into the role-shaped response.

There is no `Mutation` root. Consequently, mutation documents fail schema validation and all state changes remain REST commands. The engine supports normal GraphQL variables, aliases, fragments, field selection, validation, and introspection. Depth 20 and query length 10,000 limits bound basic query complexity.

**[INFERIDO]** `graph-gophers/graphql-go` v1.6.0 was selected for the reconstruction because it provides a schema-first Go implementation with typed resolvers and a small integration surface. The recovered task requires a dependency but does not preserve the original package/version, so that historical detail cannot be claimed.

`DisallowUnknownFields` catches accidental request fields. `appError` maps stable errors to 400/403/404/409. Authentication failures at the bearer boundary return 401.

Operations review, REST queue, and composed queue reads require the Operations role. Seller REST/composed reads require the matching Seller identity.

## Verification

**[VERIFICADO 2026-07-30]**

- `go test -race ./...`: passed.
- `go vet ./...`: passed.
- Tests cover domain scoring/totals, application lifecycle/JWT/review, HTTP authorization/CORS/flows, scheduler timing/cancellation, concurrent windows, and notification drivers.
- Playwright verifies the principal role flows; exhaustive contract permutations remain future work.
