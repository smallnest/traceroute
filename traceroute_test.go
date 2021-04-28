package traceroute

import (
	"fmt"
	"testing"
)

func printHop(hop Hop) {
	fmt.Printf("%-3d %v (%v)  %v\n", hop.TTL, hop.String(), hop.AddressString(), hop.ElapsedTime)
}

var testOption = &Option{
	port:       DEFAULT_PORT,
	maxHops:    20,
	firstHop:   1,
	timeoutMs:  DEFAULT_TIMEOUT_MS,
	nqueries:   1,
	packetSize: 0,
	privileged: false,
}

var testHost = "114.114.114.114"

func TestTrace(t *testing.T) {
	t.Log("Testing synchronous traceroute")
	out, err := Trace(testHost, testOption)
	if err == nil {
		if len(out.Hops) == 0 {
			t.Errorf("TestTraceroute failed. Expected at least one hop")
		}
	} else {
		t.Errorf("TestTraceroute failed due to an error: %v", err)
	}

	for _, hop := range out.Hops {
		printHop(hop)
	}
}

func TestTraceChannel(t *testing.T) {
	t.Log("Testing asynchronous traceroute")
	c := make(chan Hop)
	go func() {
		for {
			hop, ok := <-c
			if !ok {
				fmt.Println()
				return
			}
			printHop(hop)
		}
	}()

	out, err := Trace(testHost, testOption, c)
	if err == nil {
		if len(out.Hops) == 0 {
			t.Errorf("TestTracerouteChannel failed. Expected at least one hop")
		}
	} else {
		t.Errorf("TestTraceroute failed due to an error: %v", err)
	}
}
