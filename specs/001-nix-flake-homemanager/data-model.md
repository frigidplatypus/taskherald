# Data Model: Nix Flake & Home-Manager Support

**Feature**: 001-nix-flake-homemanager
**Date**: 2025-10-29

## Configuration Schema

### Home-Manager Module Configuration

```nix
services.taskherald = {
  enable = true;  # Required: enables the service
  settings = {
    ntfy_topic = "my-tasks";        # Required: notification topic
    ntfy_server = "https://ntfy.sh";  # Optional: defaults to ntfy.sh
    taskherald_interval = 60;       # Optional: check interval in seconds
  };
};
```

### Validation Rules

- `enable`: Boolean, required
- `ntfy_topic`: String, required when enabled, no default
- `ntfy_server`: String, optional, must be valid URL if provided
- `taskherald_interval`: Integer, optional, must be positive if provided

### Environment Variables Mapping

The home-manager module translates configuration to environment variables:

```
NTFY_TOPIC=${settings.ntfy_topic}
NTFY_SERVER=${settings.ntfy_server}  # Only if set
TASKHERALD_INTERVAL=${toString settings.taskherald_interval}  # Only if set
```

## Service Configuration

### Systemd User Service

- **Service Name**: `taskherald`
- **Type**: `simple`
- **ExecStart**: Path to taskherald binary from flake
- **Environment**: Variables set from home-manager config
- **Restart**: `always` (automatic recovery)
- **WantedBy**: `default.target` (starts on login)

### Dependencies

- **Requires**: Network connectivity for ntfy.sh
- **Wants**: Taskwarrior installation
- **After**: Network setup

## Build Configuration

### Flake Inputs

```nix
inputs = {
  nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  flake-utils.url = "github:numtide/flake-utils";
};
```

### Package Definition

- **buildGoModule**: Standard Go build approach
- **vendorHash**: Computed from go.mod/go.sum
- **meta**: Standard package metadata (description, license, maintainers)

## State Transitions

### Service Lifecycle

1. **Disabled** → **Enabled**: Home-manager applies configuration
2. **Enabled** → **Running**: Systemd starts service on login
3. **Running** → **Stopped**: Systemd stops service on logout
4. **Configuration Change** → **Restart**: Service restarts with new environment

### Build States

1. **Source** → **Built**: `nix build` compiles Go code
2. **Built** → **Installed**: Home-manager installs package and service
3. **Installed** → **Running**: Service starts with configuration
