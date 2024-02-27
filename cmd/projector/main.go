package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/MrRTi/projector-go/pkg/cli"
)

func main() {
	opts, err := cli.GetOpts()
	if err != nil {
		log.Fatal("unable to get options", err)
	}

	config, err := cli.NewConfig(opts)
	if err != nil {
		log.Fatal("unable to get config", err)
	}

	proj := cli.NewProjector(config)

	if config.Operation == cli.Print {
		if len(config.Args) == 0 {
			data := proj.GetValueAll()
			jsonString, err := json.Marshal(data)
			if err != nil {
				log.Fatal("this line should never be reached", err)
			}

			fmt.Printf("%v", string(jsonString))
		} else if value, ok := proj.GetValue(config.Args[0]); ok {
			fmt.Printf("%v", value)	
		}
	}

	if config.Operation == cli.Add {
		proj.SetValue(config.Args[0], config.Args[1])
		proj.Save()
	}

	if config.Operation == cli.Remove {
		proj.RemoveValue(config.Args[0])
		proj.Save()
	}
}
