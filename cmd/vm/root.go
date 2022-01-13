package vm

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "machines",
	Short: "commands to manage a virtual machine",
}

func init() {
	RootCmd.AddCommand(vmDeletionCmd())
}
