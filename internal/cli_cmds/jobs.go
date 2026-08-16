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

// noun
var jobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "Manage Jobs",
}

// Verb List
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists the active jobs",
	RunE: func(cmd *cobra.Command, args []string) error {
		var orderBy string
		if sortDirection != "" {
			orderBy = "?sort=" + sortDirection
		}

		apiTarget := "http://localhost:8080/api/v1/jobs" + orderBy

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

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a specific job",
	RunE: func(cmd *cobra.Command, args []string) error {

		Name := strings.ReplaceAll(jobName, " ", "%20")

		apiTarget := "http://localhost:8080/api/v1/jobs/" + Name

		fmt.Printf("Getting: %s\n", apiTarget)

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

var createCmd = &cobra.Command{

	Use:   "create",
	Short: "create a job",
	Long:  "Create a job by passing 1 flag name to the api call and 1 optional flag status",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {

		flagName := jobName
		flagStatus := jobStatus

		type RequestPayload struct {
			Name   string `json:"name"`
			Status string `json:"status,omitempty"`
		}

		payload := RequestPayload{
			Name:   flagName,
			Status: flagStatus,
		}

		data, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("unable to marshal data to json: %v", err)
		}

		apiTarget := "http://localhost:8080/api/v1/jobs"

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

func init() {
	listCmd.Flags().StringVarP(&sortDirection, "sort", "s", "", "default sort is asc, pass value desc for descending. sort is based off creation time")

	getCmd.Flags().StringVarP(&jobName, "name", "n", "", "The exact name of the job to retrive")
	getCmd.MarkFlagRequired("name")

	createCmd.Flags().StringVarP(&jobName, "name", "n", "", "The name of the job to be created")
	createCmd.Flags().StringVarP(&jobStatus, "status", "s", "", "The status of the job to be created")
	createCmd.MarkFlagRequired("name")

	jobsCmd.AddCommand(createCmd)
	jobsCmd.AddCommand(getCmd)
	jobsCmd.AddCommand(listCmd)
	rootCmd.AddCommand(jobsCmd)
}
