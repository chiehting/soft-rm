package main

import (
	"fmt"
	"soft-rm/pkg/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configViewCmd)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
	Long:  `Manage the configuration for soft-rm.`,
}

var configSetCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Set a configuration key",
	Long:  `Set a configuration key to a specific value.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]
		viper.Set(key, value)
		if err := config.SaveConfig(); err != nil {
			fmt.Println("Error saving config:", err)
		}
	},
}

var configViewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the current configuration",
	Long:  `View the current configuration for soft-rm.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetConfig()
		fmt.Printf("trash_path: %s\n", cfg.TrashPath)
		fmt.Printf("retention_days: %d\n", cfg.RetentionDays)
	},
}
