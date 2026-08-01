package main

import (
	"os"
	"strings"

	_ "embed"

	"github.com/chiprek/bassurance/internal/cli_cmds"
)

var version string

func main() {
	err := cli_cmds.Execute(strings.TrimSpace(version))
	if err != nil {
		os.Exit(1)
	}

}
