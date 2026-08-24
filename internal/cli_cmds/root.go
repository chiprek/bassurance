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

// creates the rootCmd
func NewRootCmd(cfg *Config, version string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "bassurance",
		Short:   "Bassurance CLI tool",
		Long:    "The CLI tool made to connect to the Btask API server",
		Version: version,
	}
	//attaches the jobs and units cmd to root cmd.
	rootCmd.AddCommand(NewJobCmd(cfg))
	rootCmd.AddCommand(NewUnitCmd(cfg))
	rootCmd.AddCommand(NewSubAsmCmd(cfg))

	return rootCmd

}
