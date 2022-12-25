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
	var ip string

	switch len(os.Args) {
	case 2:
		contname = string(os.Args[1])
	case 1:
		contname = ""
	default:
		fmt.Println("Usage: ", os.Args[0], "[cont_name]")
		return
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
			ip = container.NetworkSettings.Networks["bridge"].IPAddress
			fmt.Printf("%s\n", ip)
		}
	}
}
