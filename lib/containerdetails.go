package containerdetails

import (
	"bytes"
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// ContainerData is returned after each iteration
type ContainerData struct {
	ContainerName string                 `json:"ContainerName"`
	Platform      string                 `json:"Platform"`
	AppArmor      string                 `json:"AppArmor"`
	CurrentStatus string                 `json:"Status"`
	Pid           int                    `json:"Pid"`
	Privleged     bool                   `json:"Privleged"`
	CapAdd        []string               `json:"CapAdd"`
	CapDrop       []string               `json:"CapDrop"`
	Image         string                 `json:"ImageName"`
	ImageHash     string                 `json:"ImageHash"`
	RunCommand    []string               `json:"RunCommand"`
	Command       []string               `json:"Command"`
	EntryPoint    []string               `json:"EntryPoint"`
	TTY           bool                   `json:"TTY"`
	CreatedDate   string                 `json:"CreatedDate"`
	WorkingDir    string                 `json:"WorkingDir"`
	Logs          string                 `json:"Logs"`
	Mounts        []string               `json:"Mounts"`
	PortBindings  map[string]interface{} `json:"PortBindings"`
}

// GetContainerData pulls configuration data for the running containers
func GetContainerData(containerID string) *ContainerData {
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
	friendlyName := containerInfo.Name               // Container Friendly Name
	platform := containerInfo.Platform               // Linux, windows?
	apparmor := containerInfo.AppArmorProfile        // App Armor info
	currentStatus := containerInfo.State.Status      // Get the containers current status
	pid := containerInfo.State.Pid                   // Host Process ID for container
	privleged := containerInfo.HostConfig.Privileged // is the container running with --priv
	capAdd := containerInfo.HostConfig.CapAdd        // Any added capabilities?
	capDrop := containerInfo.HostConfig.CapDrop      // Any dropped capabilities?
	configImage := containerInfo.Config.Image        // Container Image
	imageHash := containerInfo.Image                 // Image hash
	runCommand := containerInfo.Args                 // Run command
	command := containerInfo.Config.Cmd              // Config Command
	entryPoint := containerInfo.Config.Entrypoint    // Container entrypoint
	tty := containerInfo.Config.Tty                  // tty
	createdDate := containerInfo.Created             // Container launch date/time
	workingDir := containerInfo.Config.WorkingDir    // Container working directory

	// Pull Container Logs
	logOptions := types.ContainerLogsOptions{
		Follow:     false,
		ShowStdout: true,
		ShowStderr: true,
	}

	buf := new(bytes.Buffer)
	out, err := cli.ContainerLogs(ctx, containerID, logOptions)
	if err != nil {
		panic(err)
	}

	buf.ReadFrom(out)
	logString := buf.String()

	// fmt.Println(logString)

	// Top
	/*
		fmt.Println("")
		fmt.Println("---- Container TOP ----")
		arguments := []string{"ps"}
		top, err := cli.ContainerTop(ctx, container.ID, arguments)
		if err != nil {
			panic(err)
		}
		fmt.Println(top)
	*/

	portBindings := containerInfo.HostConfig.PortBindings
	fmt.Printf("Port Bindings: ")
	fmt.Println(portBindings)

	portInfo := containerInfo.NetworkSettings.Ports
	fmt.Printf("Port Info: ")
	fmt.Println(portInfo)

	// fmt.Println("Is privileged: " + strconv.FormatBool(privleged))

	binds := containerInfo.HostConfig.Binds
	fmt.Printf("Binds?: ")
	fmt.Println(binds)

	mounts := containerInfo.Mounts
	fmt.Printf("Mounts: ")
	fmt.Println(mounts)

	// Pull associate the variables with the containerData struct
	contData := &ContainerData{
		ContainerName: friendlyName,
		Platform:      platform,
		AppArmor:      apparmor,
		Pid:           pid,
		Privleged:     privleged,
		CapAdd:        capAdd,
		CapDrop:       capDrop,
		CurrentStatus: currentStatus,
		Image:         configImage,
		ImageHash:     imageHash,
		RunCommand:    runCommand,
		Command:       command,
		EntryPoint:    entryPoint,
		TTY:           tty,
		CreatedDate:   createdDate,
		WorkingDir:    workingDir,
		Logs:          logString,
	}

	// Returned the raw container struct json
	return contData
}
