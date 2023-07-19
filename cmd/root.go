package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:   "pdgen",
		Short: "The Data Dictionary Generator",
		Long:  "PDGEN is command line tools to generate a data dictionary in different formats.",
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}

// init function to initialize de commanda line
func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
}
