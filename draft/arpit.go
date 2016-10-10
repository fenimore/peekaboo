package peekaboo

import (
	"fmt"
	"net"

	"github.com/mdlayher/arp"
)

func Find() {
	ifaces, _ := net.Interfaces()
	// The third for me is the relevant one.
	client, err := arp.NewClient(&ifaces[1])
	if err != nil {
		fmt.Println(err)
	}

	target := net.ParseIP("10.0.18.70")
	err = client.Request(target)
	if err != nil {
		fmt.Println(err)
	}
}
