package cmd

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"
	"github.com/spf13/cobra"
	"lb-replicator/cmd/vm"
	"log"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "anx-tools",
	Short: "anx-tools is a collection of tools for the Anexia API",
}

func init() {
	setupLogger()
	RootCmd.AddCommand(replicationCmd())
	RootCmd.AddCommand(vm.RootCmd)
}

var logger logr.Logger

func setupLogger() {
	logger = stdr.New(log.New(os.Stdout, "anx-tools", 0))

}
