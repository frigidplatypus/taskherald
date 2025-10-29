---

description: "Task list template for feature implementation"
---

# Tasks: TaskHerald Notifications

**Input**: Design documents from `/specs/001-taskherald-notifications/`
**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md, contracts/

**Tests**: Tests are OPTIONAL - not requested in specification.

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Path Conventions

- **Single project**: `src/`, `tests/` at repository root
- Paths shown below assume single project structure

## Dependencies

User stories can be implemented independently after foundational tasks complete. US1 (P1) should be implemented first as MVP. US2-US4 can run in parallel after US1.

## Implementation Strategy

Start with MVP: Implement US1 for core notification functionality. Then add US2, US3, US4 incrementally. Each user story delivers independent value.

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure

- [X] T001 Create Go module and basic project structure in src/
- [X] T002 Initialize Go project with ntfy.sh client dependency

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story can be implemented

- [X] T003 Setup configuration loading from environment variables in src/config.go
- [X] T004 Implement Taskwarrior CLI integration for task export in src/taskwarrior.go
- [X] T005 Setup structured logging with timestamps in src/logger.go
- [X] T006 Implement ntfy client with priority/tag mapping in src/ntfy.go
- [X] T007 Create main service loop structure in src/main.go
- [X] T008 Implement 60-second task check interval in src/main.go

## Phase 3: User Story 1 - Receive Task Notifications (Priority: P1) 🎯 MVP

**Goal**: Send notifications for tasks with notification_date matching current time

**Independent Test**: Set a task's notification_date to current time, run service, verify notification sent to ntfy.sh

- [X] T009 [US1] Implement task filtering for current notifications in src/notifications.go
- [X] T010 [US1] Implement notification message formatting with project/description/due in src/notifications.go
- [X] T011 [US1] Implement UDA update after successful notification in src/taskwarrior.go

## Phase 4: User Story 2 - Startup Summary Notifications (Priority: P2)

**Goal**: Send summary of missed notifications on startup

**Independent Test**: Stop service, create past-due tasks, restart service, verify summary notification sent

- [X] T012 [US2] Implement startup summary logic for past 7 days in src/startup.go

## Phase 5: User Story 3 - Startup Task Verification (Priority: P3)

**Goal**: Log tasks with notification_date on startup for verification

**Independent Test**: Start service, check logs for up to 10 tasks or "no tasks found" message

- [X] T013 [US3] Implement startup task logging in src/startup.go

## Phase 6: User Story 4 - Custom Ntfy Configuration (Priority: P4)

**Goal**: Support custom ntfy server and default topic configuration

**Independent Test**: Configure custom server/topic, send notification, verify delivery to custom destination

- [X] T014 [US4] Implement custom server/topic configuration in src/config.go
- [X] T015 [US4] Implement per-task topic override via ntfy_topic UDA in src/notifications.go

## Phase 7: Polish & Cross-Cutting Concerns

**Purpose**: Final integration, error handling, and deployment

- [X] T016 Add comprehensive error handling and logging for failures in src/main.go
- [X] T017 Create systemd service unit file in systemd/taskherald.service
- [X] T018 Add graceful shutdown handling in src/main.go
- [X] T019 Update README with installation and usage instructions