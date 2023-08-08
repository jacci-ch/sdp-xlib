package netx

import (
	"errors"
	"github.com/jacci-ch/sdp-xlib/valuex"
	"net"
)

var (
	ErrAddrNotFound = errors.New("netx: valid address not found")
)

// GetInterfaceByPrefixIn
// Returns a list of network interface list which name has at least one
// of the prefix in the list.
func GetInterfaceByPrefixIn(prefixes []string) ([]*net.Interface, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, errors.New("netx: " + err.Error())
	}

	var result []*net.Interface
	for cc, i := range interfaces {
		if valuex.HasPrefixIn(i.Name, prefixes) {
			result = append(result, &interfaces[cc])
		}
	}

	return result, nil
}

func GetFirstInterfaceAddr() (string, error) {
	interfaces, err := GetInterfaceByPrefixIn([]string{"eth0", "en"})
	if err != nil {
		return "", errors.New("netx: " + err.Error())
	}

	var ipv6 net.IP
	for _, i := range interfaces {
		if addrs, err := i.Addrs(); err == nil && len(addrs) != 0 {
			for _, addr := range addrs {
				var ip net.IP

				switch addr.(type) {
				case *net.IPNet:
					ip = addr.(*net.IPNet).IP
				case *net.IPAddr:
					ip = addr.(*net.IPAddr).IP
				}

				if ipv4 := ip.To4(); ipv4 != nil {
					return ip.String(), nil // the first IPv4 address
				} else if ipv6 == nil {
					ipv6 = ip
				}
			}
		}
	}

	if ipv6 != nil {
		return ipv6.String(), nil
	}

	return "", ErrAddrNotFound
}
