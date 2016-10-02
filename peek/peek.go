package main

import (
	"fmt"

	"github.com/polypmer/peekaboo"
)

func main() {
	n, _ := peekaboo.LocalNetwork()
	fmt.Println("Scanning for alives on: ", n)
	// TODO: does this work with the IPNet string?
	// insteald of 0/24 it's something like 144/24
	// It does send pings to the right hosts,
	// but doesn't pick up on all of them?
	alives, _ := peekaboo.AliveHosts(n.String()) //("192.168.1.0/24") //

	macs := peekaboo.Macs(alives)
	fmt.Println(macs)
	// This takes way too long
	//peekaboo.PortScan(alives)
}
