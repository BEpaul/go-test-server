package main

import (
	"flag"

	"github.com/BEpaul/go-test-server/init/cmd"
)

var configPathFlag = flag.String("config", "./config.toml", "config file not found")

func main() {
	flag.Parse()
	cmd.NewCmd(*configPathFlag)
}
