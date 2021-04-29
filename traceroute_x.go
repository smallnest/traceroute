package traceroute

import (
	"net"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

// TraceX uses the given dest (hostname) and option to execute a traceroute
// from your machine to the remote host by using golang.org/x/net/icmp.
//
// Outbound packets are UDP packets and inbound packets are ICMP.
//
// Returns a TracerouteResult which contains an array of hops. Each hop includes
// the elapsed time and its IP address.
func TraceX(dest string, opt *Option, c ...chan Hop) (result TraceResult, err error) {
	result.Hops = []Hop{}
	dstAddr, err := destAddr(dest)
	if err != nil {
		return result, err
	}
	result.DestinationAddress = dstAddr

	socketIP, err := socketIP()
	if err != nil {
		return result, err
	}

	seq := 0

	port := opt.Port()
	ttl := opt.FirstHop()
	nqueries := 0
	for {
		if nqueries >= opt.nqueries {
			ttl++
			nqueries = 0
		}
		nqueries++

		var conn *icmp.PacketConn
		var err error
		if opt.Privileged() {
			conn, err = icmp.ListenPacket("ip4:icmp", socketIP)
		} else {
			conn, err = icmp.ListenPacket("udp4", socketIP)
		}

		if err != nil {
			return result, err
		}

		conn.SetDeadline(time.Now().Add(time.Duration(opt.TimeoutMs()) * time.Millisecond))

		pconn := conn.IPv4PacketConn()
		pconn.SetControlMessage(ipv4.FlagTTL, true)
		err = pconn.SetTTL(ttl)
		if err != nil {
			return result, err
		}

		seq++
		wm := icmp.Message{
			Type: ipv4.ICMPTypeEcho,
			Code: 0,
			Body: &icmp.Echo{
				ID:   seq,
				Seq:  seq,
				Data: make([]byte, opt.PacketSize()),
			},
		}
		wb, err := wm.Marshal(nil)
		if err != nil {
			notify(Hop{Success: false, TTL: ttl}, c)
			conn.Close()
			continue
		}

		var dst net.IPAddr
		dst.IP = dstAddr[:]

		start := time.Now()
		if opt.Privileged() {
			if _, err := pconn.WriteTo(wb, nil, &dst); err != nil {
				notify(Hop{Success: false, TTL: ttl}, c)
				conn.Close()
				continue
			}
		} else {
			// increment destination port for each probe
			currentPort := port
			if !opt.fixedDstPort {
				port++
				if port == 33535 {
					port = opt.Port()
				}
			}
			if _, err := conn.WriteTo(wb, &net.UDPAddr{IP: net.ParseIP(dest), Port: currentPort}); err != nil {
				notify(Hop{Success: false, TTL: ttl}, c)
				conn.Close()
				continue
			}
		}

		rb := make([]byte, 1500)
		n, _, peer, err := pconn.ReadFrom(rb)
		conn.Close()
		if err != nil {
			notify(Hop{Success: false, TTL: ttl}, c)
			conn.Close()
			continue
		}
		rm, err := icmp.ParseMessage(1, rb[:n])
		if err != nil {
			notify(Hop{Success: false, TTL: ttl}, c)
			conn.Close()
			continue
		}
		elapsed := time.Since(start)

		if err == nil && ((rm.Type == ipv4.ICMPTypeTimeExceeded) || (rm.Type == ipv4.ICMPTypeEchoReply)) {
			peerHost := peer.String()
			if !opt.Privileged() {
				peerHost, _, _ = net.SplitHostPort(peerHost)
			}
			currAddr, _ := parseAddr(peerHost)

			hop := Hop{Success: true, Address: currAddr, N: n, ElapsedTime: elapsed, TTL: ttl}
			if opt.ResolveHost() {
				currHost, err := net.LookupAddr(hop.AddressString())
				if err == nil {
					hop.Host = currHost[0]
				}
			}

			notify(hop, c)
			result.Hops = append(result.Hops, hop)

			if ttl >= opt.MaxHops() || currAddr == dstAddr {
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
