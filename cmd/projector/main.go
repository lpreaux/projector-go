package main

import (
	"encoding/json"
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

	projector := cli.NewProjector(config)

	if config.Operation == cli.Print {
		if len(config.Args) == 0 {
			data := projector.GetAllValue()
			jsonString, err := json.Marshal(data)
			if err != nil {
				log.Fatalf("this line should never be reached %v", err)
			}
			fmt.Printf("%v", string(jsonString))
		} else if value, ok := projector.GetValue(config.Args[0]); ok {
			fmt.Printf("%v", value)
		}
	}

	if config.Operation == cli.Add {
		projector.SetValue(config.Args[0], config.Args[1])
		projector.Save()
	}

	if config.Operation == cli.Remove {
		projector.RemoveValue(config.Args[0])
		projector.Save()
	}
}
