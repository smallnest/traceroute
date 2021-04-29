package main

import (
	"flag"
	"fmt"
	"net"
	"sync"

	"github.com/smallnest/traceroute"
)

var (
	lastTTL  = 0
	lastAddr = ""
)

func printHop(hop traceroute.Hop) {
	addr := fmt.Sprintf("%v.%v.%v.%v", hop.Address[0], hop.Address[1], hop.Address[2], hop.Address[3])
	hostOrAddr := addr
	if hop.Host != "" {
		hostOrAddr = hop.Host
	}

	isNewTTL := lastTTL != hop.TTL
	if isNewTTL {
		if hop.TTL == *f {
			fmt.Printf("%-3d ", hop.TTL)
		} else {
			fmt.Printf("\n%-3d ", hop.TTL)
		}

		lastTTL = hop.TTL
	}

	if !hop.Success {
		fmt.Print(" *")
		return
	}

	if lastAddr == addr && !isNewTTL { // only print elapsed time
		fmt.Printf(" %.2f ms", float64(hop.ElapsedTime.Microseconds())/1000)
	} else {
		lastAddr = addr
		newLine := "\n    "
		if isNewTTL {
			newLine = ""
		}
		fmt.Printf("%s %v (%v)  %.2f ms", newLine, hostOrAddr, addr, float64(hop.ElapsedTime.Microseconds())/1000)
	}
}

var (
	m          = flag.Int("m", 32, `Set the max time-to-live (max number of hops) used in outgoing probe packets (default is 64)`)
	f          = flag.Int("f", traceroute.DEFAULT_FIRST_HOP, `Set the first used time-to-live, e.g. the first hop (default is 1)`)
	q          = flag.Int("q", 1, `Set the number of probes per "ttl" to nqueries (default is one probe).`)
	e          = flag.Bool("e", true, "Firewall evasion mode.  Use fixed destination ports for UDP and TCP probes.")
	privileged = flag.Bool("P", true, `privileged or not`)
)

func main() {
	flag.Parse()
	host := flag.Arg(0)
	opt := *traceroute.DefaultOption
	opt.SetNRequeries(*q)
	opt.SetMaxHops(*m)
	opt.SetFirstHop(*f)
	if *privileged {
		opt.EnablePrivileged()
	} else {
		opt.DisablePrivileged()
	}
	if *e {
		opt.EnableFixedDstPort()
	} else {
		opt.DisableFixedDstPort()
	}

	ipAddr, err := net.ResolveIPAddr("ip", host)
	if err != nil {
		return
	}

	fmt.Printf("traceroute to %v (%v), %v hops max, %v byte packets\n", host, ipAddr, opt.MaxHops(), opt.PacketSize())

	var wg sync.WaitGroup
	wg.Add(1)
	c := make(chan traceroute.Hop)
	go func() {
		defer wg.Done()
		for {
			hop, ok := <-c
			if !ok {
				fmt.Println()
				return
			}
			printHop(hop)
		}
	}()

	_, err = traceroute.TraceX(host, &opt, c)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	wg.Wait()
}
