package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "silgo",
		Short: "A golang project",
		Long:  `A golang project.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
