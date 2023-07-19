package cmd

import (
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print client version",
	Long:  `Print the pdgen client version information`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("pdgen v1.0.0")
	},
}

// init function to initialize de commanda line
func init() {
	rootCmd.AddCommand(versionCmd)
}
