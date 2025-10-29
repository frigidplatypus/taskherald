# Feature Specification: TaskHerald Notifications

**Feature Branch**: `001-taskherald-notifications`  
**Created**: 2025-10-29  
**Status**: Draft  
**Input**: User description: "Taskherald is a single user service that regularly retrieves a users taskwarrior tasks and sends notifications via ntfy.sh when a taskwarrior UDA (notification_date) is the current date/time. It also has the following features:

- When a notification is sent the UDA taskherald_notified is updated with the current timestamp.
- When Taskherald starts and there are previously unsent messages with a notification_date, a summary message is sent. The timespan of previously sent is no more than 7 days ago.
- When Taskherald starts, it logs up to 10 tasks with a notification_date set to the console to verify the service can reach the local taskwarrior tasks. If no tasks are found, that is logged as well.
- A custom ntfy.sh server may be specified
- A custom ntfy-topic may be specified
- The UDA ntfy_topic can be set on a task to override the topic the message will be delivered to for that task
- The interval at with "task export" runs is 60s
- Taskherald does not retry sending failed messages. It is logged."

## User Scenarios & Testing _(mandatory)_

<!--
  IMPORTANT: User stories should be PRIORITIZED as user journeys ordered by importance.
  Each user story/journey must be INDEPENDENTLY TESTABLE - meaning if you implement just ONE of them,
  you should still have a viable MVP (Minimum Viable Product) that delivers value.

  Assign priorities (P1, P2, P3, etc.) to each story, where P1 is the most critical.
  Think of each story as a standalone slice of functionality that can be:
  - Developed independently
  - Tested independently
  - Deployed independently
  - Demonstrated to users independently
-->

### User Story 1 - Receive Task Notifications (Priority: P1)

As a user, I want to receive notifications via ntfy.sh when a task's notification_date matches the current date/time, so that I am reminded of important tasks.

**Why this priority**: This is the core functionality delivering immediate value for task reminders.

**Independent Test**: Can be tested by setting a task's notification_date to the current time, running the service, and verifying the notification is sent to ntfy.sh.

**Acceptance Scenarios**:

1. **Given** a task with notification_date set to the current time, **When** the service checks tasks, **Then** a notification is sent to the default ntfy topic.
2. **Given** a task with notification_date set to the current time and ntfy_topic UDA set, **When** the service checks tasks, **Then** a notification is sent to the task's specified topic.
3. **Given** a task with notification_date in the past, **When** the service checks tasks, **Then** no notification is sent.

---

### User Story 2 - Startup Summary Notifications (Priority: P2)

As a user, I want to receive a summary notification on service startup for any previously unsent notifications within the last 7 days, so that I don't miss important reminders while the service was down.

**Why this priority**: Ensures continuity of notifications after service interruptions.

**Independent Test**: Can be tested by stopping the service, creating tasks with past notification_dates within 7 days, restarting the service, and verifying the summary notification is sent.

**Acceptance Scenarios**:

1. **Given** the service was stopped and tasks have notification_date within the last 7 days without taskherald_notified, **When** the service starts, **Then** a summary notification is sent listing those tasks.
2. **Given** tasks have notification_date older than 7 days, **When** the service starts, **Then** no summary notification is sent for those tasks.
3. **Given** tasks already have taskherald_notified set, **When** the service starts, **Then** no summary notification is sent for those tasks.

---

### User Story 3 - Startup Task Verification (Priority: P3)

As a user, I want the service to log up to 10 tasks with notification_date on startup to verify it can access Taskwarrior, so that I know the service is properly configured.

**Why this priority**: Provides confidence in service setup and Taskwarrior integration.

**Independent Test**: Can be tested by starting the service and checking console logs for task details or a message indicating no tasks found.

**Acceptance Scenarios**:

1. **Given** tasks exist with notification_date, **When** the service starts, **Then** up to 10 tasks are logged to the console.
2. **Given** no tasks have notification_date, **When** the service starts, **Then** a message is logged indicating no tasks found.

---

### User Story 4 - Custom Ntfy Configuration (Priority: P4)

As a user, I want to configure custom ntfy.sh server and default topic, so that I can use my own ntfy instance or organize notifications.

**Why this priority**: Allows customization for different deployment scenarios.

**Independent Test**: Can be tested by configuring custom server/topic, sending a notification, and verifying it arrives at the custom destination.

**Acceptance Scenarios**:

1. **Given** a custom ntfy server is configured, **When** a notification is sent, **Then** it is sent to the custom server.
2. **Given** a custom default topic is configured, **When** a notification is sent without task-specific topic, **Then** it is sent to the custom topic.

---

## Clarifications

### Session 2025-10-29

- Q: How is the service configured (custom server, topic, interval, state file path)? → A: Environment variables
- Q: What is the structure for tracking sent notifications to avoid duplicates? → A: Via taskherald_notified UDA in Taskwarrior
- Q: How should the notification message handle missing project or due fields? → A: Use empty string

## Edge Cases

<!--
  ACTION REQUIRED: The content in this section represents placeholders.
  Fill them out with the right edge cases.
-->

- What happens if Taskwarrior is not installed or accessible?
- What happens if ntfy.sh server is unreachable?
- What happens if project or due is missing from a task? Use empty string in notification message.
- What happens if multiple tasks have the same notification_date?
- What happens if notification_date is in the future?

## Requirements _(mandatory)_

<!--
  ACTION REQUIRED: The content in this section represents placeholders.
  Fill them out with the right functional requirements.
-->

### Functional Requirements

- **FR-001**: System MUST check Taskwarrior tasks every 60 seconds using "task export".
- **FR-002**: System MUST send notification via ntfy.sh when task notification_date equals current date/time.
- **FR-003**: System MUST update taskherald_notified UDA with current timestamp after sending notification.
- **FR-004**: System MUST send summary notification on startup for unsent notifications within last 7 days.
- **FR-005**: System MUST log up to 10 tasks with notification_date on startup; log message if none found.
- **FR-006**: System MUST support custom ntfy.sh server configuration.
- **FR-007**: System MUST support custom default ntfy topic configuration.
- **FR-008**: System MUST use task's ntfy_topic UDA to override notification topic if set.
- **FR-009**: System MUST log failed notification sends without retrying.
- **FR-010**: System MUST read configuration from environment variables: NTFY_SERVER, NTFY_TOPIC, TASKHERALD_INTERVAL, TASKHERALD_STATE_FILE.
- **FR-011**: System MUST use empty string for missing project or due fields in notification message format: "Project:%project% %description% Due:%due%".

### Key Entities _(include if feature involves data)_

- **Task**: Represents a Taskwarrior task with UDAs: notification_date (datetime), taskherald_notified (timestamp), ntfy_topic (string)

## Success Criteria _(mandatory)_

### Measurable Outcomes

- **SC-001**: Notifications are sent within 5 seconds of notification_date matching current time.
- **SC-002**: Service starts up and sends summary notifications within 30 seconds.
- **SC-003**: Service correctly logs task verification on startup in under 10 seconds.
- **SC-004**: 99% of notifications are delivered successfully when ntfy.sh is available.
- **SC-005**: No duplicate notifications sent for the same task within a 60-second interval.

## Constitution Alignment

_GATE: Must verify compliance with TaskHerald Constitution v1.1.0 before finalizing spec._

- **Service Reliability**: Service runs as systemd service with automatic recovery; notifications ensure continuous task awareness.
- **CLI Contract**: All Taskwarrior interactions use `task export` command with proper parsing.
- **Security**: Operates with least privilege; notifications contain only necessary task information.
- **Observability**: All actions logged with timestamps; errors from failed sends reported.
- **Simplicity**: Core notification logic remains minimal; no complex retry mechanisms.
- **Deployment Constraints**: Uses Go, systemd, Taskwarrior CLI, and ntfy.sh; network access limited to ntfy.sh.
