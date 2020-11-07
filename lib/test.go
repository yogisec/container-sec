package lib

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type containerData struct {
	ContainerName string `json:"containername"`
	Platform      string `json:"platform"`
	CurrentStatus string `json:"status"`
	Image         string `json:"imagename"`
}

// GetContainerData pulls configuration data for the running containers
func GetContainerData(containerID string) string {
	fmt.Println("Container ID: " + containerID)

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	// Inspect the container
	containerInfo, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		panic(err)
	}

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

	// Logs
	fmt.Println("")
	fmt.Println("---- Container LOGS ----")
	logOptions := types.ContainerLogsOptions{
		Follow:     false,
		ShowStdout: true,
		ShowStderr: true,
	}

	out, err := cli.ContainerLogs(ctx, containerID, logOptions)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)

	/*
		type ContainerData struct {
		ContainerName string `json:"containername"`
		Platform string `json:"platform"`
		CurrentStatus string `json:"status"`
		Image string `json:"imagename"`
		}
	*/
	contData := &containerData{ContainerName: friendlyName, Platform: platform, CurrentStatus: currentStatus, Image: configImage}

	e, err := json.Marshal(contData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(e))

	return containerID
}

// Dothing I guess all exported functions need a comment explaining why?
func Dothing() {
	s := "12:39:50:2d:a3:b1"

	md5 := md5.Sum([]byte(s))
	sha1 := sha1.Sum([]byte(s))
	sha256 := sha256.Sum256([]byte(s))

	fmt.Printf("%x\n", md5)
	fmt.Printf("%x\n", sha1)
	fmt.Printf("%x\n", sha256)

	fmt.Println("We made it here")
}
