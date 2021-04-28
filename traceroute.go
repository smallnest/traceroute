// Package traceroute provides functions for executing a tracroute to a remote
// host.
package traceroute

import (
	"net"
	"syscall"
	"time"
)

// TraceResult type
type TraceResult struct {
	DestinationAddress [4]byte
	Hops               []Hop
}

func notify(hop Hop, channels []chan Hop) {
	for _, c := range channels {
		c <- hop
	}
}

func closeNotify(channels []chan Hop) {
	for _, c := range channels {
		close(c)
	}
}

// Trace uses the given dest (hostname) and option to execute a traceroute
// from your machine to the remote host.
//
// Outbound packets are UDP packets and inbound packets are ICMP.
//
// Returns a TracerouteResult which contains an array of hops. Each hop includes
// the elapsed time and its IP address.
func Trace(dest string, opt *Option, c ...chan Hop) (result TraceResult, err error) {
	result.Hops = []Hop{}
	destAddr, err := destAddr(dest)
	if err != nil {
		return result, err
	}
	result.DestinationAddress = destAddr
	socketAddr, err := socketAddr()
	if err != nil {
		return
	}

	timeoutMs := (int64)(opt.TimeoutMs())
	tv := syscall.NsecToTimeval(1000 * 1000 * timeoutMs)

	port := opt.Port()

	ttl := opt.FirstHop()
	nqueries := 0
	for {
		if nqueries >= opt.nqueries {
			ttl++
			nqueries = 0
		}
		nqueries++

		// log.Println("TTL: ", ttl)
		start := time.Now()

		// Set up the socket to receive inbound packets
		typ := syscall.SOCK_DGRAM
		if opt.Privileged() {
			typ = syscall.SOCK_RAW
		}
		recvSocket, err := syscall.Socket(syscall.AF_INET, typ, syscall.IPPROTO_ICMP)
		if err != nil {
			return result, err
		}

		// Set up the socket to send packets out.
		sendSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
		if err != nil {
			return result, err
		}
		// This sets the current hop TTL
		err = syscall.SetsockoptInt(sendSocket, 0x0, syscall.IP_TTL, ttl)
		if err != nil {
			return result, err
		}
		// This sets the timeout to wait for a response from the remote host
		err = syscall.SetsockoptTimeval(recvSocket, syscall.SOL_SOCKET, syscall.SO_RCVTIMEO, &tv)
		if err != nil {
			return result, err
		}

		// increment destination port for each probe
		currentPort := port
		if !opt.fixedDstPort {
			port++
			if port == 65535 {
				port = DEFAULT_PORT
			}
		}

		// Bind to the local socket to listen for ICMP packets
		err = syscall.Bind(recvSocket, &syscall.SockaddrInet4{Port: 0, Addr: socketAddr})
		if err != nil {
			return result, err
		}

		// Send a single null byte UDP packet
		sendData := make([]byte, opt.PacketSize())
		err = syscall.Sendto(sendSocket, sendData, 0, &syscall.SockaddrInet4{Port: currentPort, Addr: destAddr})
		if err != nil {
			syscall.Close(recvSocket)
			return result, err
		}

		p := make([]byte, opt.PacketSize()+52)
		n, from, err := syscall.Recvfrom(recvSocket, p, 0)
		elapsed := time.Since(start)
		syscall.Close(recvSocket)
		syscall.Close(sendSocket)

		if err == nil {
			currAddr := from.(*syscall.SockaddrInet4).Addr

			hop := Hop{Success: true, Address: currAddr, N: n, ElapsedTime: elapsed, TTL: ttl}

			if opt.ResolveHost() {
				currHost, err := net.LookupAddr(hop.AddressString())
				if err == nil {
					hop.Host = currHost[0]
				}
			}

			notify(hop, c)

			result.Hops = append(result.Hops, hop)

			if ttl >= opt.MaxHops() || currAddr == destAddr {
				closeNotify(c)
				return result, nil
			}
		} else {
			notify(Hop{Success: false, TTL: ttl}, c)
			if ttl >= opt.MaxHops() {
				closeNotify(c)
				return result, nil
			}
		}

	}
}
