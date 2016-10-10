package peekaboo

import "github.com/mostlygeek/arp"

func Macs(ips []string) []string {
	macs := make([]string, 0)
	table := arp.Table()
	for _, ip := range ips {
		macs = append(macs, table[ip])
	}
	return macs
}
