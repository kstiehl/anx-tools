package vm

import (
	"context"
	"github.com/go-logr/logr"
	"github.com/spf13/cobra"
	"lb-replicator/machines"
)

func vmDeletionCmd(logger logr.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use: "deleteByPrefix",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := logr.NewContext(context.Background(), logger)
			for _, prefix := range args {
				machines.DeleteVMByPrefix(ctx, prefix)
			}
		},
	}
	return cmd
}
