# Quickstart: TaskHerald Notifications

**Date**: 2025-10-29
**Feature**: TaskHerald Notifications

## Prerequisites

- Go 1.21 installed
- Taskwarrior installed and configured
- systemd (for service management)
- ntfy.sh account (optional, for custom server)

## Installation

1. Clone repository and build:
   ```bash
   git clone <repo>
   cd taskherald
   go build -o taskherald ./cmd/taskherald
   ```

2. Install binary:
   ```bash
   sudo cp taskherald /usr/local/bin/
   ```

## Configuration

Set environment variables:

```bash
export NTFY_SERVER=https://ntfy.sh  # Optional, defaults to ntfy.sh
export NTFY_TOPIC=mytasks           # Default notification topic
export TASKHERALD_INTERVAL=60       # Check interval in seconds
export TASKHERALD_STATE_FILE=/var/lib/taskherald/notifications.json  # State file path
```

## Running

### Manual Test

```bash
./taskherald
```

### As Systemd Service

1. Create service file `/etc/systemd/system/taskherald.service`:

   ```ini
   [Unit]
   Description=TaskHerald Notification Service
   After=network.target

   [Service]
   Type=simple
   User=taskherald
   Environment=NTFY_TOPIC=mytasks
   ExecStart=/usr/local/bin/taskherald
   Restart=always

   [Install]
   WantedBy=multi-user.target
   ```

2. Enable and start:
   ```bash
   sudo systemctl enable taskherald
   sudo systemctl start taskherald
   ```

## Task Setup

Add notification dates to tasks:

```bash
task add "Buy groceries" notification_date:2025-10-30T09:00:00
task add "Meeting" project:work priority:H notification_date:2025-10-30T14:00:00 ntfy_topic:work
```

## Verification

Check logs:
```bash
journalctl -u taskherald -f
```

Check state file:
```bash
cat /var/lib/taskherald/notifications.json
```