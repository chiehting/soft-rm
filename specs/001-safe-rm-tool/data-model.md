# Data Model: Configuration

The configuration for the `soft-rm` tool will be stored in a JSON file located at `~/.config/soft-rm/config.json`.

## Configuration Schema

The JSON file will have the following structure:

```json
{
  "trash_path": "/path/to/your/trash",
  "retention_days": 30
}
```

### Fields

- **`trash_path`**
  - **Type**: `string`
  - **Description**: The absolute path to the directory where deleted files will be moved.
  - **Default**: `~/.trash`
  - **Validation**: Must be a valid, writable directory path.

- **`retention_days`**
  - **Type**: `integer`
  - **Description**: The number of days to keep files in the trash before they are permanently deleted by the background cleanup process.
  - **Default**: `30`
  - **Validation**: Must be a non-negative integer. A value of `0` would mean files are deleted by the cleanup process almost immediately.
