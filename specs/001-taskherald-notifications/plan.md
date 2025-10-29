# Implementation Plan: TaskHerald Notifications

**Branch**: `001-taskherald-notifications` | **Date**: 2025-10-29 | **Spec**: spec.md
**Input**: Feature specification from `/specs/[###-feature-name]/spec.md`

**Note**: This template is filled in by the `/speckit.plan` command. See `.specify/templates/commands/plan.md` for the execution workflow.

## Summary

TaskHerald is a single Go binary systemd service that regularly checks Taskwarrior tasks for notification_date matches and sends ntfy.sh notifications. On startup, it sends summaries of missed notifications and logs task verification. Configuration via environment variables, tracks sent notifications in JSON file.

## Technical Context

<!--
  ACTION REQUIRED: Replace the content in this section with the technical details
  for the project. The structure here is presented in advisory capacity to guide
  the iteration process.
-->

## Technical Context

**Language/Version**: Go 1.21  
**Primary Dependencies**: ntfy client library for Go (NEEDS CLARIFICATION: which library, e.g., github.com/binwiederhier/ntfy)  
**Storage**: Taskwarrior UDAs for tracking sent notifications  
**Testing**: Go testing framework  
**Target Platform**: Linux (systemd compatible)  
**Project Type**: Single binary service  
**Performance Goals**: Check tasks every 60s, send notifications within 5s of notification_date match  
**Constraints**: Single user, least privilege, external network only to ntfy.sh  
**Scale/Scope**: Handle <1000 active Taskwarrior tasks per user

## Constitution Check

_GATE: Must pass before Phase 0 research. Re-check after Phase 1 design._

- **Service Reliability**: Design must include systemd service integration and automatic failure recovery.
- **CLI Contract**: All Taskwarrior interactions must use the `task` command with input/output validation.
- **Security**: Implementation must operate with least privilege and avoid exposing sensitive data.
- **Observability**: All actions must be logged with timestamps; errors must be actionable.
- **Simplicity**: Feature set must remain minimal, avoiding unnecessary complexity.
- **Deployment Constraints**: Only Go, systemd, Taskwarrior CLI, and ntfy.sh for notifications permitted; external network access limited to ntfy.sh.

## Project Structure

### Documentation (this feature)

```text
specs/[###-feature]/
├── plan.md              # This file (/speckit.plan command output)
├── research.md          # Phase 0 output (/speckit.plan command)
├── data-model.md        # Phase 1 output (/speckit.plan command)
├── quickstart.md        # Phase 1 output (/speckit.plan command)
├── contracts/           # Phase 1 output (/speckit.plan command)
└── tasks.md             # Phase 2 output (/speckit.tasks command - NOT created by /speckit.plan)
```

### Source Code (repository root)

<!--
  ACTION REQUIRED: Replace the placeholder tree below with the concrete layout
  for this feature. Delete unused options and expand the chosen structure with
  real paths (e.g., apps/admin, packages/something). The delivered plan must
  not include Option labels.
-->

```text
# [REMOVE IF UNUSED] Option 1: Single project (DEFAULT)
src/
├── models/
├── services/
├── cli/
└── lib/

tests/
├── contract/
├── integration/
└── unit/

# [REMOVE IF UNUSED] Option 2: Web application (when "frontend" + "backend" detected)
backend/
├── src/
│   ├── models/
│   ├── services/
│   └── api/
└── tests/

frontend/
├── src/
│   ├── components/
│   ├── pages/
│   └── services/
└── tests/

# [REMOVE IF UNUSED] Option 3: Mobile + API (when "iOS/Android" detected)
api/
└── [same as backend above]

ios/ or android/
└── [platform-specific structure: feature modules, UI flows, platform tests]
```

**Structure Decision**: [Document the selected structure and reference the real
directories captured above]

## Complexity Tracking

> **Fill ONLY if Constitution Check has violations that must be justified**

| Violation                  | Why Needed         | Simpler Alternative Rejected Because |
| -------------------------- | ------------------ | ------------------------------------ |
| [e.g., 4th project]        | [current need]     | [why 3 projects insufficient]        |
| [e.g., Repository pattern] | [specific problem] | [why direct DB access insufficient]  |
