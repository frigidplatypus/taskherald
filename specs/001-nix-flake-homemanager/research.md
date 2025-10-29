# Research: Nix Flake & Home-Manager Support

**Feature**: 001-nix-flake-homemanager
**Date**: 2025-10-29

## Research Tasks Completed

### Task: Research flake.nix structure for Go projects

**Decision**: Use standard Nixpkgs Go build approach with buildGoModule
**Rationale**: Standard approach for Go projects in Nix, provides reproducible builds and dependency management
**Alternatives considered**:

- Custom build script: More complex, less maintainable
- Use existing Go tooling: Doesn't provide Nix benefits like reproducibility

### Task: Research home-manager module structure

**Decision**: Create module in flake outputs.homeManagerModules.taskherald
**Rationale**: Standard home-manager module location, allows importing via flake
**Alternatives considered**:

- Separate home-manager module repo: Unnecessary complexity for single feature
- Inline module in configuration: Less reusable

### Task: Research user-level systemd service configuration

**Decision**: Use home-manager's systemd.user.services option
**Rationale**: Proper user-level service management with automatic start/stop on login/logout
**Alternatives considered**:

- System-level service: Requires root, conflicts with user-level design
- Manual service files: Less declarative, harder to manage

### Task: Research Nix flake best practices for packages

**Decision**: Follow nixpkgs conventions with proper meta attributes
**Rationale**: Ensures discoverability and follows community standards
**Alternatives considered**:

- Minimal flake: Less useful for users
- Custom structure: Confusing for Nix users

## Implementation Approach

Based on research, the implementation will:

1. Create `flake.nix` with:

   - Go build using `buildGoModule`
   - Output package as `packages.x86_64-linux.default`
   - Home-manager module in `homeManagerModules.taskherald`

2. Home-manager module will:

   - Accept `services.taskherald` configuration
   - Create systemd user service
   - Set environment variables from config
   - Enable service by default when configured

3. Follow Nix community patterns for:
   - Proper meta attributes (description, license, maintainers)
   - Reproducible builds
   - Cross-platform support (x86_64-linux only as specified)
