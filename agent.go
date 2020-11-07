package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"
	"net"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	L "./lib"
)

/*

  go get github.com/docker/docker/client

  Client Documentation:
  https://pkg.go.dev/github.com/docker/docker/client
  Type Documentation:
  https://pkg.go.dev/github.com/docker/docker/api/types

  UUID Documentation:
  https://pkg.go.dev/github.com/google/uuid

*/

func getMacAddr() ([]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var as []string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as = append(as, a)
		}
	}
	return as, nil
}

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

		fmt.Println("######## Looping Through Containers Pulling Data #########")
		containerDetails := L.GetContainerData(container.ID)
		fmt.Println("All Done With Container " + containerDetails)

		/*
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

			platform := containerInfo.Platform
			fmt.Printf("Platform: ")
			fmt.Println(platform)

			apparmor := containerInfo.AppArmorProfile
			fmt.Printf("AppArmor Profile: ")
			fmt.Println(apparmor)

			stateRunning := containerInfo.State.Running
			currentStatus := containerInfo.State.Status
			fmt.Println("Is running: " + strconv.FormatBool(stateRunning))
			fmt.Println("Current Status: " + currentStatus)

			portBindings := containerInfo.HostConfig.PortBindings
			fmt.Printf("Port Bindings: ")
			fmt.Println(portBindings)

			portInfo := containerInfo.NetworkSettings.Ports
			fmt.Printf("Port Info: ")
			fmt.Println(portInfo)

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
			fmt.Printf("Command: ")
			fmt.Println(command)

			entryPoint := containerInfo.Config.Entrypoint
			fmt.Printf("Entrypoint: ")
			fmt.Println(entryPoint)

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
		*/
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
		/*
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

		// Get Mac Address
		as, err := getMacAddr()
		if err != nil {
			log.Fatal(err)
		}

		mac1 := as[0]
		// fmt.Println(reflect.TypeOf(mac1))
		aid := sha256.Sum256([]byte(mac1))
		fmt.Printf("AID: %x\n", aid)
		fmt.Printf("All Done")

	}
}
