package peekaboo

import "net"

var (
	Dot10  = "10.0.0.1/16"
	Dot192 = "192.168.1.0/24"
)

// Hosts finds all possible IP addresses in
// range according to the CIDR string.
func Possible(cidr string) ([]string, error) {
	// example: 192.168.1.0/24
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	// 0 and 255 are not to be called.
	// 255 is the broadcast
	// also 255 is the highest number in a byte.
	return ips[1 : len(ips)-1], nil
}
