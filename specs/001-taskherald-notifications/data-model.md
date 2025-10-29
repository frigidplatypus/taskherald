# Data Model: TaskHerald Notifications

**Date**: 2025-10-29
**Feature**: TaskHerald Notifications

## Entities

### Task
Represents a Taskwarrior task retrieved via `task export`.

**Fields**:
- `uuid` (string): Unique identifier from Taskwarrior
- `description` (string): Task description
- `project` (string, optional): Project name
- `due` (datetime, optional): Due date
- `tags` (array of strings): Task tags
- `priority` (string, enum: H, M, L, or empty): Task priority
- `udas` (object): User-defined attributes
  - `notification_date` (datetime): When to send notification
  - `taskherald_notified` (datetime, optional): When notification was sent
  - `ntfy_topic` (string, optional): Override notification topic

**Validation Rules**:
- `uuid` must be valid UUID format
- `notification_date` must be valid datetime
- `priority` must be one of: H, M, L, or empty string
- `tags` array elements must be non-empty strings

**State Transitions**:
- Unnotified: `taskherald_notified` is null
- Notified: `taskherald_notified` set to current timestamp

### SentNotification
Tracks notifications sent to avoid duplicates.

**Fields**:
- `uuid` (string): Task UUID
- `notified_at` (datetime): When notification was sent

**Validation Rules**:
- `uuid` must match a valid Task UUID
- `notified_at` must be valid datetime

**Relationships**:
- One-to-one with Task (by UUID)
- Stored in JSON file as array of SentNotification objects