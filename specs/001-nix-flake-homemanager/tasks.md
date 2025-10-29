# Tasks: Nix Flake & Home-Manager Support

**Input**: Design documents from `/specs/001-nix-flake-homemanager/`
**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md, contracts/

**Tests**: No test tasks included - feature specification does not request TDD approach.

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Path Conventions

- **Single project**: `src/`, `tests/` at repository root
- **Packaging feature**: Adds `flake.nix`, `flake.lock` at repository root
- Paths shown below follow repository root structure

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization for Nix flake support

- [ ] T001 Create flake.nix with basic Nixpkgs setup and Go build inputs
- [ ] T002 Configure flake-utils for multi-system support (x86_64-linux focus)

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core Nix infrastructure that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [ ] T003 Verify Go module structure and dependencies in go.mod
- [ ] T004 Test basic Go build works: `go build -o taskherald ./src`
- [ ] T005 Confirm TaskHerald binary runs and shows expected output

**Checkpoint**: Foundation ready - Nix flake implementation can now begin

---

## Phase 3: User Story 1 - Install TaskHerald via Nix Flake (Priority: P1) 🎯 MVP

**Goal**: Enable users to build and install TaskHerald using Nix flake

**Independent Test**: Run `nix build` and verify taskherald binary is produced and executable

### Implementation for User Story 1

- [ ] T006 [US1] Implement buildGoModule configuration in flake.nix for TaskHerald
- [ ] T007 [US1] Configure flake outputs to provide taskherald package for x86_64-linux
- [ ] T008 [US1] Set taskherald as default package in flake.nix
- [ ] T009 [US1] Add package metadata (description, license, maintainers) to flake.nix
- [ ] T010 [US1] Test flake build: `nix build` produces working binary
- [ ] T011 [US1] Verify built binary functionality matches go build output

**Checkpoint**: At this point, User Story 1 should be fully functional - users can build TaskHerald via Nix flake

---

## Phase 4: User Story 2 - Configure via Home-Manager Module (Priority: P2)

**Goal**: Provide declarative TaskHerald configuration through home-manager module

**Independent Test**: Import home-manager module and verify systemd service is created with correct environment variables

### Implementation for User Story 2

- [ ] T012 [US2] Create home-manager module structure in flake.nix outputs
- [ ] T013 [US2] Implement services.taskherald configuration options (enable, settings)
- [ ] T014 [US2] Add validation for required ntfy_topic setting
- [ ] T015 [US2] Configure optional ntfy_server and taskherald_interval settings
- [ ] T016 [US2] Implement systemd user service creation in home-manager module
- [ ] T017 [US2] Set up environment variables mapping from config to service
- [ ] T018 [US2] Configure service dependencies and restart behavior
- [ ] T019 [US2] Test home-manager module configuration and service creation

**Checkpoint**: At this point, User Stories 1 AND 2 should both work independently - users can build via flake AND configure via home-manager

---

## Phase 5: User Story 3 - Automatic Service Management (Priority: P3)

**Goal**: Ensure TaskHerald service starts automatically on user login

**Independent Test**: Verify systemd service is enabled and starts automatically on login

### Implementation for User Story 3

- [ ] T020 [US3] Configure systemd service WantedBy for automatic login start
- [ ] T021 [US3] Set service Type to simple for proper lifecycle management
- [ ] T022 [US3] Configure Restart=always for automatic failure recovery
- [ ] T023 [US3] Test service auto-start behavior on login/logout
- [ ] T024 [US3] Verify service survives system suspend/resume cycles

**Checkpoint**: All user stories should now be independently functional - complete Nix flake and home-manager integration

---

## Phase 6: Polish & Cross-Cutting Concerns

**Purpose**: Final improvements and validation

- [ ] T025 Update README.md with Nix flake installation instructions
- [ ] T026 Add flake.lock to version control
- [ ] T027 Test cross-platform compatibility (x86_64-linux only as specified)
- [ ] T028 Validate quickstart.md instructions work end-to-end
- [ ] T029 Run final integration test with real Taskwarrior tasks

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - can start immediately
- **Foundational (Phase 2)**: Depends on Setup completion - BLOCKS all user stories
- **User Stories (Phase 3+)**: All depend on Foundational phase completion
  - User stories can then proceed in parallel (if staffed)
  - Or sequentially in priority order (P1 → P2 → P3)
- **Polish (Final Phase)**: Depends on all desired user stories being complete

### User Story Dependencies

- **User Story 1 (P1)**: Can start after Foundational (Phase 2) - No dependencies on other stories
- **User Story 2 (P2)**: Can start after Foundational (Phase 2) - No dependencies on other stories
- **User Story 3 (P3)**: Can start after Foundational (Phase 2) - Depends on US2 (service configuration)

### Within Each User Story

- Flake configuration before testing
- Service configuration before auto-start setup
- Implementation before integration testing

### Parallel Opportunities

- All Setup tasks marked [P] can run in parallel
- All Foundational tasks can run in parallel (within Phase 2)
- Once Foundational phase completes, User Stories 1 and 2 can start in parallel
- User Story 3 depends on User Story 2 completion

### Parallel Example: User Stories 1 & 2

```bash
# Developer A: User Story 1 (Flake building)
Task: "Implement buildGoModule configuration in flake.nix"
Task: "Configure flake outputs to provide taskherald package"
Task: "Test flake build produces working binary"

# Developer B: User Story 2 (Home-manager module)
Task: "Create home-manager module structure in flake.nix"
Task: "Implement services.taskherald configuration options"
Task: "Test home-manager module configuration and service creation"
```

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup
2. Complete Phase 2: Foundational (CRITICAL - blocks all stories)
3. Complete Phase 3: User Story 1
4. **STOP and VALIDATE**: Test `nix build` produces working TaskHerald binary
5. Deploy/demo if ready

### Incremental Delivery

1. Complete Setup + Foundational → Foundation ready
2. Add User Story 1 → Test flake build independently → Deploy/Demo (MVP!)
3. Add User Story 2 → Test home-manager config independently → Deploy/Demo
4. Add User Story 3 → Test auto-start behavior → Deploy/Demo
5. Each story adds value without breaking previous stories

### Parallel Team Strategy

With multiple developers:

1. Team completes Setup + Foundational together
2. Once Foundational is done:
   - Developer A: User Story 1 (flake building)
   - Developer B: User Story 2 (home-manager module)
   - Developer C: User Story 3 (service auto-start) - after US2 complete

---

## Notes

- [P] tasks = different files, no dependencies
- [Story] label maps task to specific user story for traceability
- Each user story should be independently completable and testable
- Commit after each task or logical group
- Stop at any checkpoint to validate story independently
- This is a packaging feature - no Go code changes required, only Nix configuration
