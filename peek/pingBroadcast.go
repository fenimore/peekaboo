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
			d = strings.TrimSuffix(d, ":")
			d = strings.TrimPrefix(d, "from ")
			fmt.Println(d)
			ips = append(ips, d)

		}
	}
	//fmt.Println(ips)
}
