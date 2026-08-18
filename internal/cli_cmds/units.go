package cli_cmds

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

// this func is dedicated to all unit management related tasks
func NewUnitCmd(cfg *Config) *cobra.Command {
	//root Cmd of units
	unitsCmd := &cobra.Command{
		Use:   "units",
		Short: "Manage Units",
	}

	var sortDirection string
	// This lists all available units
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List All units",
		RunE: func(cmd *cobra.Command, args []string) error {
			var query string
			if sortDirection != "" {
				query = "?sort=" + sortDirection
			}

			apiTarget := fmt.Sprintf("%s/units%s", cfg.APIUrl, query)

			resp, err := http.Get(apiTarget)
			if err != nil {
				return fmt.Errorf("failed to reach API: %w", err)
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return fmt.Errorf("failed to read response body: %w", err)
			}
			fmt.Println(string(body))
			return nil
		},
	}
	//listcmd has 1 flag which is sort called by -s --sort this by default is sorted in ascending order
	listCmd.Flags().StringVarP(&sortDirection, "sort", "s", "", "sort order: asc or desc")

	// Attach cmds to unitsCmd
	unitsCmd.AddCommand(listCmd)

	return unitsCmd
}
