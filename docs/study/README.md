# Regenerated Application Study Package

## Purpose

This package explains how the locally regenerated application implements the recovered requirements and how its architecture can be defended technically. It is designed for human study, presentations, and ingestion into tools such as NotebookLM or Gemini.

This is **derived documentation**. The authoritative requirements remain in [../../specs](../../specs), accepted decisions in [../decisions](../decisions), current architecture in [../architecture](../architecture), and declared task status in [../plans/implementation-plan.md](../plans/implementation-plan.md).

## Provenance labels

- **[DOCUMENTADO]**: explicitly stated by a recovered specification, contract, ADR, plan, or README.
- **[OBSERVADO]**: directly visible in this local source snapshot.
- **[VERIFICADO]**: confirmed by a command, test, or build during this documentation pass.
- **[INFERIDO]**: a reasonable technical conclusion, not a demonstrated historical motive.
- **[DEFENSA]**: an architecture argument suitable for a presentation.
- **[RECOMENDACIÓN]**: proposed improvement, not approved behavior.
- **[BRECHA]**: mismatch among requirements, documentation, implementation, or tests.
- **[NO VERIFICABLE]**: would require the unavailable original cloud application or other missing evidence.

## Recommended reading paths

### Fast orientation

1. [Reconstruction context](01-reconstruction-context.md)
2. [Architecture explained](03-architecture-explained.md)
3. [End-to-end flows](06-end-to-end-flows.md)
4. [Senior architecture defense](11-senior-architecture-defense.md)

### Deep technical study

1. [Source and documentation map](02-source-and-documentation-map.md)
2. [Backend deep dive](04-backend-deep-dive.md)
3. [Frontend deep dive](05-frontend-deep-dive.md)
4. [Technology decisions](07-technology-decisions.md)
5. [Requirements traceability](08-requirements-traceability.md)
6. [Documentation/code audit](09-documentation-code-audit.md)
7. [Testing and quality](10-testing-and-quality.md)

### Presentation preparation

1. [Senior architecture defense](11-senior-architecture-defense.md)
2. [Presentation guide](12-presentation-guide.md)
3. [Question bank](13-question-bank.md)
4. [Glossary](14-glossary.md)

### LLM ingestion

- Use [the consolidated guide](llm/consolidated-study-guide.md) when the tool accepts few files.
- Use [the source manifest](llm/source-manifest.md) and all specialist documents when the tool supports a larger corpus.
- Start from [the semantic index](llm/corpus-index.md) for suggested prompts and topic routing.

## Core method

Every important behavior is studied through the same chain:

```text
Requirement
→ specification and contract
→ accepted decision
→ declared task
→ regenerated code
→ technology
→ test evidence
→ observed result
```

The package never treats a checked task as sufficient proof. Code establishes what is present in the regenerated snapshot; tests establish only what they actually exercise.

## Snapshot summary

- Analysis date: 2026-07-30.
- Backend: Go 1.22 module `seller-dashboard`.
- Frontend: Vite, React, TypeScript, pnpm 10.12.1.
- Source tree: available locally without Git metadata.
- Backend verification: `go test ./...` and `go vet ./...` passed.
- Frontend verification: installed Vitest passed 5 session/routing assertions, TypeScript and Vite built successfully, and Playwright passed 3 integrated browser scenarios.
- Backend verification after alignment: `go test -race ./...` and `go vet ./...` pass across domain, application, HTTP, scheduler, memory, and notification packages.
- GraphQL verification after alignment: the backend has a parsed read-only schema, typed resolvers, standard field selection/variables/aliases/fragments/validation, role authorization, and mutation rejection.
- Remaining material gaps: production-grade security, persistence, notification infrastructure, and deeper isolated component coverage.

## Maintenance rule

When code or normative documents change, update traceability and audit first, then propagate the result to the explanatory and LLM documents. Never edit a recovered specification merely to make this study package look consistent.
