package main

import (
	"fmt"
	"os"
	"path/filepath"
	"soft-rm/pkg/config"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(emptyTrashCmd)
}

var emptyTrashCmd = &cobra.Command{
	Use:   "empty-trash",
	Short: "Empty the trash",
	Long:  `Permanently delete all items in the trash directory.`, 
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetConfig()
		trashPath := cfg.TrashPath

		files, err := filepath.Glob(filepath.Join(trashPath, "*"))
		if err != nil {
			fmt.Println("Error reading trash directory:", err)
			return
		}

		if len(files) == 0 {
			fmt.Println("Trash is already empty.")
			return
		}

		fmt.Printf("Are you sure you want to permanently delete %d items? [y/N] ", len(files))
		var response string
		fmt.Scanln(&response)

		if response == "y" || response == "Y" {
			for _, file := range files {
				if err := os.RemoveAll(file); err != nil {
					fmt.Fprintf(os.Stderr, "Error deleting %s: %v\n", file, err)
				}
			}
			fmt.Println("Trash emptied.")
		} else {
			fmt.Println("Operation cancelled.")
		}
	},
}
