# Feature Specification: 安全的 rm 命令工具 (Safe rm Tool)

**Feature Branch**: `001-safe-rm-tool`
**Created**: 2025-12-23
**Status**: Draft
**Input**: User description: "本專案為建立一個安全的 rm 命令工具。用戶可以做設定路徑跟日期，例如 .trash 跟 30d，當執行 rm 命令時不會直接刪除目標，而是先將目標移動到指定的路徑下，例如 .trash/yyyyddmmss_target。並且在背景刪除指定日期前的檔案或目錄，例如刪除 30d 前的檔案或目錄。"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - 安全刪除 (Safe Deletion) (Priority: P1)

使用者希望使用 `rm` 命令時，檔案會被移至垃圾桶資料夾，而非永久刪除，以便能救回誤刪的檔案。
(As a user, I want to use `rm` to move files to a trash directory instead of deleting them permanently, so that I can recover accidentally deleted files.)

**Why this priority**: 這是此工具最核心的功能。 (This is the core functionality of the tool.)

**Independent Test**: 可以透過執行 `rm` 命令並檢查檔案是否已移至垃圾桶資料夾來獨立測試此功能。 (This can be tested independently by running the `rm` command and checking if the file is moved to the trash directory.)

**Acceptance Scenarios**:

1.  **Given** 一個檔案 `my_file.txt` (a file `my_file.txt`),
    **When** 我執行 `rm my_file.txt` (I run `rm my_file.txt`),
    **Then** 該檔案會被移動到 `.trash/YYYYMMDDHHMMSS_my_file.txt` (the file is moved to `.trash/YYYYMMDDHHMMSS_my_file.txt`).
2.  **Given** 一個資料夾 `my_dir` (a directory `my_dir`),
    **When** 我執行 `rm -r my_dir` (I run `rm -r my_dir`),
    **Then** 該資料夾會被移動到 `.trash/YYYYMMDDHHMMSS_my_dir` (the directory is moved to `.trash/YYYYMMDDHHMMSS_my_dir`).

---

### User Story 2 - 自訂垃圾桶路徑 (Custom Trash Directory) (Priority: P2)

使用者希望能自訂垃圾桶資料夾的路徑，以便將刪除的檔案存放在所選的位置。
(As a user, I want to configure the trash directory path, so that I can store deleted files in a location of my choice.)

**Why this priority**: 提供使用者自訂的彈性。 (Provides flexibility for the user.)

**Independent Test**: 可以透過設定垃圾桶路徑、執行 `rm` 命令，並檢查檔案是否已移至新的垃圾桶路徑來獨立測試此功能。 (This can be tested independently by setting the trash path, running `rm`, and checking if the file is in the new location.)

**Acceptance Scenarios**:

1.  **Given** 我已將垃圾桶路徑設定為 `/my/custom/trash` (I have configured the trash path to `/my/custom/trash`),
    **When** 我執行 `rm my_file.txt` (I run `rm my_file.txt`),
    **Then** 該檔案會被移動到 `/my/custom/trash/YYYYMMDDHHMMSS_my_file.txt` (the file is moved to `/my/custom/trash/YYYYMMDDHHMMSS_my_file.txt`).

---

### User Story 3 - 自動清理垃圾桶 (Automatic Trash Cleanup) (Priority: P3)

使用者希望垃圾桶內的檔案在經過一段可設定的時間後能被自動刪除，以自動管理磁碟空間。
(As a user, I want files in the trash directory to be automatically deleted after a configurable period, so that my disk space is managed automatically.)

**Why this priority**: 這是確保工具不會無限地消耗磁碟空間的關鍵功能。 (This is a key feature to ensure the tool does not consume disk space indefinitely.)

**Independent Test**: 可以透過設定保留期限，並在垃圾桶中放置不同時間戳記的檔案，然後驗證背景清理工作是否只刪除過期的檔案來獨立測試此功能。 (This can be tested independently by configuring the retention period, placing files with different timestamps in the trash, and verifying that the background job only deletes expired files.)

**Acceptance Scenarios**:

1.  **Given** 我已將保留期限設定為 30 天，且垃圾桶中有一個超過 30 天的檔案 (I have configured the retention period to 30 days, and there is a file in the trash directory older than 30 days),
    **When** 背景清理工作執行時 (the background process runs),
    **Then** 該檔案會被永久刪除 (that file is permanently deleted).
2.  **Given** 我已將保留期限設定為 30 天，且垃圾桶中有一個未滿 30 天的檔案 (I have configured the retention period to 30 days, and there is a file in the trash directory newer than 30 days),
    **When** 背景清理工作執行時 (the background process runs),
    **Then** 該檔案不會被刪除 (that file is not deleted).

---

### Edge Cases

- 當垃圾桶資料夾不存在時會發生什麼事？ (What happens if the trash directory does not exist?)
- 如果目標檔案已經存在於垃圾桶中，該如何處理？ (How to handle if a file with the same name already exists in the trash?)
- 如果使用者沒有權限將檔案移動到垃圾桶資料夾，該如何處理？ (What if the user does not have permission to move the file to the trash directory?)
- 如果背景清理工作在刪除檔案時失敗，該如何處理？ (What if the background cleanup job fails while deleting a file?)

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: 系統必須攔截 `rm` 命令。 (The system MUST intercept the `rm` command.)
- **FR-002**: 系統必須將指定的檔案或資料夾移動到設定的垃圾桶資料夾。 (The system MUST move the specified file or directory to a configured trash directory.)
- **FR-003**: 系統必須在移動的檔案/資料夾名稱前加上時間戳記。 (The system MUST prepend a timestamp to the moved file/directory name.)
- **FR-004**: 系統必須允許使用者設定垃圾桶資料夾的路徑。 (The system MUST allow users to configure the trash directory path.)
- **FR-005**: 系統必須允許使用者設定保留期限（以天為單位）。 (The system MUST allow users to configure a retention period (in days).)
- **FR-006**: 系統必須有一個背景程序，定期掃描垃圾桶資料夾。 (The system MUST have a background process that periodically scans the trash directory.)
- **FR-007**: 背景程序必須刪除早於設定保留期限的檔案/資料夾。 (The background process MUST delete files/directories older than the configured retention period.)
- **FR-008**: 設定必須儲存在使用者可設定的位置（例如 `~/.config/soft-rm/config.json`）。 (Configuration MUST be stored in a user-configurable location (e.g., `~/.config/soft-rm/config.json`).)
- **FR-009**: 系統必須提供一個命令來檢視目前的設定。 (The system MUST provide a command to view the current configuration.)
- **FR-010**: 系統必須提供一個手動清空垃圾桶的指令。 (The system MUST provide a command to manually empty the trash.)

### Key Entities *(include if feature involves data)*

- **設定 (Configuration)**: 儲存 `trash_path` 和 `retention_days`。 (Stores `trash_path` and `retention_days`.)
- **垃圾桶項目 (TrashedItem)**: 代表垃圾桶中的一個檔案或資料夾，包含其原始路徑和刪除時間戳記。 (Represents a file or directory in the trash, with its original path and deletion timestamp.)

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: 99% 的 `rm` 命令執行後，檔案會被移至垃圾桶，而非被刪除。 (99% of `rm` command executions result in the file being moved to the trash, not deleted.)
- **SC-002**: 超過保留期限的檔案會在過期後 24 小時內被背景程序刪除。 (Files older than the retention period are deleted by the background process within 24 hours of expiring.)
- **SC-003**: `rm` 包裝器的效能延遲應小於 50 毫秒。 (The performance overhead of the `rm` wrapper is less than 50ms.)