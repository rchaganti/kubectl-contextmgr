package main

import (
	"contextmgr/cmd"
	"fmt"
	"os"
	"time"
)

var (
	version = "v0.0.7"
	commit  = "none"
	date    = "unknown"
)

func main() {
	// Set the version information
	date = time.Now().Format(time.RFC3339)
	cmd.SetVersionInfo(version, commit, date)

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
