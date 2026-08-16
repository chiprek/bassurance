package cli_cmds

import "github.com/spf13/cobra"

//noun
var unitsCmd = &cobra.Command{
	Use: "units",
	Short: "Manage Units",
}

//verb list
var unitsListCmd = &cobra.Command{
	Use: "list",
	Short: "lists All units",
	RunE: (cfg *Config) func(cmd *cobra.Command, args []string) error {
		var orderBy string
		if sortDirection != "" {
			orderBy = "?sort=" + sortDirection
		}
		
	},
	
}