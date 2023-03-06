package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "silgo",
		Short: "A golang project",
		Long: `
	     ___                       ___       ___           ___     
	    /\  \          ___        /\__\     /\  \         /\  \    
	   /::\  \        /\  \      /:/  /    /::\  \       /::\  \   
	  /:/\ \  \       \:\  \    /:/  /    /:/\:\  \     /:/\:\  \  
	 _\:\~\ \  \      /::\__\  /:/  /    /:/  \:\  \   /:/  \:\  \ 
	/\ \:\ \ \__\  __/:/\/__/ /:/__/    /:/__/_\:\__\ /:/__/ \:\__\
	\:\ \:\ \/__/ /\/:/  /    \:\  \    \:\  /\ \/__/ \:\  \ /:/  /
	 \:\ \:\__\   \::/__/      \:\  \    \:\ \:\__\    \:\  /:/  / 
	  \:\/:/  /    \:\__\       \:\  \    \:\/:/  /     \:\/:/  /  
	   \::/  /      \/__/        \:\__\    \::/  /       \::/  /   
	    \/__/                     \/__/     \/__/         \/__/    
		`,
	}
)

// Execute executes the root command.
func Execute() error {
	fmt.Println(rootCmd.Long)
	return rootCmd.Execute()
}
