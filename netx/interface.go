// Copyright 2023 to now() The SDP Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package netx

import (
	"errors"
	"github.com/jacci-ch/sdp-xlib/valuex"
	"net"
)

// InterfacesWithPrefixIn
//
// Returns a list of network interfaces which name has at least one
// of the prefix in the list.
func InterfacesWithPrefixIn(prefixes []string) ([]*net.Interface, error) {
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

// FirstInterfaceAddr - retrieves the first non-virtual interfaces IP address.
// The prefixes:
//
//	eth - the classical linux
//	en  - ens33 (centos 7+); en0 - mac OS X.
//
// This function use prefix match for local HW interfaces.
func FirstInterfaceAddr() (string, error) {
	interfaces, err := InterfacesWithPrefixIn([]string{"eth", "en"})
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

	// No IPv4 is found, so we have to returns an IPv6 address.
	if ipv6 != nil {
		return ipv6.String(), nil
	}

	return "", errors.New("netx: valid address not found")
}
