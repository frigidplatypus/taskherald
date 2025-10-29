# Quickstart: Nix Flake & Home-Manager Setup

**Feature**: 001-nix-flake-homemanager
**Date**: 2025-10-29

## Prerequisites

- Nix with flakes enabled
- Home-manager installed
- Taskwarrior installed and configured

## Installation

### 1. Add TaskHerald flake to your configuration

Add the TaskHerald flake to your `flake.nix`:

```nix
{
  inputs = {
    # ... your other inputs
    taskherald.url = "github:your-org/taskherald";  # Replace with actual repo
  };

  outputs = { self, nixpkgs, taskherald, ... }: {
    # ... your outputs
  };
}
```

### 2. Configure home-manager

Add TaskHerald to your home-manager configuration:

```nix
{ config, pkgs, ... }:

{
  imports = [
    # ... your other imports
  ];

  # Import the TaskHerald home-manager module
  home-manager.users.youruser = { config, ... }: {
    imports = [
      # ... your other imports
    ];

    # TaskHerald service configuration
    services.taskherald = {
      enable = true;
      settings = {
        ntfy_topic = "my-task-notifications";
        ntfy_server = "https://ntfy.sh";  # Optional
        taskherald_interval = 60;          # Optional
      };
    };
  };
}
```

### 3. Apply configuration

```bash
# Rebuild and switch to new configuration
home-manager switch

# Or if using NixOS with flakes:
sudo nixos-rebuild switch
```

## Verification

### Check service status

```bash
# Check if service is running
systemctl --user status taskherald

# View service logs
journalctl --user -u taskherald -f
```

### Test notifications

Create a test task with a notification:

```bash
# Create task that notifies in 1 minute
task add "Test notification" notification_date:1min

# Wait for notification on your ntfy topic
```

## Configuration Options

| Option                | Required | Default           | Description                                      |
| --------------------- | -------- | ----------------- | ------------------------------------------------ |
| `enable`              | Yes      | `false`           | Enable TaskHerald service                        |
| `ntfy_topic`          | Yes*     | -                 | Ntfy topic for notifications (mutually exclusive with ntfy_topic_file) |
| `ntfy_topic_file`     | Yes*     | -                 | Path to file containing ntfy topic (mutually exclusive with ntfy_topic) |
| `ntfy_server`         | No       | `https://ntfy.sh` | Ntfy server URL                                  |
| `taskherald_interval` | No       | `60`              | Check interval in seconds                        |

\* Exactly one of `ntfy_topic` or `ntfy_topic_file` must be specified.

## Troubleshooting

### Service not starting

Check service status:

```bash
systemctl --user status taskherald
journalctl --user -u taskherald
```

### No notifications

1. Verify Taskwarrior is installed: `task version`
2. Check ntfy topic configuration
3. Verify network connectivity to ntfy server
4. Check TaskHerald logs for errors

### Build failures

Ensure flakes are enabled in Nix:

```bash
# Check nix version and flake support
nix --version
nix flake --help
```

## Advanced Usage

### Custom ntfy server

```nix
services.taskherald = {
  enable = true;
  settings = {
    ntfy_topic = "work-tasks";
    ntfy_server = "https://ntfy.yourdomain.com";
  };
};
```

### Frequent checks

```nix
services.taskherald = {
  enable = true;
  settings = {
    ntfy_topic = "urgent-tasks";
    taskherald_interval = 30;  # Check every 30 seconds
  };
};
```
