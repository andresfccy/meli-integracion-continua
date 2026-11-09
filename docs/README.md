# Documentation Index

This is the central map for project documentation. Use it to find the authoritative source before changing code, specs, tests, or documentation.

## Start here by need

| Need | Read first | Then read |
| --- | --- | --- |
| Run or understand the project quickly | [../README.md](../README.md) | [development/TESTING_GUIDELINES.md](development/TESTING_GUIDELINES.md) |
| Understand AI-agent rules | [../AGENTS.md](../AGENTS.md) | [development/agent-workflow.md](development/agent-workflow.md) |
| Change product behavior | [../specs/README.md](../specs/README.md) | [decisions/README.md](decisions/README.md) |
| Change backend code | [architecture/backend.md](architecture/backend.md) | [../backend/README.md](../backend/README.md), [../specs/api-contract.md](../specs/api-contract.md) |
| Change frontend code | [architecture/frontend.md](architecture/frontend.md) | [../frontend/README.md](../frontend/README.md), [../specs/frontend-contract.md](../specs/frontend-contract.md) |
| Change priority or notifications | [../specs/SPEC-003-priority-notifications.md](../specs/SPEC-003-priority-notifications.md) | [architecture/notifications.md](architecture/notifications.md) |
| Check implementation status | [plans/implementation-plan.md](plans/implementation-plan.md) | [plans/implementation-backlog.md](plans/implementation-backlog.md) |
| Review decisions | [decisions/README.md](decisions/README.md) | The related ADR files in [decisions/](decisions/) |
| Review historical context | [archive/README.md](archive/README.md) | [archive/BINACLE.md](archive/BINACLE.md) |
| Study the regenerated application and prepare an architecture defense | [study/README.md](study/README.md) | [study/llm/consolidated-study-guide.md](study/llm/consolidated-study-guide.md) |

## Documentation precedence

When documents disagree, use this order:

1. Current code and package metadata for what is implemented.
2. Active specs for approved behavior and contracts.
3. [plans/implementation-plan.md](plans/implementation-plan.md) for implementation status.
4. [decisions/README.md](decisions/README.md) and ADRs for accepted decisions and supersession history.
5. Architecture docs for design boundaries and rationale.
6. Root and module READMEs for onboarding and local maps.
7. Development docs for workflow and verification.
8. Archive docs only for historical context.

## Normative documentation

- [../AGENTS.md](../AGENTS.md): required operating guide for AI agents.
- [../specs/README.md](../specs/README.md): requirements and contract index.
- [decisions/README.md](decisions/README.md): decision index and preserved decision log.
- [architecture/README.md](architecture/README.md): architecture overview.
- [development/CONTRIBUTING.md](development/CONTRIBUTING.md): contribution rules.
- [development/TESTING_GUIDELINES.md](development/TESTING_GUIDELINES.md): verification commands.

## Historical documentation

Historical documents are preserved under [archive/](archive/) and must not be used as active guidance unless promoted into specs, decisions, architecture, or development docs.

## Derived study documentation

[study/README.md](study/README.md) is a derived learning and architecture-defense package for the locally regenerated application. It cross-references normative sources and observed code, but it does not redefine requirements, contracts, decisions, or implementation status.

## Maintenance rules

- Put each concept in one authoritative document.
- Link to source-of-truth content instead of copying it.
- Keep current implementation separate from approved future direction.
- Move executed plans and deprecated guidance to [archive/](archive/).
- Update this index whenever documentation paths or authority change.
