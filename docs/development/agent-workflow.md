# Agent Workflow

This guide expands the root [../../AGENTS.md](../../AGENTS.md) operating rules for multi-step agent work.

## Planning

- Start from the documentation index and relevant source-of-truth files.
- Inspect code before changing documentation that describes implementation.
- Keep changes small and scoped.
- List affected files in the task summary.
- Protect unrelated changes discovered through `git status`.

## Documentation updates

Update the document that owns the concept:

- Requirement or contract: [../../specs/README.md](../../specs/README.md)
- Implementation status or task sequencing: [../plans/implementation-plan.md](../plans/implementation-plan.md)
- Architecture boundary: [../architecture/README.md](../architecture/README.md)
- Durable decision: [../decisions/README.md](../decisions/README.md)
- Workflow or verification: files in [./](./)
- Historical context: [../archive/README.md](../archive/README.md)

Do not copy source-of-truth text into multiple places. Link to it.

## Completion

Before reporting completion:

1. Run relevant verification or explain why it could not run.
2. Validate changed documentation links and paths.
3. Review the diff for unrelated changes.
4. Summarize modified files, verification, and any pending human decisions.
