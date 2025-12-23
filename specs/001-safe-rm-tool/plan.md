# Implementation Plan: 安全的 rm 命令工具 (Safe rm Tool)

**Branch**: `001-safe-rm-tool` | **Date**: 2025-12-23 | **Spec**: [spec.md](./spec.md)
**Input**: Feature specification from `/specs/001-safe-rm-tool/spec.md`

**Note**: This template is filled in by the `/speckit.plan` command.

## Summary

本專案旨在建立一個安全的 `rm` 命令工具，它會將刪除的檔案移至垃圾桶，並在背景自動清理過期檔案，以防止意外的永久性資料遺失。
(This project aims to create a safe `rm` command tool that moves deleted files to a trash directory and automatically cleans up expired files in the background to prevent accidental permanent data loss.)

## Technical Context

**Language/Version**: Golang (1.2x)
**Primary Dependencies**:
- **CLI Framework**: `github.com/spf13/cobra`
- **Configuration**: `github.com/spf13/viper`
- **Background Processing**: Standard library `goroutines`
**Storage**: Filesystem for trash directory and a JSON file for configuration (`~/.config/soft-rm/config.json`).
**Testing**: Standard library `testing`
**Target Platform**: Cross-platform (Linux, macOS, Windows)
**Project Type**: single/cli
**Performance Goals**:
- File move operation overhead < 50ms.
- Background cleanup should utilize minimal system resources and not interfere with foreground tasks.
**Constraints**:
- The tool MUST be installed as a wrapper or alias for the standard `rm` command.
- The background cleanup process MUST run detached from the user's shell session.
- The final product should be a single, self-contained binary.

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

- **主要開發語言 (Primary Development Language):** Yes, the implementation will be in Golang.
- **執行效率 (Execution Efficiency):** Yes, the use of compiled Go and lightweight goroutines is designed for high performance and low overhead.
- **最小可行性產品 (Minimum Viable Product):** Yes, the plan focuses on the core features of safe deletion and background cleanup.
- **語言 (Language):** Yes, all user-facing output will be in Traditional Chinese.

## Project Structure

### Documentation (this feature)

```text
specs/001-safe-rm-tool/
├── plan.md              # This file
├── research.md          # Background process research
├── data-model.md        # Configuration file structure
├── quickstart.md        # User guide
└── tasks.md             # Implementation tasks (to be generated)
```

### Source Code (repository root)
```text
# Single project (DEFAULT)
cmd/soft-rm/
├── main.go
├── root.go
├── config.go
└── empty_trash.go
internal/
├──- trash/
│   └── trash.go
└──- cleaner/
    └── cleaner.go
pkg/
└──- config/
    └── config.go
tests/
├── integration/
└── unit/
```

**Structure Decision**: The standard Go CLI project structure is adopted to separate command definitions (`cmd`), internal business logic (`internal`), and reusable packages (`pkg`).

## Complexity Tracking

> **Fill ONLY if Constitution Check has violations that must be justified**

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| N/A       | -          | -                                   |
