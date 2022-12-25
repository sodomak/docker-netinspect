package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {

	var contname string

	if len(os.Args) > 2 {
		fmt.Println("Usage: ", os.Args[0], "[cont_name]")
		return
	}

	if len(os.Args) > 1 {
		contname = string(os.Args[1])
	} else {
		contname = ""
	}

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if strings.Contains(container.Names[0], contname) {
			fmt.Printf("%s\t", container.Names[0][1:])
			for _, v := range container.NetworkSettings.Networks {
				fmt.Printf("%s\n", v.IPAddress)
			}
		}
	}
}

// NetworkSettings Networks IPAddress
