package machines

import (
	"context"
	"github.com/anexia-it/go-anxcloud/pkg/client"
	"github.com/anexia-it/go-anxcloud/pkg/vsphere"
	"github.com/go-logr/logr"
	"strings"
	"sync"
)

var waitgroup = &sync.WaitGroup{}

func DeleteVMByPrefix(ctx context.Context, name string) {
	c, err := client.New(client.TokenFromEnv(false))
	if err != nil {
		panic(c)
	}

	get, err := vsphere.NewAPI(c).VMList().Get(ctx, 1, 100)
	if err != nil {
		panic(err)
	}

	deleteChan := make(chan string)
	waitgroup.Add(1)
	go deleteVM(ctx, c, deleteChan)
	for _, vmInfo := range get {
		if strings.HasPrefix(vmInfo.Name, name) {
			deleteChan <- vmInfo.Identifier
		}
	}
	close(deleteChan)
	waitgroup.Wait()
}

func deleteVM(ctx context.Context, cli client.Client, c chan string) {
	logger := logr.FromContextOrDiscard(ctx)
	api := vsphere.NewAPI(cli)
	for id := range c {
		logger.Info("Deleting VM", "ID", id)
		_, err := api.Provisioning().VM().Deprovision(ctx, id, false)
		if err != nil {
			panic(err)
		}
	}

	waitgroup.Done()
}
