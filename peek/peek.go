package main

import "github.com/polypmer/peekaboo"

func main() {
	//fmt.Println("Pinging Hosts in CIDR")
	//alives, _ := peekaboo.AliveHosts("192.168.1.0/24")

	//fmt.Println(alives)
	//a, _ := peekaboo.LocalAddress()
	//fmt.Println(a)
	n, _ := peekaboo.LocalNetwork()
	//fmt.Println(n)

	// TODO: does this work with the IPNet string?
	// insteald of 0/24 it's something like 144/24
	alives, _ := peekaboo.AliveHosts(n.String())

	//fmt.Println(alives)

	peekaboo.PortScan(alives)
}
