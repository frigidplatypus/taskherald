<!--
Sync Impact Report
Version change: 1.0.0 → 1.1.0
Modified principles: None
Added sections: None
Removed sections: None
Templates requiring updates:
✅ plan-template.md (updated constitution gates)
✅ spec-template.md (added constitution alignment section)
✅ tasks-template.md (added principle-driven foundational tasks)
⚠ commands/ (directory missing, manual check needed)
Follow-up TODOs: RATIFICATION_DATE
-->

# TaskHerald Constitution

## Core Principles

### Service Reliability

The service MUST run continuously as a systemd service and recover automatically from failures.
Rationale: Ensures uninterrupted task management and user trust.

### CLI Contract

All interactions with Taskwarrior MUST use the `task` command, with strict input/output validation.
Rationale: Guarantees compatibility and prevents data corruption.

### Security

The service MUST operate with least privilege and never expose sensitive data.
Rationale: Protects user privacy and system integrity.

### Observability

All actions MUST be logged with timestamps; errors MUST be reported and actionable.
Rationale: Enables troubleshooting and accountability.

### Simplicity

The codebase MUST remain minimal, with only essential features implemented.
Rationale: Reduces maintenance burden and risk of bugs.

## Deployment Constraints

Only Go, systemd, Taskwarrior CLI, and ntfy.sh for notifications are permitted. External network access is limited to ntfy.sh.

## Development Workflow

All changes require local testing. Version must be bumped according to semantic versioning rules:

- MAJOR: Breaking changes to principles or governance
- MINOR: New principle/section added or expanded
- PATCH: Clarifications or non-semantic refinements

## Governance

Amendments are made by direct edit to this file. Version bump per semantic rules. Compliance reviewed before each deployment.

**Version**: 1.1.0 | **Ratified**: TODO(RATIFICATION_DATE): original adoption date unknown | **Last Amended**: 2025-10-29
