package main

import (
	"os"

	"github.com/thediymaker/slurmbeat/cmd"

	_ "github.com/thediymaker/slurmbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
