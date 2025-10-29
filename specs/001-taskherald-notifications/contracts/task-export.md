# Contract: Taskwarrior CLI Export

**Date**: 2025-10-29
**Purpose**: Define interface for retrieving tasks with future notification dates

## Command

```
task notification_date.after:now export
```

**Parameters**:
- `notification_date.after:now`: Filter tasks where notification_date is after current time

**Output**: JSON array of task objects

## Response Format

```json
[
  {
    "uuid": "string",
    "description": "string",
    "project": "string (optional)",
    "due": "ISO datetime string (optional)",
    "tags": ["string"],
    "priority": "H|M|L (optional)",
    "udas": {
      "notification_date": "ISO datetime string",
      "taskherald_notified": "ISO datetime string (optional)",
      "ntfy_topic": "string (optional)"
    }
  }
]
```

## Error Handling

- Command exits with code 0 on success
- Empty array if no tasks match
- Command fails if Taskwarrior not configured (exit code >0)

## Assumptions

- Taskwarrior installed and configured
- User has read access to task database
- JSON output format enabled