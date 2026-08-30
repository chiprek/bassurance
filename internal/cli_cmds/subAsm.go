package cli_cmds

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

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
	var createStatus string
	createCmd := &cobra.Command{
		Use:   "create [unitSN]",
		Short: "Create sub assembly",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			type requestParams struct {
				Name         string `json:"name"`
				SerialNumber string `json:"serial_number"`
				Status       string `json:"status"`
			}
			payload := requestParams{
				Name:         createName,
				SerialNumber: createSN,
				Status:       createStatus,
			}
			unitSN = args[0]
			escapedName := strings.ReplaceAll(unitSN, " ", "%20")

			data, err := json.Marshal(payload)
			if err != nil {
				return fmt.Errorf("unable to marshal data to json: %v", err)
			}

			apiTarget := fmt.Sprintf("%s/units/%s/sub-assemblies", cfg.APIUrl, escapedName)

			resp, err := http.Post(apiTarget, ContentTypeJson, bytes.NewBuffer(data))
			if err != nil {
				return fmt.Errorf("Error making POST request: %w", err)
			}
			defer resp.Body.Close()

			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				return fmt.Errorf("failed to read response body: %w", err)
			}

			if resp.StatusCode == 200 {
				fmt.Println("Sucess: ", resp.Status)
				fmt.Println("Response body: ", string(bodyBytes))
			} else {
				fmt.Println("Error returned: ", resp.Status)
				fmt.Println("Respose body: ", string(bodyBytes))
			}

			return nil
		},
	}
	createCmd.Flags().StringVarP(&createName, "name", "n", "", "name of what the sub assembly is eg: Diesel heater.")
	createCmd.Flags().StringVarP(&createSN, "serial", "s", "", "Insert the serial number if the sub assembly has one eg: measuring head")
	createCmd.Flags().StringVarP(&createStatus, "status", "S", "", "Status of the sub asembly eg: completed ongoing or prepairing.")

	var listUnitSN string
	var sortDirection string

	listAsm := &cobra.Command{
		Use:   "list [unit SN]",
		Short: "list the sub assemblies on the unit",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var query string
			if sortDirection != "" {
				query = "?sort=" + sortDirection
			}
			listUnitSN = args[0]
			escapedName := strings.ReplaceAll(listUnitSN, " ", "%20")
			apiTarget := fmt.Sprintf("%s/units/%s/sub-assemblies/%s", cfg.APIUrl, escapedName, query)

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
	listAsm.Flags().StringVarP(&sortDirection, "sort", "s", "", "Sort order: asc or desc")

	subAsmCmd.AddCommand(createCmd)
	subAsmCmd.AddCommand(listAsm)
	return subAsmCmd
}
