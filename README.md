# soft-rm

A safer `rm` command that moves files to a trash directory instead of deleting them permanently.

This tool provides a safety net, moving specified files and directories to a trash folder instead of immediate, irreversible deletion. It also includes functionality to manage this trash folder, such as emptying it automatically based on a retention policy or manually.

## Installation

To install `soft-rm`, use the `go install` command:

```sh
go install github.com/chiehting/soft-rm@latest
```

## Usage

To use `soft-rm`, simply pass the files or directories you wish to "delete" as arguments.

```sh
soft-rm [file1.txt] [directory1] ...
```

### Flags

-   `-f`, `--force`: Ignores and suppresses errors if a specified file does not exist.
-   `-r`, `--recursive`: This flag is ignored and is included only for compatibility with the standard `rm` command's syntax.

### Commands

#### empty-trash
Permanently deletes all items from the trash directory.

```sh
soft-rm empty-trash
```

## Configuration

`soft-rm` creates a configuration file at `~/.config/soft-rm/config.json` upon its first run.

The default configuration is:
```json
{
  "trash_path": "~/.trash",
  "retention_days": 30
}
```

-   **`trash_path`**: The directory where deleted files are stored.
-   **`retention_days`**: The number of days to keep files in the trash before they are automatically deleted. The cleanup process runs in the background.

## Development

To build the project from source:

```sh
go build -o soft-rm ./cmd/soft-rm
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request.

1.  Fork the repository.
2.  Create your feature branch (`git checkout -b feature/AmazingFeature`).
3.  Commit your changes (`git commit -m 'Add some AmazingFeature'`).
4.  Push to the branch (`git push origin feature/AmazingFeature`).
5.  Open a pull request.
