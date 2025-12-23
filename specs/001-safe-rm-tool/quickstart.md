# Quickstart: soft-rm

This guide explains how to install, configure, and use the `soft-rm` tool.

## Installation

1.  Download the latest binary for your operating system from the releases page.
2.  Place the binary in a directory that is in your system's `PATH` (e.g., `/usr/local/bin`).
3.  Create an alias for the `rm` command in your shell's configuration file (e.g., `.bashrc`, `.zshrc`) to point to the `soft-rm` executable:

    ```sh
    alias rm='/path/to/your/soft-rm'
    ```

4.  Reload your shell configuration (e.g., `source ~/.bashrc`) or open a new terminal.

## Configuration

The `soft-rm` tool can be configured using the `config` subcommand.

### View Configuration

To see the current configuration, run:

```sh
soft-rm config view
```

### Set Configuration

To set the trash directory path:

```sh
soft-rm config set trash_path /my/custom/trash
```

To set the retention period (in days):

```sh
soft-rm config set retention_days 60
```

The configuration is stored at `~/.config/soft-rm/config.json`.

## Usage

### Safely Deleting Files

Use the `rm` command as you normally would. The files will be moved to your configured trash directory instead of being permanently deleted.

```sh
# This will move my_document.txt to the trash
rm my_document.txt

# This will move the my_folder directory to the trash
rm -r my_folder
```

### Manually Emptying the Trash

To permanently delete all files and directories currently in the trash:

```sh
soft-rm empty-trash
```
