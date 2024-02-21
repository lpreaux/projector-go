package main

import (
	"fmt"
	"github.com/lpreaux/projector-go/pkg/cli"
	"log"
)

func main() {
	opts, err := cli.GetOpts()
	if err != nil {
		log.Fatalf("unable to get options %v", err)
	}

	config, err := cli.NewConfig(opts)
	if err != nil {
		log.Fatalf("unable to get config %v", err)
	}

	fmt.Printf("opts: %+v", opts)
	fmt.Printf("config: %+v", config)
}
