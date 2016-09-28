package main

import (
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func main() {

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		fmt.Fprintln(os.Stderr, "No Label String specified Example: com.intzen.nginx_name=MainNginx")
		os.Exit(1)
	}

	containerLabelFilter := argsWithoutProg[0]
	defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
	cli, err := client.NewClient("unix:///var/run/docker.sock", "v1.22", nil, defaultHeaders)
	if err != nil {
		panic(err)
	}

	filters := filters.NewArgs()
	filters.Add("label", containerLabelFilter)

	containers, errorGetContainers := cli.ContainerList(context.Background(), types.ContainerListOptions{
		Filter: filters,
	})

	if errorGetContainers != nil {
		panic(errorGetContainers)
	}
	if len(containers) > 0 {
		for _, container := range containers {
			errorKill := cli.ContainerKill(context.Background(), container.ID, "HUP")
			if errorKill != nil {
				panic(errorKill)
			}
			fmt.Println("HUP sent to " + container.ID)
		}
	} else {
		fmt.Println("No containers found for label search " + containerLabelFilter)
	}
}
