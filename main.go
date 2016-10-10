package main

import (
	"errors"
	"fmt"
	"net"
	"os/exec"
	"strings"

	"github.com/klauspost/oui"
	"github.com/mostlygeek/arp"
)

// LocalAddress returns the local address.
func LocalAddress() (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP, nil
			}
		}
	}
	return nil, errors.New("No IP found")
}

// BroadcastPing returns the pings from a broadcast on the network.
func BroadcastPing(ip string) ([]string, error) {
	output, err := exec.Command("ping", "-b", "-c2", ip).Output()
	if err != nil {
		return nil, err
	}
	//fmt.Println(string(output))
	data := strings.Split(string(output), "\n")
	ips := make([]string, 0)
	for _, d := range data {
		if strings.HasPrefix(d, "64") {
			parts := strings.Split(d, ":")
			ip := strings.Split(parts[0], "from ")
			ips = append(ips, string(ip[1]))

		}
	}
	return ips, nil
}

func Macs(ips []string) []string {
	macs := make([]string, 0)
	table := arp.Table()
	for _, ip := range ips {
		macs = append(macs, table[ip])
	}
	return macs
}

func main() {

	devices, err := BroadcastPing("10.0.255.255")
	if err != nil {
		fmt.Println(err)
	}
	macs := Macs(devices)

	db, err := oui.OpenStaticFile("oui.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Total Devices, %d, and mac addresses %d",
		len(devices), len(macs))

	entry, err := db.Query(macs[0])
	if err == nil {
		fmt.Println(entry)
	} else {
		fmt.Println(err)
	}
}
