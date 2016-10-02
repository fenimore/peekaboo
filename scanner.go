package peekaboo

import (
	"fmt"
	"time"

	portscanner "github.com/anvie/port-scanner"
)

func PortScan(ips []string) {

	for _, ip := range ips {
		fmt.Println("scanning ", ip)
		// Scanner
		ps := portscanner.NewPortScanner(ip, time.Second)
		openedPorts := ps.GetOpenedPort(20, 10000)
		//fmt.Println(openedPorts)
		for i := 0; i < len(openedPorts); i++ {
			port := openedPorts[i]
			fmt.Print(ip, " ", port, " [open]")
			fmt.Println("  -->  ", ps.DescribePort(port))
		}
	}

}
