# Home-Manager Module Interface Contract

**Feature**: 001-nix-flake-homemanager
**Date**: 2025-10-29

## Module Path

`homeManagerModules.taskherald`

## Configuration Options

### services.taskherald

Type: `submodule`

#### enable (required)

- **Type**: `boolean`
- **Default**: `false`
- **Description**: Whether to enable the TaskHerald service

#### settings (optional)

- **Type**: `submodule`
- **Description**: Configuration settings for TaskHerald

##### settings.ntfy_topic (required when enabled, mutually exclusive with ntfy_topic_file)

- **Type**: `nullOr string`
- **Default**: `null`
- **Description**: Ntfy topic for notifications

##### settings.ntfy_topic_file (required when enabled, mutually exclusive with ntfy_topic)

- **Type**: `nullOr path`
- **Default**: `null`
- **Description**: Path to file containing ntfy topic
- **Validation**: Non-empty string

##### settings.ntfy_server (optional)

- **Type**: `string`
- **Default**: `"https://ntfy.sh"`
- **Description**: Ntfy server URL
- **Validation**: Valid URL format

##### settings.taskherald_interval (optional)

- **Type**: `integer`
- **Default**: `60`
- **Description**: Check interval in seconds
- **Validation**: Positive integer

## Generated Outputs

### Systemd User Service

**Service Name**: `taskherald`

**Service Configuration**:

```ini
[Unit]
Description=TaskHerald notification service
After=network.target

[Service]
Type=simple
ExecStart=${pkgs.taskherald}/bin/taskherald
Environment=NTFY_TOPIC=${config.services.taskherald.settings.ntfy_topic}
Environment=NTFY_SERVER=${config.services.taskherald.settings.ntfy_server}
Environment=TASKHERALD_INTERVAL=${toString config.services.taskherald.settings.taskherald_interval}
Restart=always

[Install]
WantedBy=default.target
```

### Environment Variables

The module sets the following environment variables for the service:

- `NTFY_TOPIC`: From `settings.ntfy_topic` (if set)
- `NTFY_TOPIC_FILE`: From `settings.ntfy_topic_file` (if set)
- `NTFY_SERVER`: From `settings.ntfy_server` (only if set)
- `TASKHERALD_INTERVAL`: From `settings.taskherald_interval` (only if set)

## Dependencies

- **Requires**: `pkgs.taskherald` (from flake packages)
- **Optional**: Taskwarrior CLI installation
- **Network**: Access to configured ntfy server

## Error Handling

- **Missing ntfy_topic or ntfy_topic_file**: Configuration fails if neither is set when enabled
- **Both ntfy_topic and ntfy_topic_file set**: Configuration fails if both are set
- **Invalid ntfy_topic**: Configuration fails if empty when set
- **Invalid ntfy_topic_file**: Configuration fails if file doesn't exist or is unreadable
- **Invalid ntfy_server**: Configuration fails if invalid URL format
- **Invalid taskherald_interval**: Configuration fails if not positive integer
- **Missing taskherald package**: Build fails if flake package not available
