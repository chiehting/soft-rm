package main

import (
	"fmt"
	"os"
	"soft-rm/internal/cleaner"
	"soft-rm/internal/trash"
	"soft-rm/pkg/config"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{

	Use:   "soft-rm [file1] [file2]...",

	Short: "A safe rm command",

	Long: `A safer rm command that moves files to a trash directory instead of de

leting them permanently.`,

	Args: cobra.ArbitraryArgs,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		for _, arg := range args {
			if err := trash.MoveToTrash(arg); err != nil {
				fmt.Fprintf(os.Stderr, "Error moving %s to trash: %v\n", arg, err)
			}
		}
		cleaner.SpawnCleanupProcess()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var recursive, force bool

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "ignored")
	rootCmd.Flags().BoolVarP(&force, "force", "f", false, "ignored")
}

func initConfig() {
	if _, err := config.LoadConfig(); err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}
}
