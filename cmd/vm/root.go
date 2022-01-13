package vm

import (
	"github.com/go-logr/logr"
	"github.com/spf13/cobra"
)

func RootCmd(logger logr.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "machines",
		Short: "commands to manage a virtual machine",
	}
	cmd.AddCommand(vmDeletionCmd(logger))

	return cmd
}
