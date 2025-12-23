# Research: Background Process Management

## Decision

For the background cleanup process, we will use a standard, detached **goroutine**.

## Rationale

The user's requirement is to not block the shell while cleaning up the trash directory. A scheduled job (like a cron job) was explicitly rejected.

A simple, effective, and idiomatic Go solution is to spawn a goroutine to perform the cleanup task. To ensure it doesn't block the main `rm` command, we can launch it and immediately detach, letting it run to completion in the background.

The process will be:
1. The `soft-rm` command is executed.
2. After successfully moving the target file/directory to the trash, the application will spawn a new goroutine for the cleanup task.
3. The main application (the `rm` command wrapper) will exit immediately, returning control to the user's shell.
4. The background goroutine will:
    a. Acquire a lock file (e.g., `.trash/.lock`) to prevent multiple cleanup processes from running simultaneously. If the lock is already held, the goroutine will exit immediately.
    b. Scan the trash directory for files/directories older than the configured retention period.
    c. Delete the expired items.
    d. Release the lock.
    e. Exit.

This approach is lightweight, has no external dependencies, and directly satisfies the user's requirements.

## Alternatives Considered

- **Dedicated Daemon Process**: Running a separate, long-lived daemon process to watch the trash directory. This was rejected as it adds significant complexity to the installation and management of the tool, violating the "MVP First" and "Simplicity" principles.
- **Third-party Job Scheduling Libraries**: Using a library like `gocron`. This was deemed overkill for a simple, intermittent task and would add an unnecessary dependency. The standard library provides all the necessary tools.
