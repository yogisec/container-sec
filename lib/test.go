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

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type containerData struct {
	ContainerName string   `json:"ContainerName"`
	Platform      string   `json:"Platform"`
	AppArmor      string   `json:"AppArmor"`
	CurrentStatus string   `json:"Status"`
	Pid           int      `json:"Pid"`
	Privleged     bool     `json:"Privleged"`
	CapAdd        []string `json:"CapAdd"`
	CapDrop       []string `json:"CapDrop"`
	Image         string   `json:"ImageName"`
	ImageHash     string   `json:"ImageHash"`
	RunCommand    []string `json:"RunCommand"`
	Command       []string `json:"Command"`
	EntryPoint    []string `json:"EntryPoint"`
	TTY           bool     `json:"TTY"`
	CreatedDate   string   `json:"CreatedDate"`
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
	logs, outlogs := io.Copy(os.Stdout, out)

	fmt.Println(logs)
	fmt.Println("OutLogs: ")
	fmt.Println(outlogs)

	/*
		type ContainerData struct {
		ContainerName string `json:"containername"`
		Platform string `json:"platform"`
		CurrentStatus string `json:"status"`
		Image string `json:"imagename"`
		}
	*/
	contData := &containerData{
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
	}

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
