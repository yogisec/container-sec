package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

//  go get github.com/docker/docker/client
// https://godoc.org/github.com/docker/docker/client
// https://pkg.go.dev/github.com/docker/docker/client

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {

		fmt.Println("---- Container HOSTCONFIG ----")
		fmt.Println(container.HostConfig)
		fmt.Println("")

		fmt.Println("---- Container ID ----")
		fmt.Println(container.ID)
		fmt.Println("")

		fmt.Println("---- Container DIFF ----")
		// Pull the diff of the file system
		diffResult, err := cli.ContainerDiff(ctx, container.ID)
		if err != nil {
			panic(err)
		}
		fmt.Println(diffResult)
		fmt.Println("")

		fmt.Println("---- Container INFO ----")
		// Inspect the container
		containerInfo, err := cli.ContainerInspect(ctx, container.ID)
		if err != nil {
			panic(err)
		}
		fmt.Println(containerInfo)
		fmt.Println("")

		// Container Attrs
		fmt.Println("---- Container ATTRS ----")
		friendlyName := containerInfo.Name
		fmt.Printf("Container Name: ")
		fmt.Println(friendlyName)

		stateRunning := containerInfo.State.Running
		currentStatus := containerInfo.State.Status
		fmt.Println("Is running: " + strconv.FormatBool(stateRunning))
		fmt.Println("Current Status: " + currentStatus)

		portBindings := containerInfo.HostConfig.PortBindings
		fmt.Printf("Port Bindings: ")
		fmt.Println(portBindings)

		pid := containerInfo.State.Pid
		fmt.Printf("Container PID: ")
		fmt.Println(pid)

		privleged := containerInfo.HostConfig.Privileged
		fmt.Println("Is privileged: " + strconv.FormatBool(privleged))

		capAdd := containerInfo.HostConfig.CapAdd
		fmt.Printf("CapAdd: ")
		fmt.Println(capAdd)

		capDrop := containerInfo.HostConfig.CapDrop
		fmt.Printf("CappDrop: ")
		fmt.Println(capDrop)

		binds := containerInfo.HostConfig.Binds
		fmt.Printf("Binds?: ")
		fmt.Println(binds)

		mounts := containerInfo.Mounts
		fmt.Printf("Mounts: ")
		fmt.Println(mounts)

		imageHash := containerInfo.Image
		fmt.Println("Image Hash: " + imageHash)

		configImage := containerInfo.Config.Image
		fmt.Printf("Container Image: ")
		fmt.Println(configImage)

		runCommand := containerInfo.Args
		fmt.Printf("Run Command: ")
		fmt.Println(runCommand)

		command := containerInfo.Config.Cmd
		fmt.Printf("Entry Command: ")
		fmt.Println(command)

		tty := containerInfo.Config.Tty
		fmt.Printf("TTY: ")
		fmt.Println(tty)

		createdDate := containerInfo.Created
		fmt.Printf("Container Created: ")
		fmt.Println(createdDate)

		workingDir := containerInfo.Config.WorkingDir
		fmt.Printf("Working Dir: ")
		fmt.Println(workingDir)

		fmt.Println("")

		// Top
		/*fmt.Println("")
		fmt.Println("---- Container TOP ----")
		arguments := []string{"ps"}
		top, err := cli.ContainerTop(ctx, container.ID, arguments)
		if err != nil {
			panic(err)
		}
		fmt.Println(top)
		*/

		// Logs
		fmt.Println("")
		fmt.Println("---- Container LOGS ----")
		logOptions := types.ContainerLogsOptions{
			Follow:     false,
			ShowStdout: true,
			ShowStderr: true,
		}

		out, err := cli.ContainerLogs(ctx, container.ID, logOptions)
		if err != nil {
			panic(err)
		}
		io.Copy(os.Stdout, out)

		/*
			logsOptions := cli.ContainerLogs{
				Container:    container.ID,
				OutputStream: os.Stdout,
				ErrorStream:  os.Stderr,
				Follow:       false,
				Stdout:       true,
				Stderr:       true,
			}
			if err := cli.Logs(logsOptions); err != nil {
				panic(err)
			}
		*/

		/*
			fmt.Println("---- Container Logs ---- ")
			// Container Logs
			// containerLogs, err := cli.ContainerList(ctx, container.
			if err != nil {
				panic(err)
			}
			fmt.Println(containerLogs)
		*/
	}
}
