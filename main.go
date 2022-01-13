package main

import (
	"fmt"
	"lb-replicator/cmd"
	"os"
)

func main() {
	if _, isSet := os.LookupEnv("ANEXIA_TOKEN"); !isSet {
		_, _ = fmt.Fprint(os.Stderr, "'ANEXIA_TOKEN' env var must be set'")
		os.Exit(-1)
	}
	_ = cmd.RootCmd.Execute()
}
