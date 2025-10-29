# TaskHerald

A single-user systemd service that monitors Taskwarrior tasks and sends notifications via ntfy.sh when tasks reach their notification_date.

## Features

- Regularly checks Taskwarrior tasks every 60 seconds
- Sends ntfy.sh notifications for due tasks
- Supports custom ntfy servers and topics
- Per-task topic overrides via UDA
- Startup summary for missed notifications
- Priority and tag mapping to ntfy
- Graceful shutdown

## Installation

1. Build the binary:
   ```bash
   go build -o taskherald ./src
   ```

2. Install:
   ```bash
   sudo cp taskherald /usr/local/bin/
   sudo mkdir -p /var/lib/taskherald
   sudo useradd -r -s /bin/false taskherald
   sudo chown taskherald:taskherald /var/lib/taskherald
   ```

3. Configure systemd:
   ```bash
   sudo cp systemd/taskherald.service /etc/systemd/system/
   sudo systemctl enable taskherald
   sudo systemctl start taskherald
   ```

## Configuration

TaskHerald is configured via environment variables. Set them in the systemd service file or globally.

### Environment Variables

| Variable | Description | Default | Example |
|----------|-------------|---------|---------|
| `NTFY_SERVER` | URL of the ntfy server to send notifications to | `https://ntfy.sh` | `https://ntfy.example.com` |
| `NTFY_TOPIC` | Default topic for notifications | `taskherald` (or `taskherald-RANDOM` for ntfy.sh) | `my-tasks` |
| `TASKHERALD_INTERVAL` | How often to check for due tasks (in seconds) | `60` | `30` |

### Setting Environment Variables

For systemd service, edit `/etc/systemd/system/taskherald.service`:

```ini
[Service]
Environment=NTFY_SERVER=https://ntfy.example.com
Environment=NTFY_TOPIC=my-tasks
Environment=TASKHERALD_INTERVAL=30
```

Then reload and restart:

```bash
sudo systemctl daemon-reload
sudo systemctl restart taskherald
```

## Task Setup

Add notification dates to tasks:

```bash
task add "Buy groceries" notification_date:2025-10-30T09:00:00
task add "Meeting" project:work priority:H notification_date:2025-10-30T14:00:00 ntfy_topic:work
```

## Usage

The service runs automatically. Check logs:

```bash
journalctl -u taskherald -f
```

## Development

- Source in `src/`
- Build with `go build ./src`
- Test with `go test ./src`