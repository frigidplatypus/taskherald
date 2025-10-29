# Research: TaskHerald Notifications

**Date**: 2025-10-29
**Feature**: TaskHerald Notifications

## Go Version

**Decision**: Go 1.21
**Rationale**: Latest stable LTS version with good performance and ecosystem support. Provides modern language features while maintaining stability.
**Alternatives considered**: Go 1.20 (older LTS), Go 1.22 (newer but less tested in production)

## Ntfy Client Library

**Decision**: github.com/binwiederhier/ntfy (official Go client)
**Rationale**: Official client library maintained by ntfy project, supports all features including tags, priority, custom servers.
**Alternatives considered**: HTTP client with manual requests (more complex), other unofficial libraries (less maintained)

## Task Export JSON Format

**Decision**: Standard Taskwarrior JSON export format
**Rationale**: Taskwarrior's `task export` produces consistent JSON with all task fields including UDAs. UUID as primary key, notification_date and other UDAs accessible.
**Alternatives considered**: Custom parsing of `task` command output (error-prone), database access (not supported by Taskwarrior)

## Sent Notifications JSON File Structure

**Decision**: JSON array of objects with UUID and timestamp
**Rationale**: Simple, human-readable, easy to parse in Go. Tracks which tasks have been notified to avoid duplicates.
**Alternatives considered**: SQLite database (overkill for simple tracking), in-memory only (loses state on restart)