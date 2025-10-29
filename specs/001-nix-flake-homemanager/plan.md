# Implementation Plan: Nix Flake & Home-Manager Support

**Branch**: `001-nix-flake-homemanager` | **Date**: 2025-10-29 | **Spec**: [specs/001-nix-flake-homemanager/spec.md](specs/001-nix-flake-homemanager/spec.md)
**Input**: Feature specification from `/specs/001-nix-flake-homemanager/spec.md`

**Note**: This template is filled in by the `/speckit.plan` command. See `.specify/templates/commands/plan.md` for the execution workflow.

## Summary

Add Nix flake support for building TaskHerald and home-manager module for declarative service configuration. The flake will provide taskherald as default package for x86_64-linux, and include a home-manager module that creates a user-level systemd service with configurable environment variables.

## Technical Context

<!--
  ACTION REQUIRED: Replace the content in this section with the technical details
  for the project. The structure here is presented in advisory capacity to guide
  the iteration process.
-->

## Technical Context

**Language/Version**: Go (existing TaskHerald codebase)  
**Primary Dependencies**: Nix flakes, home-manager  
**Storage**: N/A (packaging feature, no data storage)  
**Testing**: Existing Go test suite  
**Target Platform**: x86_64-linux (specified in requirements)  
**Project Type**: Packaging/deployment (adds Nix files to existing Go project)  
**Performance Goals**: Flake build completes within 30 seconds  
**Constraints**: x86_64-linux architecture only, user-level systemd service  
**Scale/Scope**: Single package output, single systemd service per user

## Constitution Check

_GATE: Must pass before Phase 0 research. Re-check after Phase 1 design._

- **Service Reliability**: ✅ Design includes systemd service integration via home-manager with automatic restart
- **CLI Contract**: ✅ No changes to Taskwarrior CLI interactions
- **Security**: ✅ User-level service with least privilege, no sensitive data exposure
- **Observability**: ✅ Inherits existing TaskHerald logging capabilities
- **Simplicity**: ✅ Minimal Nix configuration focused on packaging and service management
- **Deployment Constraints**: ✅ Uses only Nix flakes and home-manager, no additional external dependencies

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

```text
# Root level additions
flake.nix              # Nix flake definition
flake.lock             # Flake lock file (generated)

# Existing structure unchanged
src/                   # Go source code (existing)
go.mod                 # Go module (existing)
go.sum                 # Go dependencies (existing)
```

**Structure Decision**: This is a packaging feature that adds Nix files to the existing Go project structure. No changes to the core Go codebase are required.

## Complexity Tracking

> **Fill ONLY if Constitution Check has violations that must be justified**

| Violation                  | Why Needed         | Simpler Alternative Rejected Because |
| -------------------------- | ------------------ | ------------------------------------ |
| [e.g., 4th project]        | [current need]     | [why 3 projects insufficient]        |
| [e.g., Repository pattern] | [specific problem] | [why direct DB access insufficient]  |
