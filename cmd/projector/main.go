package main

import (
	"fmt"
	"log"

	"github.com/MrRTi/projector-go/pkg/cli"
)

func main() {
	opts, err := cli.GetOpts()
	if err != nil {
		log.Fatal("unable to get options %v", err)
	}

	config, err := cli.NewConfig(opts)
	if err != nil {
		log.Fatal("unable to get config %v", err)
	}

	fmt.Printf("opts: %+v", config)
}
