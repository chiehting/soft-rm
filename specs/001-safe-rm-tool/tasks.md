# Tasks: 安全的 rm 命令工具 (Safe rm Tool)

**Input**: Design documents from `/specs/001-safe-rm-tool/`
**Prerequisites**: plan.md, spec.md

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure

- [X] T001 Initialize Go module: `go mod init soft-rm` in project root
- [X] T002 Add dependencies: `go get github.com/spf13/cobra github.com/spf13/viper`
- [X] T003 [P] Configure linting tools: Create `.golangci.yml` in project root

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story can be implemented

- [X] T004 Implement configuration loading logic in `pkg/config/config.go` to read from `~/.config/soft-rm/config.json`
- [X] T005 Create the main application entrypoint in `cmd/soft-rm/main.go`
- [X] T006 Set up the root Cobra command in `cmd/soft-rm/root.go` which will act as the `rm` wrapper

---

## Phase 3: User Story 1 - 安全刪除 (Safe Deletion) (Priority: P1) 🎯 MVP

**Goal**: Intercept `rm` command to move files to a trash directory instead of permanently deleting them.
**Independent Test**: Run `soft-rm <file>` and verify the file is moved to the configured trash directory with a timestamp.

### Implementation for User Story 1

- [X] T007 [US1] Implement the core "trash" logic in `internal/trash/trash.go` containing the function to move a file/directory to the trash.
- [X] T008 [US1] Integrate the trash logic into the root command's execution in `cmd/soft-rm/root.go`.
- [X] T009 [US1] Implement the background cleanup process spawner in `internal/cleaner/cleaner.go` as defined in `research.md`.
- [X] T010 [US1] Trigger the background cleaner from the root command in `cmd/soft-rm/root.go` after a file is successfully moved to the trash.

**Checkpoint**: At this point, User Story 1 should be fully functional. The `soft-rm` command should move files to the trash and trigger a background cleanup.

---

## Phase 4: User Story 2 - 自訂垃圾桶路徑 (Custom Trash Directory) (Priority: P2)

**Goal**: Allow users to view and update the trash path and retention days configuration.
**Independent Test**: Use the `config` subcommand to set a new `trash_path`, then use `soft-rm` and verify the file goes to the new path.

### Implementation for User Story 2

- [X] T011 [US2] Create the `config` subcommand and add it to the root command in `cmd/soft-rm/config.go`.
- [X] T012 [US2] Implement the `set` logic for `trash_path` and `retention_days` in `cmd/soft-rm/config.go`.
- [X] T013 [US2] Implement the `view` logic in `cmd/soft-rm/config.go` to display the current configuration.

**Checkpoint**: At this point, User Story 2 should be fully functional. Users can configure the tool.

---

## Phase 5: User Story 3 - 手動清理垃圾桶 (Manual Trash Cleanup) (Priority: P3)

**Goal**: Allow users to manually empty the trash directory.
**Independent Test**: Add files to the trash, run the `empty-trash` command, and verify the trash directory is empty.

### Implementation for User Story 3

- [X] T014 [US3] Create the `empty-trash` subcommand and add it to the root command in `cmd/soft-rm/empty_trash.go`.
- [X] T015 [US3] Implement the logic to permanently delete all items in the trash directory within the `empty-trash` command.

**Checkpoint**: All user stories should now be independently functional.

---

## Phase N: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories

- [X] T016 [P] Add unit tests for the trash logic in `internal/trash/trash_test.go`
- [X] T017 [P] Add unit tests for the configuration logic in `pkg/config/config_test.go`
- [X] T018 Add integration tests in `tests/integration/` to simulate CLI command executions.
- [X] T019 Create user documentation in `docs/usage.md`.
- [X] T020 Automate cross-platform builds for release using a GitHub Action workflow in `.github/workflows/release.yml`.

---

## Dependencies & Execution Order

- **Setup (Phase 1)** and **Foundational (Phase 2)** must be completed before any user stories.
- **User Story 1 (P1)** is the core MVP and can be implemented after the Foundational phase.
- **User Story 2 (P2)** and **User Story 3 (P3)** can be implemented in any order after User Story 1 is complete.

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup
2. Complete Phase 2: Foundational
3. Complete Phase 3: User Story 1
4. **STOP and VALIDATE**: Test User Story 1 independently (`go test ./...` and manual testing).
5. Deploy/demo the core `soft-rm` functionality.
