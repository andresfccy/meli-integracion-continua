# Reconstruction Context and Evidence Rules

## What this repository represents

**[NO VERIFICABLE]** The original application remains in an inaccessible cloud environment. Its exact source, runtime configuration, data, deployment history, and test results cannot be inspected.

**[DOCUMENTADO]** The recovered repository documentation defines a seller dashboard MVP with Seller, Buyer, and Operations roles; Go and React runtimes; REST commands; authenticated read-only GraphQL composition; in-memory persistence; priority review; and notification drivers.

**[OBSERVADO]** This repository contains a regenerated Go backend and React/TypeScript frontend that implement a meaningful subset of those definitions.

The safe description is:

> The recovered documentation establishes the intended behavior and decisions. The regenerated local application materializes those definitions in the way observed here. The reconstruction is useful for study and technical defense, but it cannot prove that the unavailable original application had identical code.

## Evidence hierarchy

Different questions use different sources:

1. “What does the regenerated code do?” — inspect code and run tests.
2. “What behavior was approved?” — inspect active specifications and contracts.
3. “Why was a durable direction accepted?” — inspect ADRs and the numbered decision log.
4. “Was a task declared complete?” — inspect the implementation plan.
5. “Did the original cloud version work exactly this way?” — **[NO VERIFICABLE]**.

## How to speak about decisions

Use precise wording:

| Evidence | Safe phrase |
| --- | --- |
| ADR explicitly compares alternatives | “The documented decision was…” |
| Technology appears in metadata but lacks an ADR | “The recovered project constrains the implementation to…” |
| Code structure suggests a reason | “A technically reasonable interpretation is…” |
| Preparing an interview answer | “The architecture can be defended because…” |
| Original behavior is unknown | “I cannot verify whether the inaccessible original used the same implementation.” |

Do not say “I chose X because Y” unless the recovered record actually supports that historical claim. Say “The decision record chooses X; Y is the documented or technically defensible rationale.”

## Snapshot limitations

**[OBSERVADO]**

- The directory contains no `.git`; history, authorship, commits, branches, and prior diffs are unavailable.
- The backend remains compact, while domain, application, HTTP, scheduler, memory, and notification behavior now have focused tests.
- Seller, Buyer, Operations, session, and GraphQL-client files were expanded into readable modules; some small legacy files and the stylesheet remain compressed.
- The documented scheduler package, UTC calculation, cancellation, bootstrap wiring, and thread-safe window acquisition are present.
- The backend now contains an embedded read-only GraphQL schema and uses `github.com/graph-gophers/graphql-go` v1.6.0.
- **[DOCUMENTADO]** TASK-003B required a GraphQL dependency and read-only schema. **[INFERIDO]** The recovered material does not identify the exact original package/version, so this library is a reconstruction choice that satisfies the documented boundary.
- Backend coverage spans six packages; frontend auth/client coverage remains small and browser E2E is not present.
- Build output and installed dependencies are present, but they are not substitutes for source-level verification.

## Why this remains valuable

**[DEFENSA]** Architecture reasoning does not require pretending that every historical detail is known. A senior defense can clearly separate:

- constraints,
- intended contracts,
- observed implementation,
- accepted trade-offs,
- reconstruction gaps,
- and a production evolution path.

That separation demonstrates stronger judgment than overstating certainty.

## Questions for study

1. Which statements in an architecture defense are historically documented?
2. Which are reasonable technical inferences?
3. What evidence would be required to claim parity with the original cloud application?
4. Why can a regenerated implementation still be useful for explaining architecture?
