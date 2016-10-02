package main

import (
	"fmt"

	"github.com/polypmer/peekaboo"
)

func main() {
	//n, _ := peekaboo.LocalNetwork()
	// fmt.Println("Scanning for alives on: ", n)
	// TODO: does this work with the IPNet string?
	// insteald of 0/24 it's something like 144/24
	// It does send pings to the right hosts,
	// but doesn't pick up on all of them?
	alives, _ := peekaboo.AliveHosts("192.168.1.0/24") //n.String())

	fmt.Println("Found: ", alives)

	peekaboo.PortScan(alives)
}
