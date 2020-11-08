package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
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

type fullContainerDetails struct {
	ContainerID      string          `json:"ContainerID"`
	ContainerDetails L.ContainerData `json:"ContainerDetails"`
}

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
	var fullDetails []fullContainerDetails

	// Get Mac Address
	as, err := getMacAddr()
	if err != nil {
		log.Fatal(err)
	}

	mac1 := as[0]
	// fmt.Println(reflect.TypeOf(mac1))
	aid := sha256.Sum256([]byte(mac1))
	fmt.Printf("AID: %x\n", aid)

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
		fullDetails = append(fullDetails, &fullContainerDetails{"ContainerID": container.ID, "ContainerDetails": containerDetails})
		fmt.Println("All Done With Container " + container.ID)

	}

	containerDataJSON, err := json.Marshal(fullDetails)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(containerDataJSON))

	fmt.Println("All Done")
}
