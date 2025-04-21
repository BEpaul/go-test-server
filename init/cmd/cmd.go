package cmd

import (
	"github.com/BEpaul/go-test-server/config"
	"github.com/BEpaul/go-test-server/network"
	"github.com/BEpaul/go-test-server/repository"
	"github.com/BEpaul/go-test-server/service"
)

type Cmd struct {
	config     *config.Config
	network    *network.Network
	repository *repository.Repository
	service    *service.Service
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filePath),
	}

	c.repository = repository.NewRepository()
	c.service = service.NewService(c.repository)
	c.network = network.NewNetwork(c.service)

	c.network.ServerStart(c.config.Server.Port)

	return c
}
