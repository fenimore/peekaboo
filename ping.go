// ping takes a CIDR, Classless Inter-Domain Routing string
// and pings all possible ips in it's range (0-255). Then
// It'll return a list of Alive IPs.
// See: https://gist.githubusercontent.com/kotakanbe/d3059af990252ba89a82/raw/b5b95e447a81987cb0746884add43c07012e6012/ipcalc.go
// For the basis of this code
package peekaboo

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
)

// Broadcast returns the pings from a broadcast on the network.
func BroadcastPing(ip string) ([]string, error) {
	output, err := exec.Command("ping", "-b", "-c2", ip).Output()
	if err != nil {
		return nil, err
	}
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

// Pong struct exists, because I need to know the status
// of all the pings, not just the successful ones
// INORDER to count the total, and know when ALL
// have been iterated through. I think.
type host struct {
	ip    string
	alive bool
}

// Hosts finds all possible IP addresses in
// range according to the CIDR string.
func hosts(cidr string) ([]string, error) {
	// example: 192.168.1.0/24
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	return ips[1 : len(ips)-1], nil
}

// Incremement?
// TODO: wtf does this do?
// http://play.golang.org/p/m8TNTtygK0
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// As the pingChan is filled with ips, execute the Unix ping
// Command, and if the input returns not at error, add that
// ip to the alive, or pong, channel
func ping(pingChan <-chan string, pongChan chan<- host) {
	for ip := range pingChan {
		var alive bool
		// ping command, -c Count -t ttl
		_, err := exec.Command("ping", "-c1", "-t1", ip).Output()
		if err != nil {
			alive = false
		} else {
			alive = true
		}
		pongChan <- host{ip: ip, alive: alive}
	}
}

// pong receives the pings and sorts the Alive ones. It needs also
// the alive ones, so that way it can be sure that ALL have been checked
// and therefore it can stop checking. Otherwise I'd just pass the live
// straight from ping to done.
// hostNum is the number of Hosts being pinged
func pong(hostNum int, pongChan <-chan host, doneChan chan<- []host) {
	var alives []host
	for i := 0; i < hostNum; i++ {
		res := <-pongChan // take latest ping result
		// fmt.Println("received:", res)
		if res.alive {
			alives = append(alives, res)
		}
		// pass the alives to the done channel
	}
	doneChan <- alives
}

// AliveHosts returns a string of alive ip addresses.
func AliveHosts(cidr string) ([]string, error) {
	hosts, err := hosts(cidr)
	if err != nil {
		return nil, err
	}
	threadNum := 100
	pingChan := make(chan string, 100)      // 100 at a time
	pongChan := make(chan host, len(hosts)) // 255 typically
	doneChan := make(chan []host)

	for i := 0; i < threadNum; i++ {
		// have 100 open at a time
		go ping(pingChan, pongChan)
	}

	// Receive the pongs and pass to done.
	// Do this 255 times. OR until done, qoui.
	go pong(len(hosts), pongChan, doneChan)

	// Send each possible host ip to the pingChan
	// this'll get picked up by the ping func.
	for _, ip := range hosts {
		pingChan <- ip
		fmt.Println("sent: ", ip)
	}

	done := <-doneChan
	alives := make([]string, 0)
	for _, a := range done {
		alives = append(alives, a.ip)
	}

	return alives, nil
}
