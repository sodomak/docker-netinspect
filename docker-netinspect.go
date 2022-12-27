package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

const VERSION string = "1.1"

func main() {

	var contname string
	var show_version bool
	var show_help bool

	flag.BoolVar(&show_version, "V", false, "Print version and exit")
	flag.BoolVar(&show_help, "h", false, "Print help and exit")

	flag.Parse()

	if show_version {
		fmt.Println("Version: ", VERSION)
		return
	}

	if show_help {
		help()
		return
	}

	switch len(os.Args) {
	case 2:
		contname = string(os.Args[1])
	case 1:
		contname = ""
	default:
		help()
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
			//TODO: Better solution? We need to iterate the networks due to different names on different systems
			for _, v := range container.NetworkSettings.Networks {
				fmt.Printf("%s\n", v.IPAddress)
			}
		}
	}
}

func help() {
	fmt.Println()
	fmt.Println("Usage: ", os.Args[0], "[<flag>] [cont_name]")
	fmt.Println()
	fmt.Println("Args: ")
	flag.PrintDefaults()
	//flag.Usage()
	fmt.Println()
	fmt.Println("App version:", VERSION)
	fmt.Println()
	fmt.Println("Source: https://github.com/sodomak/docker-netinspect")
	fmt.Println()
	return
}
