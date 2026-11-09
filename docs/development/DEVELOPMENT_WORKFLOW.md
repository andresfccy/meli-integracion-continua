# Development Workflow

1. Inspect repository state, code, package metadata, docs, and relevant specs.
2. Plan affected files, acceptance criteria, verification commands, dependencies, and risks.
3. Review the plan against approved specs and decisions.
4. Implement the smallest deterministic change.
5. Test with task-specific commands.
6. Verify build, formatting, linting, and documentation links where applicable.
7. Document behavior, command, contract, or structure changes in the owning document.

## Stop conditions

- Required behavior is not specified.
- A change would alter an accepted requirement, contract, or decision without approval.
- A new dependency is needed but not justified.
- A requested change requires modifying read-only files.
- Tests fail for reasons not understood after focused investigation.

## Plan history

If a task creates a standalone `*-PLAN.md`, preserve its execution history in [../archive/BINACLE.md](../archive/BINACLE.md) and move executed plan content to archive or remove the standalone file after preservation.
