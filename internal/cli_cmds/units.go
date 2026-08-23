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

	//create cmd
	var createUnitSn string
	var createUnitJob string
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "create a unit from a serial number",
		RunE: func(cmd *cobra.Command, args []string) error {

			escapedName := strings.ReplaceAll(createUnitJob, " ", "%20")

			type RequestParams struct {
				Serial string `json:"serial_number"`
			}

			payload := RequestParams{
				Serial: createUnitSn,
			}

			data, err := json.Marshal(payload)
			if err != nil {
				return fmt.Errorf("unable to marshal data to json: %v", err)
			}

			apiTarget := fmt.Sprintf("%s/jobs/%s/units", cfg.APIUrl, escapedName)

			resp, err := http.Post(apiTarget, ContentTypeJson, bytes.NewBuffer(data))
			if err != nil {
				return fmt.Errorf("Error making POST request: %w", err)
			}
			defer resp.Body.Close()

			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				return fmt.Errorf("failed to read response body: %w", err)
			}

			if resp.StatusCode == 201 {
				fmt.Println("Success: ", resp.Status)
				fmt.Println(string(bodyBytes))
			} else {
				fmt.Println("Error returned: ", resp.Status)
				fmt.Println("Response body:", string(bodyBytes))
			}

			return nil
		},
	}
	// seting the flags for createCmd
	createCmd.Flags().StringVarP(&createUnitSn, "serial_number", "s", "", "Serial number of the unit to be created")
	createCmd.Flags().StringVarP(&createUnitJob, "job", "j", "", "The name of the job this unit will be attached to")
	createCmd.MarkFlagsRequiredTogether("serial_number", "job")
	// Attach cmds to unitsCmd
	unitsCmd.AddCommand(listCmd)
	unitsCmd.AddCommand(createCmd)

	return unitsCmd
}
