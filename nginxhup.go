package main

import (
	"fmt"
	"os"

	"github.com/docker/engine-api/client"
	"golang.org/x/net/context"
)

func main() {

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		fmt.Fprintln(os.Stderr, "No container specified")
		os.Exit(1)
	}

	containerName := argsWithoutProg[0]
	defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
	cli, err := client.NewClient("unix:///var/run/docker.sock", "v1.22", nil, defaultHeaders)
	if err != nil {
		panic(err)
	}

	errorKill := cli.ContainerKill(context.Background(), containerName, "HUP")
	if errorKill != nil {
		panic(errorKill)
	}

	fmt.Println("HUP sent")
	/*
		information, errInspect := cli.ContainerInspect(context.Background(), containerName)
		if errInspect != nil {
			panic(errInspect)
		}

		fmt.Println(information.ID)
	*/
}
