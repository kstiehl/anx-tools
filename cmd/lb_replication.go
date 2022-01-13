package cmd

import (
	"context"
	"fmt"
	"github.com/anexia-it/go-anxcloud/pkg/api"
	"github.com/anexia-it/go-anxcloud/pkg/client"
	"github.com/go-logr/logr"
	"github.com/spf13/cobra"
	"lb-replicator/lbaas/components"
	"lb-replicator/lbaas/replication"
	"os"
	"sync"
)

func replicationCmd() *cobra.Command {
	var targetLB string
	var sourceLB string

	cmd := &cobra.Command{
		Use:   "lb-relplicate",
		Short: "replicate a loadbalancer configuration from one lb to another",
		Run: func(cmd *cobra.Command, args []string) {
			if targetLB == "" || sourceLB == "" {
				_ = cmd.Usage()
			}
			anxClient := createClient()
			var config, target components.HashedLoadBalancer
			var waitGroup sync.WaitGroup
			waitGroup.Add(2)

			ctx := logr.NewContext(context.Background(), logger)

			go func() {
				defer waitGroup.Done()
				var err error
				config, err = replication.FetchLoadBalancer(ctx, sourceLB, anxClient)
				if err != nil {
					logger.Error(err, "could not fetch source loadbalancers lb config")
					os.Exit(-1)
				}
			}()

			go func() {
				defer waitGroup.Done()
				var err error
				target, err = replication.FetchLoadBalancer(ctx, targetLB, anxClient)
				if err != nil {
					logger.Error(err, "could not fetch target loadbalancers lb config")
					os.Exit(-1)
				}
			}()

			waitGroup.Wait()
			fmt.Printf("%v\n", target)
			fmt.Printf("%v\n", config)

			err := replication.SyncLoadBalancer(ctx, anxClient, config, target)
			if err != nil {
				panic(err)
			}
		},
	}

	cmd.PersistentFlags().StringVarP(&targetLB, "target", "t", "",
		"specify the target loadbalancer. the loadbalancer that will receive the configuration")
	cmd.PersistentFlags().StringVarP(&sourceLB, "source", "s", "",
		"specify the source loadbalancer. the loadbalancer from which the configuration is copied")
	err := cmd.MarkPersistentFlagRequired("source")
	if err != nil {
		panic(err)
	}
	err = cmd.MarkPersistentFlagRequired("target")
	if err != nil {
		panic(err)
	}
	return cmd
}

func createClient() api.API {
	client, err := api.NewAPI(api.WithClientOptions(client.TokenFromEnv(false)))
	if err != nil {
		panic(err)
	}
	return client
}
