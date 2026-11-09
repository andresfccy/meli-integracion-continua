# Architecture and Implementation Question Bank

## Architecture

### Why not microservices?

**Short:** No independent scaling or ownership need justifies distributed complexity.  
**Expanded:** ADR-001 intentionally keeps one deployment while separating code through packages and ports. Microservices would add network failures, contract versioning, tracing, and operations.  
**Evidence:** ADR-001.  
**Reprompt:** Split only when measured scaling, reliability, ownership, or release needs diverge.

### What makes this hexagonal?

Application defines ports; memory and console implement outbound adapters; HTTP is inbound; domain does not import infrastructure. Bootstrap selects implementations.

### Is the application fully hexagonal?

Yes in dependency direction and adapter composition. The single service remains coarse, while GraphQL is isolated in a schema-backed inbound adapter whose resolvers call application read use cases.

## APIs

### Why REST for writes?

Commands express intent and invariants. A reply or approval is not a generic patch. ADR-002 keeps mutation ownership explicit.

### Why GraphQL for reads?

It composes role-shaped dashboards through an embedded schema, normal GraphQL documents, variables, aliases, fragments, field selection, validation, introspection, and typed Go resolvers. It intentionally has no mutation root.

### What is wrong with the current GraphQL implementation?

It uses an embedded schema and a real parser/executor, so field validation, selection, variables, aliases, fragments, and introspection work. Typed resolvers enforce Operations/Seller/Buyer authorization and call application read use cases. No `Mutation` root exists.

## Backend

### Why is priority in domain?

It is deterministic business policy independent of HTTP, schedule, storage, or provider.

### Why is the review in application?

Review coordinates time eligibility, persistence, queue composition, event creation, and delivery ports.

### Why use interfaces?

Application expresses capabilities without importing memory or console. This supports isolated tests and replacement.

### What is the largest backend debt?

The largest remaining debts are browser-level tests, source seed security, exact money representation, and in-memory infrastructure.

## Frontend

### Why organize by feature?

Role capability changes stay localized. Shared session/transport remains centralized.

### Does TypeScript make API data safe?

No. Types disappear at runtime; the JSON still needs validation. `gql<any>` further disables compile-time help.

### Why sessionStorage?

ADR-003 limits authentication to the browser session for the MVP. It is not a production security guarantee.

## Data

### Why store item name and price on the order?

Purchase history must not change when catalog values change. The order item is a snapshot.

### Is float64 correct for money?

Not for production financial correctness. Integer minor units or decimals are safer.

### Why in memory?

Deterministic, infrastructure-free MVP. It must be replaced for durability and multiple instances.

## Lifecycle

### Why both status and lifecycleStatus?

`open/resolved` preserves coarse compatibility; lifecycle identifies active, seller-answered, and closed phases. Neither belongs to order status.

### Why does seller reply not close the question?

Recovered decisions 41/54 require buyer approval. The seller changes the party expected to act; the buyer closes.

## Notifications and scheduling

### Why a driver?

The domain/application decide what to notify; the driver knows how. `NotifiableFactory` normalizes Sellers and Ops members, each target owns channel IDs, and the dispatcher maps logical channel types to drivers. In the MVP `email` maps to Console to simulate sending. Provider changes should not change classification or targeting.

### Is daily scheduling implemented?

Yes. The inbound scheduler calculates the next 00:00 UTC execution, supports cancellation, invokes the shared review use case, and relies on atomic in-memory window acquisition to suppress duplicate scheduled execution.

### Can manual review resend?

Yes. Manual review adds a nanosecond override suffix and delegates again. Application, HTTP, and Playwright tests verify the manual path; scheduled duplicate-window behavior remains separately protected.

## Testing

### Do tests prove the project?

No single suite proves the project. Race-enabled Go tests cover domain, application, HTTP/GraphQL, scheduler, memory, and drivers; Vitest covers session/routing helpers; the production build passes; and Playwright covers the principal role flows. More isolated component and negative security permutations would further deepen evidence.

### What tests should be added first?

Playwright now covers seeded login/error, Buyer→Seller→Buyer lifecycle, Operations seller filtering, and manual review. The next testing investment is deeper React component isolation, more security/error permutations, and durable-adapter contracts.

## Reconstruction

### Is this the original application?

No verifiable claim can be made. It is a local regeneration from recovered documentation.

### How can you defend decisions you may not have historically made?

Separate explicit decision records from inferred technical rationale. Defend what the architecture represents, disclose uncertainty, and avoid invented history.

### What demonstrates senior judgment here?

Traceability, trade-off analysis, consistent boundaries, explicit risk acceptance, recognition of evidence gaps, and a prioritized production evolution path.
