package main

import (
	"fmt"

	"github.com/polypmer/peekaboo"
)

func main() {
	ips, err := peekaboo.Possible(peekaboo.Dot192)
	if err != nil {
		fmt.Println(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
