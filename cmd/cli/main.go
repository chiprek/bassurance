package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := buildGetJobCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

}

func buildGetJobCommand() *cobra.Command {
	getjob := &cobra.Command{
		Use: "getjob",
	}
	getjob.AddCommand()
	return getjob
}
