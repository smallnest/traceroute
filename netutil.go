package traceroute

import (
	"errors"
	"net"
	"strings"
)

// import (
// 	"errors"
// 	"net"
// )

// func socketIP() (ip string, err error) {
// 	addrs, err := net.InterfaceAddrs()
// 	if err != nil {
// 		return
// 	}

// 	for _, a := range addrs {
// 		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
// 			if len(ipnet.IP.To4()) == net.IPv4len {
// 				return ipnet.IP.To4().String(), nil
// 			}
// 		}
// 	}
// 	err = errors.New("You do not appear to be connected to the Internet")
// 	return
// }

func socketIPv6() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if len(ipnet.IP.To4()) != net.IPv4len {
				localIP := ipnet.IP.To16().String()
				if localIP == "::1" || strings.HasPrefix(localIP, "fe80::") {
					continue
				}

				return ipnet.IP.To16().String(), nil
			}
		}
	}
	err = errors.New("You do not appear to be connected to the Internet")
	return
}

// // Given a host name convert it to a 4 byte IP address.
// func destAddr(dest string) (destAddr [4]byte, err error) {
// 	addrs, err := net.LookupHost(dest)
// 	if err != nil {
// 		return
// 	}
// 	addr := addrs[0]

// 	ipAddr, err := net.ResolveIPAddr("ip", addr)
// 	if err != nil {
// 		return
// 	}
// 	copy(destAddr[:], ipAddr.IP.To4())
// 	return
// }

// func parseAddr(addr string) (destAddr [4]byte, err error) {
// 	ipAddr, err := net.ResolveIPAddr("ip", addr)
// 	if err != nil {
// 		return
// 	}
// 	copy(destAddr[:], ipAddr.IP.To4())
// 	return
// }
