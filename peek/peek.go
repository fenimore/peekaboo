package main

import (
	"fmt"

	"github.com/polypmer/peekaboo"
)

func main() {
	//fmt.Println("Pinging Hosts in CIDR")
	//alives, _ := peekaboo.AliveHosts("192.168.1.0/24")

	//fmt.Println(alives)
	a, _ := peekaboo.LocalAddress()
	fmt.Println(a)
	n, _ := peekaboo.LocalNetwork()
	fmt.Println(n)

}
