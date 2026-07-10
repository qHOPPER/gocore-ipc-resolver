# gocore-ipc-resolver [v1.0]

Deterministic IPC resolver for systemd-resolved. Eliminates UDP overhead and external dependencies.

## Features
- Low-Level IPC: Direct communication via AF_UNIX (binary wire format).
- Zero-Dependency: No external libraries (uses golang.org/x/sys/unix).
- Kernel-Gatekeeping: SO_PEERCRED UID verification.
- Hardened Bastion: Cgroups v2, Seccomp, and Namespacing.
- Graceful Shutdown: Idempotent signal handling.

## Installation
1. Apply kernel hardening (`sysctl.hardening`).
2. Deploy systemd override (`systemd/override.conf`).
3. Build: `CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o gocore-ipc-resolver main.go`
