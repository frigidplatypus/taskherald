# Feature Specification: Nix Flake & Home-Manager Support

**Feature Branch**: `001-nix-flake-homemanager`  
**Created**: 2025-10-29  
**Status**: Draft  
**Input**: User description: "Taskherald should provide a flake.nix that builds and outputs taskherald as the default package for x86_64-linux. A homeManager module should also be outputted. The module should create a user level systemd service that uses the following attribute set."

## User Scenarios & Testing _(mandatory)_

### User Story 1 - Install TaskHerald via Nix Flake (Priority: P1)

As a Nix user, I want to install TaskHerald using a flake so I can easily build and install the application on my system.

**Why this priority**: This is the core requirement - without the flake, users can't install TaskHerald via Nix.

**Independent Test**: Can be fully tested by running `nix build` on the flake and verifying the taskherald binary is produced.

**Acceptance Scenarios**:

1. **Given** a flake.nix exists in the repository, **When** I run `nix build .#taskherald`, **Then** a taskherald binary is built for x86_64-linux
2. **Given** the flake has taskherald as default package, **When** I run `nix build`, **Then** taskherald binary is built
3. **Given** the built binary, **When** I execute it, **Then** it runs and shows help/version information

---

### User Story 2 - Configure via Home-Manager Module (Priority: P2)

As a Nix user with home-manager, I want to configure TaskHerald through a home-manager module so I can manage the service declaratively alongside my other system configuration.

**Why this priority**: This enables the primary use case - running TaskHerald as a user service with proper configuration.

**Independent Test**: Can be fully tested by importing the home-manager module and verifying the systemd service is created with correct configuration.

**Acceptance Scenarios**:

1. **Given** home-manager configuration with `services.taskherald.enable = true`, **When** I apply the configuration, **Then** a user-level systemd service is created
2. **Given** home-manager settings with `ntfy_topic = "my-tasks"`, **When** the service runs, **Then** notifications are sent to the "my-tasks" topic
3. **Given** optional settings like `ntfy_server` and `taskherald_interval`, **When** configured, **Then** the service uses those values instead of defaults

---

### User Story 3 - Automatic Service Management (Priority: P3)

As a user, I want the TaskHerald service to start automatically when I log in so I don't have to manually start it each time.

**Why this priority**: This ensures the service runs reliably without user intervention.

**Independent Test**: Can be fully tested by checking that the systemd service is enabled and starts on user login.

**Acceptance Scenarios**:

1. **Given** home-manager configuration is applied, **When** I log in to my user session, **Then** the taskherald service starts automatically
2. **Given** the service is running, **When** I log out and log back in, **Then** the service restarts automatically
3. **Given** configuration changes, **When** I update home-manager, **Then** the service is restarted with new configuration

---

### Edge Cases

- What happens when flake is built on unsupported architecture?
- How does system handle missing Taskwarrior configuration?
- What happens when ntfy_topic is not provided in home-manager config?
- How does service behave when network connectivity to ntfy server is lost?

## Requirements _(mandatory)_

### Functional Requirements

- **FR-001**: flake.nix MUST provide a build for taskherald binary targeting x86_64-linux
- **FR-002**: flake.nix MUST output taskherald as the default package
- **FR-003**: flake.nix MUST include home-manager module as an output
- **FR-004**: Home-manager module MUST create user-level systemd service when enabled
- **FR-005**: Home-manager module MUST require ntfy_topic setting with no default value
- **FR-006**: Home-manager module MUST accept optional ntfy_server setting
- **FR-007**: Home-manager module MUST accept optional taskherald_interval setting as integer
- **FR-008**: Systemd service MUST start automatically on user login
- **FR-009**: Systemd service MUST use environment variables for configuration
- **FR-010**: Service MUST restart when home-manager configuration is updated

### Key Entities _(include if feature involves data)_

- **Flake Output**: Nix flake providing packages and home-manager modules
- **Home-Manager Module**: Declarative configuration module for user services
- **Systemd Service**: User-level service that runs taskherald with proper environment

## Success Criteria _(mandatory)_

### Measurable Outcomes

- **SC-001**: Users can successfully build taskherald using `nix build` within 30 seconds
- **SC-002**: Home-manager module enables taskherald service in under 10 seconds
- **SC-003**: 100% of users with valid configuration can receive task notifications
- **SC-004**: Service automatically starts for 95% of user login sessions
- **SC-005**: Configuration changes apply correctly in 100% of home-manager updates

## Constitution Alignment

_GATE: Must verify compliance with TaskHerald Constitution v1.1.0 before finalizing spec._

- **Service Reliability**: Feature ensures service runs as user-level systemd service with automatic restart on login
- **CLI Contract**: No changes to Taskwarrior CLI interactions
- **Security**: User-level service runs with user privileges, no elevated permissions needed
- **Observability**: Inherits existing logging capabilities from taskherald binary
- **Simplicity**: Minimal Nix configuration focused only on packaging and service management
- **Deployment Constraints**: Uses only Nix flakes and home-manager, no additional external dependencies

## Clarifications

### Session 2025-10-29

- Q: What is the appropriate build timeout for the flake build success criteria? → A: 30 seconds
