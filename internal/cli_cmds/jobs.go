package cli_cmds

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var jobName string

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

		fmt.Println("=========\n= debug =\n=========")
		fmt.Printf("%v\n", args)

		resp, err := http.Get("http://localhost:8080/api/v1/jobs")
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

		Name := strings.Replace(jobName, " ", "%20", -1)

		apiTarget := fmt.Sprintf("http://localhost:8080/api/v1/jobs/%s", Name)

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

var createCmd = &cobra.Command{}

func init() {
	getCmd.Flags().StringVarP(&jobName, "name", "n", "", "The exact name of the job to retrive")

	getCmd.MarkFlagRequired("name")

	jobsCmd.AddCommand(getCmd)
	jobsCmd.AddCommand(listCmd)
	rootCmd.AddCommand(jobsCmd)
}
