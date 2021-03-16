package abtesting

import (
	"context"

	"github.com/hashicorp/consul/api"
)

var abClient *AbTesting

func NewAbTesting(addr, path string) error {
	config := api.DefaultConfig()
	config.Address = addr
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}
	abClient = &AbTesting{
		path:   path,
		client: client,
	}
	err = abClient.Init(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func APP() *AbTesting {
	return abClient
}
