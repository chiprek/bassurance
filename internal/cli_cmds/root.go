package cli_cmds

import (
	"github.com/spf13/cobra"
)

type Config struct {
	APIUrl string `toml:"api_url"`
	Output string `toml:"output"`
}

const (
	ContentTypeJson = "aplication/json"
)

func NewRootCmd(cfg *Config, version string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "bassurance",
		Short:   "Bassurance CLI tool",
		Long:    "The CLI tool made to connect to the Btask API server",
		Version: version,
	}

	rootCmd.AddCommand(NewJobCmd(cfg))
	rootCmd.AddCommand(NewUnitCmd(cfg))

	return rootCmd

}
