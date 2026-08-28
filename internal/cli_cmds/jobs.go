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

func NewJobCmd(cfg *Config) *cobra.Command {
	jobsCmd := &cobra.Command{
		Use:   "jobs",
		Short: "Manage Jobs",
	}
	// ListCmd
	var sortDirection string
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "lists the active jobs",
		RunE: func(cmd *cobra.Command, args []string) error {
			var query string
			if sortDirection != "" {
				query = "?sort=" + sortDirection
			}

			apiTarget := fmt.Sprintf("%s/jobs%s", cfg.APIUrl, query)

			resp, err := http.Get(apiTarget)
			if err != nil {
				return fmt.Errorf("failed to reach APIL %w", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("API returned non-200 status: %d", resp.StatusCode)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return fmt.Errorf("failed to read response body: %w", err)
			}
			fmt.Println(string(body))
			return nil
		},
	}
	listCmd.Flags().StringVarP(&sortDirection, "sort", "s", "", "sort order: asc or desc")

	// getCmd

	var getJobName string
	getCmd := &cobra.Command{
		Use:   "get [job name]",
		Short: "Retrieve a specific job",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			getJobName = args[0]

			escapedName := strings.ReplaceAll(getJobName, " ", "%20")

			apiTarget := fmt.Sprintf("%s/jobs/%s", cfg.APIUrl, escapedName)

			resp, err := http.Get(apiTarget)
			if err != nil {
				return fmt.Errorf("failed to reach the API: %w", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("API returned a non 200 status: %d", resp.StatusCode)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return fmt.Errorf("failed to read response body: %w", err)
			}

			fmt.Println(string(body))
			return nil
		},
	}

	// Create Cmd

	var createJobName string
	var createJobStatus string
	createCmd := &cobra.Command{

		Use:   "create [job name]",
		Short: "create a job",
		Long:  "Create a job by passing 1 flag name to the api call and 1 optional flag status",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			createJobName = args[0]
			type RequestPayload struct {
				Name   string `json:"name"`
				Status string `json:"status,omitempty"`
			}

			payload := RequestPayload{
				Name:   createJobName,
				Status: createJobStatus,
			}

			data, err := json.Marshal(payload)
			if err != nil {
				return fmt.Errorf("unable to marshal data to json: %v", err)
			}

			apiTarget := fmt.Sprintf("%s/jobs", cfg.APIUrl)

			resp, err := http.Post(apiTarget, ContentTypeJson, bytes.NewBuffer(data))
			if err != nil {
				return fmt.Errorf("Error making POST request: %w", err)
			}
			defer resp.Body.Close()

			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				return fmt.Errorf("failed to read response body: %w", err)
			}

			if resp.StatusCode >= 200 && resp.StatusCode <= 300 {
				fmt.Println("Success: ", resp.Status)
				fmt.Println(string(bodyBytes))
			} else {
				fmt.Println("Error returned: ", resp.Status)
				fmt.Println("Response body:", string(bodyBytes))
			}

			return nil

		},
	}
	createCmd.Flags().StringVarP(&createJobStatus, "status", "s", "", "The status of the job to be created")

	var attachJobName string
	var attachUnitSn string
	//attach job to unit
	attachCmd := &cobra.Command{
		Use:   "attach [job name] [unit SN]",
		Short: "attach job to unit ",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			attachJobName = args[0]
			attachUnitSn = args[1]

			type RequestPayload struct {
				Serial_number string `json:"serial_number"`
			}

			escapedName := strings.ReplaceAll(attachJobName, " ", "%20")

			payload := RequestPayload{
				Serial_number: attachUnitSn,
			}

			data, err := json.Marshal(payload)
			if err != nil {
				return fmt.Errorf("unable to marshal data to json: %v", err)
			}
			apiTarget := fmt.Sprintf("%s/jobs/%s/units/attach", cfg.APIUrl, escapedName)

			resp, err := http.Post(apiTarget, ContentTypeJson, bytes.NewBuffer(data))
			if err != nil {
				return fmt.Errorf("Error making POST request: %w", err)
			}
			defer resp.Body.Close()

			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				return fmt.Errorf("failed to read response body: %w", err)
			}

			if resp.StatusCode == 204 {
				fmt.Println("Success: ", resp.Status)
				fmt.Println(string(bodyBytes))
			} else {
				fmt.Println("Error returned: ", resp.Status)
				fmt.Println("Response body:", string(bodyBytes))
			}

			return nil
		},
	}

	// Attcach sub commands to parent command
	jobsCmd.AddCommand(listCmd)
	jobsCmd.AddCommand(getCmd)
	jobsCmd.AddCommand(createCmd)
	jobsCmd.AddCommand(attachCmd)

	return jobsCmd
}
