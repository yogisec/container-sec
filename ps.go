package main

import (
    "fmt"
    "github.com/prometheus/procfs"
)

func main() {
	p, err := procfs.Self()
	if err != nil {
		log.Fatalf("could not get process: %s", err)
	}

	stat, err := p.NewStat()
	if err != nil {
		log.Fatalf("could not get process stat: %s", err)
	}

	fmt.Printf("command:  %s\n", stat.Comm)
	fmt.Printf("cpu time: %fs\n", stat.CPUTime())
	fmt.Printf("vsize:    %dB\n", stat.VirtualMemory())
	fmt.Printf("rss:      %dB\n", stat.ResidentMemory())
	
}