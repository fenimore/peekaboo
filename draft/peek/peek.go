package main

import (
	"fmt"

	"github.com/polypmer/peekaboo"
)

func main() {
	n, _ := peekaboo.LocalNetwork()
	//fmt.Println("Scanning for alives on: ", n)
	// TODO: does this work with the IPNet string?
	// insteald of 0/24 it's something like 144/24
	// It does send pings to the right hosts,
	// but doesn't pick up on all of them?
	//alives, _ := peekaboo.AliveHosts("10.0.19.133/16") //n.String()) //("192.168.1.0/24") //

	//ips, _ := peekaboo.BroadcastPing("10.0.255.255")
	ips, _ := peekaboo.BroadcastPing(n.String())
	//macs := peekaboo.Macs(ips)
	//fmt.Println(macs)
	for _, ip := range ips {
		fmt.Println(ip)
	}
	//fmt.Println(len(macs), len(ips))
	//hw, _ := net.ParseMAC(macs[0])
	//peekaboo.PortScan(ips)
}
