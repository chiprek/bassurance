package cli_cmds

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

func NewSubAsmCmd(cfg *Config) *cobra.Command {
	subAsmCmd := &cobra.Command{
		Use:   "subasm",
		Short: "Manage Sub Assemblies",
	}

	var createName string
	var unitSN string
	var createSN string
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create sub assembly",
		RunE: func(cmd *cobra.Command, args []string) error {

			type requestParams struct {
				Serial string `json:"serial_number"`
			}
			payload := requestParams{
				Serial: createSN,
			}

			data, err := json.Marshal(payload)
			if err != nil {
				return fmt.Errorf("unable to marshal data to json: %v", err)
			}

			apiTarget := fmt.Sprintf("%s/units/%s/sub-assemblies", cfg.APIUrl, unitSN)

			return nil
		},
	}
	createCmd.Flags().StringVarP(&createName, "name", "n", "", "name of what the sub assembly is eg: Diesel heater.")
	createCmd.Flags().StringVarP(&createSN, "serial", "s", "", "Insert the serial number if the sub assembly has one eg: measuring head")
	return subAsmCmd
}
