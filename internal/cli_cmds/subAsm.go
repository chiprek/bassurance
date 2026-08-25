package cli_cmds

import "github.com/spf13/cobra"

func NewSubAsmCmd(cfg *Config) *cobra.Command {
	subAsmCmd := &cobra.Command{
		Use:   "subasm",
		Short: "Manage Sub Assemblies",
	}

	var createName string
	var createSN string
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create sub assembly",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	createCmd.Flags().StringVarP(&createName, "name", "n", "", "name of what the sub assembly is eg: Diesel heater.")
	createCmd.Flags().StringVarP(&createSN, "serial", "s", "", "Insert the serial number if the sub assembly has one eg: measuring head")
	return subAsmCmd
}
