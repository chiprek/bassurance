package cli_cmds

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bassurance",
	Short: "Bassurance CLI tool",
	Long:  "The CLI tool made to connect to the Btask API server",
}

func Execute(version string) error {
	rootCmd.Version = version
	return rootCmd.Execute()
}
