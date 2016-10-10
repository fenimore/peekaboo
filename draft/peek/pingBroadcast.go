package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {

	output, err := exec.Command("ping", "-b", "-c2", "10.0.255.255").Output()
	if err != nil {
		fmt.Println(err)
	}
	data := strings.Split(string(output), "\n")
	ips := make([]string, 0)
	for _, d := range data {
		if strings.HasPrefix(d, "64") {
			parts := strings.Split(d, ":")
			fmt.Println(parts)
			ip := strings.Split(parts[0], "from ")
			fmt.Println(ip)
			ips = append(ips, string(ip[1]))

		}
	}
	for _, i := range ips {
		fmt.Println(i)
	}

}
